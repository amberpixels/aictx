// Package fsutils for filesystem (billy.Filesystem) utils
package fsutils

import (
	"io"
	"os"
	"path"

	"github.com/go-git/go-billy/v5"
)

// kb as bytes in kilobytes.
const kb = 1024

// FileSizeInMb returns the size of the file in megabytes.
func FileSizeInMb(info os.FileInfo) float64 {
	return float64(info.Size()) / (kb * kb)
}

// ReadAll reads the contents of the file at path.
func ReadAll(fs billy.Filesystem, path string) ([]byte, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

// Walk traverses the billy.Filesystem starting at root and calls fn for each file or directory.
func Walk(fsys billy.Filesystem, root string, fn func(path string, info os.FileInfo, err error) error) error {
	// Stat the root to get file info.
	info, err := fsys.Stat(root)
	if err != nil {
		return fn(root, nil, err)
	}
	// Process the current entry.
	if err := fn(root, info, nil); err != nil {
		return err
	}

	// If it's a directory, use ReadDir to list entries.
	if info.IsDir() {
		entries, err := fsys.ReadDir(root)
		if err != nil {
			return fn(root, info, err)
		}
		for _, entry := range entries {
			// Use the billy util package's Join (or implement your own) to handle paths correctly.
			childPath := path.Join(root, entry.Name())
			if err := Walk(fsys, childPath, fn); err != nil {
				return err
			}
		}
	}

	return nil
}
