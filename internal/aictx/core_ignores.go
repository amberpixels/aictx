package aictx

import (
	"fmt"
	"io"
)

// CoreIgnores lists patterns that should be excluded for both tree and source modes.
//
//nolint:gochecknoglobals // Hardcoded patterns.
var CoreIgnores = []string{
	"Thumbs.db",
	"__pycache__",
}

// CoreSourceIgnores lists patterns that should be excluded only for source mode (but are ok in tree mode).
//
//nolint:gochecknoglobals // Hardcoded patterns.
var CoreSourceIgnores = []string{
	"go.sum",
	"vendor",
	"node_modules",
	"package-lock.json",
	"yarn.lock",
	"npm-debug.log",
	"__pycache__",
	"*.pyc",
	"*.pyo",
	"*.pyd",
	"*.egg-info",
	"build",
	"dist",
	"*.class",
	"target",
	"*.jar",
	"*.war",
	"*.ear",
	"Gemfile.lock",
	".bundle",
	"*.o",
	"*.obj",
	"*.exe",
	"*.so",
	"*.dSYM",
	"CMakeCache.txt",
	"CMakeFiles",
	"bin",
	"obj",
}

// PrintCoreIgnores writes the core exclude patterns to the provided writer.
func PrintCoreIgnores(w io.Writer) {
	fmt.Fprintln(w, "Core Ignores (Both Tree and Source Modes):")
	for _, pattern := range CoreIgnores {
		fmt.Fprintf(w, "  %s\n", pattern)
	}
	fmt.Fprintln(w, "\nCore Ignores (Only for Source Mode):")
	for _, pattern := range CoreSourceIgnores {
		fmt.Fprintf(w, "  %s\n", pattern)
	}
}
