package aictx

import (
	"bufio"
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/yarlson/pin"

	"github.com/amberpixels/aictx/internal/fsutils"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

// App encapsulates the configuration and dependencies for the application.
type App struct {
	Lgr *log.Logger

	// InputPath is the path to the file, directory, or git repository URL to process.
	InputPath string

	// Local, when true, forces the input to be treated as a local directory.
	Local bool

	// Include is an optional global glob pattern to include files (supports comma-separated lists).
	Include string

	// Exclude is an optional global glob pattern to exclude files (supports comma-separated lists).
	Exclude string

	// TreeEnabled indicates whether to output the directory tree.
	// Tree mode is enabled by default unless explicitly disabled.
	TreeEnabled bool

	// TreeInclude is an optional glob pattern to include files specifically for tree mode.
	// If not specified, the global Include is used.
	TreeInclude string

	// TreeExclude is an optional glob pattern to exclude files specifically for tree mode.
	// If not specified, the global Exclude is used.
	TreeExclude string

	// TreeShowHidden indicates whether hidden files should be shown in tree mode.
	TreeShowHidden bool

	// SourceEnabled indicates whether to output the file contents.
	SourceEnabled bool

	// SourceInclude is an optional glob pattern to include files specifically for source mode.
	// If not specified, the global Include is used.
	SourceInclude string

	// SourceExclude is an optional glob pattern to exclude files specifically for source mode.
	// If not specified, the global Exclude is used.
	SourceExclude string

	// SourceShowHidden indicates whether hidden files should be shown in source mode.
	SourceShowHidden bool

	// SourceThreshold is the maximum file size (in MB) allowed for source output.
	SourceThreshold float64

	// Out is the destination writer where output will be written.
	Out io.Writer

	// OutFilename holds the output file name (if not stdout) so that it can be ignored during processing.
	OutFilename string

	// NoCoreIgnores disables the hardcoded core ignore patterns.
	NoCoreIgnores bool

	// NoGitIgnore disables respecting the .gitignore file.
	NoGitIgnore bool

	// AictxIgnore holds additional exclude patterns loaded from a .aictxignore file.
	AictxIgnore []string

	// Raw, when true, concatenates file contents in source mode without headers or summary.
	Raw bool

	// Verbose, when true, prints verbose output.
	Verbose bool
}

// Run executes the main application logic.
// It will either print a tree view of the allowed files or concatenate file contents.
//
//nolint:gocognit // TODO: refactor this at some point.
func (a *App) Run(ctx context.Context) error {
	if !a.TreeEnabled && !a.SourceEnabled {
		return errors.New("at least one of tree or source mode must be enabled")
	}

	p := pin.New(".",
		pin.WithSpinnerColor(pin.ColorMagenta),
		pin.WithTextColor(pin.ColorYellow),
	)
	var pCancel context.CancelFunc

	var fsys billy.Filesystem
	//nolint:nestif // we're OK with this
	if a.InputPath == "." || a.Local {
		// Use OS filesystem.
		root := "."
		if strings.HasPrefix(a.InputPath, "/") {
			root = "/"
		}
		fsys = osfs.New(root)
		a.Local = true

		absPath, _ := filepath.Abs(a.InputPath)
		if absPath == "" {
			absPath = "."
		}

		if a.Verbose {
			p.UpdateMessage("Loading local path...")
			pCancel = p.Start(ctx)
			defer pCancel()

			p.Stop(fmt.Sprintf(`Loaded local path "%s"`, absPath))
		}
	} else {
		// Treat inputPath as a Git repository URL.
		repoURL, branch, err := ValidateGitRepoName(a.InputPath)
		if err != nil {
			return fmt.Errorf("invalid git repository URL[%s]: %w", a.InputPath, err)
		}

		strRepoURL := repoURL
		if branch != "" {
			strRepoURL = strRepoURL + " (branch " + branch + ")"
		}

		if a.Verbose {
			p.UpdateMessage(fmt.Sprintf("Cloning %s...", strRepoURL))
			pCancel = p.Start(ctx)
			defer pCancel()
		}

		gitFS, err := ReadGit(repoURL, branch)
		if err != nil {
			p.Stop(fmt.Sprintf("Failed on cloning %s", strRepoURL))
			return fmt.Errorf("failed to load git repo: %w", err)
		}
		fsys = gitFS

		// Reset input path to root.
		a.InputPath = "."

		if a.Verbose {
			p.Stop(fmt.Sprintf("Cloned %s", strRepoURL))
		}
	}

	info, err := fsys.Stat(a.InputPath)
	if err != nil {
		return fmt.Errorf("failed to access input path '%s': %w", a.InputPath, err)
	}

	// If the input is a directory, attempt to load .aictxignore.
	if info.IsDir() {
		ignorePatterns, err := loadDotIgnoreFromFS(fsys, ".aictxignore", a.InputPath)
		if err != nil {
			return fmt.Errorf("error reading .aictxignore: %w", err)
		}
		a.AictxIgnore = ignorePatterns

		// Load .gitignore patterns unless disabled by --no-git-ignore
		if !a.NoGitIgnore {
			gitIgnorePatterns, err := loadDotIgnoreFromFS(fsys, ".gitignore", a.InputPath)
			if err != nil {
				return fmt.Errorf("error reading .gitignore: %w", err)
			}
			a.AictxIgnore = append(a.AictxIgnore, gitIgnorePatterns...)
		}
	}

	if a.TreeEnabled {
		if err := a.displayTree(ctx, fsys, info, a.Out, p); err != nil {
			return err
		}
	}

	if a.SourceEnabled {
		if a.TreeEnabled {
			// let's have an empty line between tree and source
			fmt.Fprintln(a.Out)
		}
		if err := a.displaySource(ctx, fsys, info, a.Out, p); err != nil {
			return err
		}
	}

	if a.Verbose {
		cancel := p.Start(ctx)
		f, _ := fsys.Stat(a.OutFilename)
		p.Stop(fmt.Sprintf(
			"Dumped to file %s (%s)",
			f.Name(), formatSize(f.Size()),
		))
		cancel()
	}

	return nil
}

// displayTree processes tree mode: it prints a filtered directory tree.
// In this updated version, we first build a filtered tree structure, then
// print a summary line, and finally print the tree structure.
func (a *App) displayTree(ctx context.Context, fsys billy.Filesystem, info os.FileInfo, w io.Writer, p *pin.Pin) error {
	// If input is not a directory, simply print it if allowed.
	if !info.IsDir() {
		if a.isAllowed(a.InputPath, false) {
			fmt.Fprintln(w, filepath.Base(a.InputPath))
		}
		return nil
	}

	var s summary

	if a.Verbose {
		p.UpdateMessage("Calculating tree...")
		cancel := p.Start(ctx)
		defer func() {
			cancel()
			p.Stop(fmt.Sprintf(
				"Calculated tree %d files (%s)",
				s.fileCount, formatSize(s.totalSize),
			))
		}()
	}

	// Build the filtered tree structure.
	rootNode, err := a.filterTree(ctx, fsys, a.InputPath)
	if err != nil {
		if errors.Is(err, ErrFilterSkipped) {
			// nothing to show
			return nil
		}
		return fmt.Errorf("error filtering tree: %w", err)
	}

	// Compute summary.
	s = rootNode.summary()
	// Format the summary concisely.
	summaryStr := fmt.Sprintf(
		"Project Tree [%d files, %s total, max %s] (* - for binary files)",
		s.fileCount, formatSize(s.totalSize), formatSize(s.maxSize),
	)

	// Print the root node with the summary appended.
	fmt.Fprintln(w, summaryStr)

	// Print the tree structure (no further filtering here).
	rootNode.printTree("", w)
	return nil
}

// displaySource processes source mode: it builds a filtered tree of source files,
// computes a summary, prints the summary, and then prints the content of each file.
func (a *App) displaySource(ctx context.Context, fs billy.Filesystem, info os.FileInfo, w io.Writer, p *pin.Pin) error {
	var rootNode *TreeNode
	var err error

	var s summary
	if a.Verbose {
		p.UpdateMessage("Concatenating source files...")
		cancel := p.Start(ctx)
		defer func() {
			cancel()
			p.Stop(fmt.Sprintf(
				"Concatenated source of %d files (%s)",
				s.fileCount, formatSize(s.totalSize),
			))
		}()
	}

	// If the input is a file, process it directly.
	if !info.IsDir() { //nolint: nestif // we're OK with this
		// If not allowed or exceeds threshold, skip.
		if !a.isAllowed(a.InputPath, true) || exceedsThreshold(info.Size(), a.SourceThreshold) {
			return nil
		}
		rootNode = &TreeNode{
			Name:  filepath.Base(a.InputPath),
			Path:  a.InputPath,
			IsDir: false,
			Size:  info.Size(),
		}
	} else {
		// Build the filtered tree structure for source mode.
		rootNode, err = a.filterSourceTree(ctx, fs, a.InputPath)
		if err != nil {
			if errors.Is(err, ErrFilterSkipped) {
				// Nothing to display.
				return nil
			}
			return fmt.Errorf("error filtering source files: %w", err)
		}
	}

	if a.Raw {
		// Raw mode: simply print the file contents without summary or fancy headers.
		return a.printSourceFilesRaw(ctx, fs, rootNode, w)
	}

	// Compute summary.
	s = rootNode.summary()

	fmt.Fprintf(w,
		"Project Source [%d files, %s total, max %s]\n",
		s.fileCount, formatSize(s.totalSize), formatSize(s.maxSize),
	)

	// Now display the source content.
	var counter int
	if err = a.printSourceFiles(ctx, fs, rootNode, w, s.fileCount, &counter); err != nil {
		return err
	}

	return nil
}

// filterSourceTree recursively builds a tree of allowed source files/directories.
// It uses isAllowed in source mode and skips files that exceed the size threshold.
//
//nolint:gocognit // TODO: refactor this at some point.
func (a *App) filterSourceTree(ctx context.Context, fs billy.Filesystem, root string) (*TreeNode, error) {
	// Check cancellation.
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	info, err := fs.Stat(root)
	if err != nil {
		return nil, err
	}

	// For files, check allowed and threshold.
	if !info.IsDir() {
		if !a.isAllowed(root, true) || exceedsThreshold(info.Size(), a.SourceThreshold) {
			return nil, ErrFilterSkipped
		}
		return &TreeNode{
			Name:  info.Name(),
			Path:  root,
			IsDir: false,
			Size:  info.Size(),
		}, nil
	}

	// For directories.
	node := &TreeNode{
		Name:  info.Name(),
		Path:  root,
		IsDir: true,
	}

	entries, err := fs.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		childPath := filepath.Join(root, entry.Name())
		if entry.IsDir() { //nolint: nestif // we're OK with this
			// Only include directories if they or their descendants are allowed.
			ok, err := a.hasAllowed(fs, childPath, true)
			if err != nil {
				return nil, err
			}
			if !ok {
				continue
			}
			childNode, err := a.filterSourceTree(ctx, fs, childPath)
			if err != nil && !errors.Is(err, ErrFilterSkipped) {
				return nil, err
			}
			if childNode != nil {
				node.Children = append(node.Children, childNode)
			}
		} else if a.isAllowed(childPath, true) {
			childInfo, err := fs.Stat(childPath)
			if err != nil {
				return nil, err
			}
			// Skip files exceeding threshold.
			if exceedsThreshold(childInfo.Size(), a.SourceThreshold) {
				continue
			}
			node.Children = append(node.Children, &TreeNode{
				Name:  entry.Name(),
				Path:  childPath,
				IsDir: false,
				Size:  childInfo.Size(),
			})
		}
	}
	// If directory has no allowed children, return nil.
	if len(node.Children) == 0 {
		return nil, ErrFilterSkipped
	}
	return node, nil
}

