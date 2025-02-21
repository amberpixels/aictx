package aictx

import (
	"fmt"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

// ReadGit clones the given Git repository URL into an in-memory FS.
// If branch is non-empty, it will clone only that branch.
func ReadGit(repoURL, branch string) (billy.Filesystem, error) {
	storer := memory.NewStorage()
	billyFS := memfs.New()

	cloneOpts := &git.CloneOptions{
		URL: repoURL,
	}
	if branch != "" {
		cloneOpts.ReferenceName = plumbing.NewBranchReferenceName(branch)
		cloneOpts.SingleBranch = true
	}

	_, err := git.Clone(storer, billyFS, cloneOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to clone git repository: %w", err)
	}

	return billyFS, nil
}

// ValidateGitRepoName parses the repository shorthand and optional branch information.
// It supports both HTTPS and SSH URLs. The logic is as follows:
//   - If the input starts with "git@", it is treated as an SSH URL. The branch is extracted
//     from the last '@' (if present).
//   - Otherwise, if the input contains an '@', the part after the last '@' is the branch.
//   - For HTTPS URLs, if the input does not contain a slash, it is considered invalid.
//   - If the input does not start with "github.com/", it is assumed to be from GitHub and
//     "github.com/" is prepended.
//   - Finally, the URL is built as "https://<repo>.git" if not already ending with ".git".
func ValidateGitRepoName(repo string) (string, string, error) {
	repo = strings.TrimSpace(repo)
	if repo == "." || repo == "" {
		return "", "", fmt.Errorf("'%s' is not a valid repository name", repo)
	}

	var branch string

	// Check if it's an SSH URL.
	if strings.HasPrefix(repo, "git@") {
		// Find the last '@'. The first '@' is part of the SSH URL.
		firstAt := strings.Index(repo, "@")
		lastAt := strings.LastIndex(repo, "@")
		if lastAt > firstAt {
			branch = strings.TrimSpace(repo[lastAt+1:])
			repo = strings.TrimSpace(repo[:lastAt])
		}
		// Return the SSH URL as-is.
		return repo, branch, nil
	}

	// For HTTPS style, check for branch information by splitting on the last '@'.
	if strings.Contains(repo, "@") {
		lastAt := strings.LastIndex(repo, "@")
		branch = strings.TrimSpace(repo[lastAt+1:])
		repo = strings.TrimSpace(repo[:lastAt])
	}

	// Validate that the repo has a slash.
	if !strings.Contains(repo, "/") {
		return "", "", fmt.Errorf("invalid repository format: %s", repo)
	}

	repo = strings.TrimPrefix(repo, "https://")
	repo = strings.TrimPrefix(repo, "http://")
	repo = strings.TrimSuffix(repo, ".git")

	// Prepend "github.com/" if not present.
	if !strings.HasPrefix(repo, "github.com/") {
		repo = "github.com/" + repo
	}

	return "https://" + repo + ".git", branch, nil
}