// matchPattern returns true if the given pattern matches the provided path.
// The matching logic is as follows:
//
//  1. If the pattern begins with "/" then it is considered "anchored" to the
//     root. In that case, we only match if the normalized path does not contain
//     any slashes (i.e. the file is in the root folder). The leading "/" is removed
//     before matching.
//  2. If the pattern ends with "/**", it is treated as a prefix match.
//  3. If the pattern has no glob wildcards, we check whether any segment of the path
//     equals the pattern, or if the path starts with the pattern followed by a slash.
//  4. Otherwise, if the pattern contains wildcards and a slash, we match against the
//     full normalized path; if no slash is present, we match against just the base name.
func matchPattern(pattern, pathStr string) bool {
	// 1. Anchored pattern: must start with "/" to indicate matching only the root.
	if strings.HasPrefix(pattern, "/") {
		// Only match if the file is in the root folder.
		if strings.Contains(pathStr, "/") {
			return false
		}
		// Remove the leading "/" and match against the file (which is just the base name).
		stripped := pattern[1:]
		if !strings.ContainsAny(stripped, "*?[") {
			return stripped == pathStr
		}
		match, err := filepath.Match(stripped, pathStr)
		return err == nil && match
	}

	// 2. Special handling: if pattern ends with "/**", treat it as a prefix match.
	if strings.HasSuffix(pattern, "/**") {
		prefix := strings.TrimSuffix(pattern, "/**")
		return pathStr == prefix || strings.HasPrefix(pathStr, prefix+"/")
	}

	// 3. If pattern has no glob wildcards, do a literal check.
	if !strings.ContainsAny(pattern, "*?[") {
		// Check if any segment equals the pattern.
		parts := strings.Split(pathStr, "/")
		for _, part := range parts {
			if part == pattern {
				return true
			}
		}
		// Also check if the entire path starts with the pattern followed by a slash.
		return strings.HasPrefix(pathStr, pattern+"/")
	}

	// 4. For wildcard patterns:
	if strings.Contains(pattern, "/") {
		// Match against the full normalized path.
		match, err := filepath.Match(pattern, pathStr)
		return err == nil && match
	}
	// Otherwise, match against the base name.
	base := filepath.Base(pathStr)
	match, err := filepath.Match(pattern, base)
	return err == nil && match
}

// isAllowed determines whether a file should be processed, matching
// against the full normalized (slash-separated) path so that patterns
// like "internal", "**.go", and anchored patterns such as "/README.md"
// (which only match files in the root folder) work as expected.
//
//nolint:gocognit // TODO: refactor this at some point.
func (a *App) isAllowed(filePath string, isSourceMode bool) bool {
	// Immediately ignore the destination file (if OutFilename is set)
	if a.OutFilename != "" && filepath.Base(a.OutFilename) == filepath.Base(filePath) {
		return false
	}

	// Select mode-specific settings.
	var showHidden bool
	var modeInclude, modeExclude string
	if isSourceMode {
		showHidden = a.SourceShowHidden
		modeInclude = a.SourceInclude
		modeExclude = a.SourceExclude
	} else {
		showHidden = a.TreeShowHidden
		modeInclude = a.TreeInclude
		modeExclude = a.TreeExclude
	}

	// Disallow hidden files/folders if not allowed.
	if !showHidden {
		if isHidden(filePath) || isHidden(filepath.Base(filePath)) {
			return false
		}
	}

	// Normalize the file path to use forward slashes.
	normalizedPath := filepath.ToSlash(filePath)

	// 1. Apply Core Ignores unless disabled.
	if !a.NoCoreIgnores {
		for _, pattern := range CoreIgnores {
			if matchPattern(pattern, normalizedPath) {
				return false
			}
		}
		if isSourceMode {
			for _, pattern := range CoreSourceIgnores {
				if matchPattern(pattern, normalizedPath) {
					return false
				}
			}
		}
	}

	// 2. Apply additional excludes from .aictxignore (if any).
	for _, pattern := range a.AictxIgnore {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}
		if matchPattern(pattern, normalizedPath) {
			return false
		}
	}

	// 3. Determine effective include.
	effectiveInclude := cmp.Or(modeInclude, a.Include, "**")
	var matched bool
	for _, pattern := range strings.Split(effectiveInclude, ",") {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}
		if matchPattern(pattern, normalizedPath) {
			matched = true
			break
		}
	}
	if !matched {
		return false
	}

	// 4. Determine effective exclude.
	effectiveExclude := cmp.Or(modeExclude, a.Exclude)
	for _, pattern := range strings.Split(effectiveExclude, ",") {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}
		if matchPattern(pattern, normalizedPath) {
			return false
		}
	}

	return true
}

// hasAllowed checks recursively whether a given directory (or file) contains any allowed content.
func (a *App) hasAllowed(fsys billy.Filesystem, path string, isSourceMode bool) (bool, error) {
	info, err := fsys.Stat(path)
	if err != nil {
		return false, err
	}
	// For files, simply check if allowed.
	if !info.IsDir() {
		return a.isAllowed(path, isSourceMode), nil
	}
	// For directories, check each child.
	entries, err := fsys.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())

		var allowed bool
		if allowed, err = a.hasAllowed(fsys, childPath, isSourceMode); err != nil {
			return false, err
		}
		if allowed {
			return true, nil
		}
	}
	return false, nil
}

// TreeNode is a simple structure for building the filtered directory tree.
type TreeNode struct {
	Name     string
	Path     string
	IsDir    bool
	Size     int64 // File size in bytes (only used if IsDir==false)
	Children []*TreeNode
	IsBinary bool
}

type summary struct {
	fileCount int
	totalSize int64
	maxSize   int64
}

// summary recursively traverses the tree and returns:
//   - fileCount: number of files in the tree,
//   - totalSize: sum of sizes (in bytes) of all files,
//   - maxSize: maximum file size (in bytes) among all files.
func (node *TreeNode) summary() summary {
	if !node.IsDir {
		// This is a file.
		return summary{
			1, node.Size, node.Size,
		}
	}

	// For directories, iterate through children.
	var s summary
	for _, child := range node.Children {
		childS := child.summary()
		s.fileCount += childS.fileCount
		s.totalSize += childS.totalSize
		if childS.maxSize > s.maxSize {
			s.maxSize = childS.maxSize
		}
	}

	return s
}

var ErrFilterSkipped = errors.New("filter skipped")

// filterTree recursively builds a tree structure of allowed nodes.
// It applies the same isAllowed/hasAllowed logic but does not do any printing.
func (a *App) filterTree(ctx context.Context, fsys billy.Filesystem, root string) (*TreeNode, error) {
	// Check for cancellation.
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	info, err := fsys.Stat(root)
	if err != nil {
		return nil, err
	}

	node := &TreeNode{
		Name:  info.Name(),
		Path:  root,
		IsDir: info.IsDir(),
	}

	// If this is a file, include it only if allowed.
	if !info.IsDir() {
		if a.isAllowed(root, false) {
			node.Size = info.Size()
			return node, nil
		}
		return nil, ErrFilterSkipped
	}

	// For directories, process each child.
	entries, err := fsys.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		childPath := filepath.Join(root, entry.Name())
		if entry.IsDir() { //nolint: nestif // we're OK with this
			// Only include directories if they (or any of their descendants) are allowed.
			ok, err := a.hasAllowed(fsys, childPath, false)
			if err != nil {
				return nil, err
			}
			if !ok {
				continue
			}
			childNode, err := a.filterTree(ctx, fsys, childPath)
			if err != nil {
				return nil, err
			}
			if childNode != nil {
				node.Children = append(node.Children, childNode)
			}
		} else if a.isAllowed(childPath, false) {
			childInfo, err := fsys.Stat(childPath)
			if err != nil {
				return nil, err
			}
			// Read file content to determine if binary.
			data, err := fsutils.ReadAll(fsys, childPath)
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, &TreeNode{
				Name:     entry.Name(),
				Path:     childPath,
				IsDir:    false,
				Size:     childInfo.Size(),
				IsBinary: isBinary(data),
			})
		}
	}

	return node, nil
}

// printTree recursively prints the node and its children.
func (node *TreeNode) printTree(prefix string, w io.Writer) {
	if prefix == "" {
		fmt.Fprintln(w, node.Name)
	}

	childCount := len(node.Children)
	for i, child := range node.Children {
		childName := child.Name
		if !child.IsDir && child.IsBinary {
			childName += " *"
		}
		connector := "├── "
		newPrefix := prefix + "│   "
		if i == childCount-1 {
			connector = "└── "
			newPrefix = prefix + "    "
		}
		fmt.Fprintln(w, prefix+connector+childName)
		if child.IsDir && len(child.Children) > 0 {
			child.printTree(newPrefix, w)
		}
	}
}

// printSourceFiles recursively traverses the tree and prints file content for each file.
// totalFiles is the total number of files (from the summary) and fileCounter is a pointer
// to a running counter.
func (a *App) printSourceFiles(ctx context.Context, fs billy.Filesystem,
	node *TreeNode, w io.Writer, totalFiles int, fileCounter *int,
) error {
	// Check cancellation.
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// If it's a file, print its content.
	if !node.IsDir {
		*fileCounter++ // increment the counter
		// *fileCounter = *fileCounter // just to emphasize its updated
		data, err := fsutils.ReadAll(fs, node.Path)
		if err != nil {
			log.Printf("Error reading file '%s': %s", node.Path, err)
			return nil
		}
		// Skip binary files.
		if isBinary(data) {
			return nil
		}

		// Write the header including file number.
		if _, err = w.Write(fileHeader(node, *fileCounter, totalFiles)); err != nil {
			log.Printf("Error writing header for '%s': %s", node.Path, err)
		}
		if _, err = w.Write(data); err != nil {
			log.Printf("Error writing content from '%s': %s", node.Path, err)
		}
		fmt.Fprintln(w) // Separate files with a blank line.
		return nil
	}

	// If it's a directory, process its children.
	for _, child := range node.Children {
		if err := a.printSourceFiles(ctx, fs, child, w, totalFiles, fileCounter); err != nil {
			return err
		}
	}
	return nil
}

// printSourceFilesRaw recursively traverses the tree and prints the content of each file
// without any headers or summary information.
func (a *App) printSourceFilesRaw(ctx context.Context, fs billy.Filesystem, node *TreeNode, w io.Writer) error {
	// Check cancellation.
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// If it's a file, print its content.
	if !node.IsDir {
		data, err := fsutils.ReadAll(fs, node.Path)
		if err != nil {
			log.Printf("Error reading file '%s': %s", node.Path, err)
			return nil
		}
		// Skip binary files.
		if isBinary(data) {
			return nil
		}
		if _, err = w.Write(data); err != nil {
			log.Printf("Error writing content from '%s': %s", node.Path, err)
		}
		fmt.Fprintln(w) // Separate files with a blank line.
		return nil
	}

	// If it's a directory, process its children.
	for _, child := range node.Children {
		if err := a.printSourceFilesRaw(ctx, fs, child, w); err != nil {
			return err
		}
	}
	return nil
}

// loadDotIgnoreFromFS tries to open and read the .aictxignore file
// from the given root directory using the provided billy.Filesystem.
// It returns a slice of non-empty, non-comment lines.
func loadDotIgnoreFromFS(fsys billy.Filesystem, dotIgnoreFile string, root string) ([]string, error) {
	ignorePath := filepath.Join(root, dotIgnoreFile)
	f, err := fsys.Open(ignorePath)
	if err != nil {
		// Not finding the file is not an error.
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var patterns []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		patterns = append(patterns, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return patterns, nil
}
