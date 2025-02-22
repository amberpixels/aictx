<h1 align="center">
    AI Ctx
</h1>

🤖 **aictx** is a powerful CLI tool that transforms your entire codebase into a single, AI-ready text file.
Perfect for feeding context to large language models like ChatGPT, Claude, and more, it supports both local directories and remote Git repositories.
With flexible filtering options—including global and mode-specific include/exclude patterns—you control exactly which files are included,
creating a seamless narrative of your project.

## Features

- **📁 Local Directory Support**:
  Processes given local directory and its subdirectories (if needed, including hidden content as well).

- **🌐 Git Repository Support**:
  Accepts a Git repository shorthand (e.g., `github.com/amberpixels/aictx`) to process its codebase.

- **🗃️ Tree Mode**:
  Displays a structured tree view of the input with a summary (total file count, cumulative size, and largest file size).

- **📜 Source Mode**:
  Outputs the contents of allowed source files with informative headers including file number and size.
  Files exceeding a configurable size threshold are skipped.

- **🛠️ Flexible Filtering**:
  - Apply global and mode-specific glob patterns (supports comma-separated lists) to include or exclude files.
  - Automatically respects `.gitignore` if exists (can be disabled).
  - Automatically ignores common unwanted files (e.g., `vendor`, `Thumbs.db`, `__pycache__`, `node_modules`) (can be disabled).
  - Automatically ignores hidden and/or binary files (can be disabled).
  - Reads additional ignore patterns from a `.aictxignore` file in the input directory.
  - Automatically excludes the output file (default `output.txt` or a user-specified file) from processing.

## Installation
Ensure you have [Go](https://golang.org/) installed.

```bash
go install github.com/amberpixels/aictx@latest
```

## Usage and Options

```bash
Usage: aictx [<input-path>] [flags]

Arguments:
  [<input-path>]    Input directory (or git repo URL) to process

Flags:
  -h, --help                    Show context-sensitive help.
  -l, --local                   Treat inputPath arg as a local directory.
                                If inputPath is '.' it is automatically makes
                                local=true.
  -i, --include=""              Global include glob pattern (supports
                                comma-separated list)
  -x, --exclude=""              Global exclude glob pattern (supports
                                comma-separated list)
      --source.disabled         Disable source mode
      --source.include=""       Include glob pattern specific for source mode.
                                Global include is used if not specified.
      --source.exclude=""       Exclude glob pattern specific for source mode.
                                Global exclude is used if not specified.
      --source.threshold=0.1    Exclude sources for files >= threshold (Mb)
      --source.show-hidden      Show hidden files in source mode
      --tree.disabled           Disable tree mode
      --tree.include=""         Include glob pattern specific for tree mode.
                                Global include is used if not specified.
      --tree.exclude=""         Exclude glob pattern specific for tree mode.
                                Global exclude is used if not specified.
      --tree.show-hidden        Show hidden files in tree mode
  -o, --out="output.txt"        Output destination file ("stdout" for stdout)
  -v, --verbose                 Verbose mode
  -r, --raw                     Concatenate file contents in raw mode without
                                headers or summary
  -L, --list-core-ignores       List core ignore patterns and exit
      --no-core-ignores         Disable core ignore patterns
      --no-git-ignore           Disable respecting .gitignore file

```

## Examples

- **Display Both Tree and Source views for a Current Local Directory**

  ```bash
  aictx
  ```

- **Display Both Tree and Source views for a Github repo**

  ```bash
  aictx amberpixels/aictx

  # supported any type of github repo mention.
  # as well as gitlab's repos.
  aictx github.com/amberpixels/aictx
  ```

- **Include specific globs (for both Tree & Source mode) **

  ```bash
  aictx --i "*.go,go.md"
  ```
- **Process a Git Repository with a Custom Size Threshold**

  ```bash
  aictx github.com/amberpixels/aictx --source.threshold 0.2 --source.include="*.js" --out=stdout
  ```
- **List Core Ignore Patterns**

  ```bash
  aictx --list-core-ignores
  ```

- **Include everything: hidden files (and core-ignored)**

  ```bash
  aictx --no-core-ignores --show-hidden
  ```

## Project Structure

~~~text
Project Tree [9 files, 16.05 MB total, max 16.01 MB] (* - for binary files)
.
├── build
│   └── aictx *
├── cmd
│   └── aictx
│       └── main.go
├── internal
│   ├── aictx
│   │   ├── app.go
│   │   ├── core_ignores.go
│   │   ├── git.go
│   │   ├── git_test.go
│   │   └── helpers.go
│   └── fsutils
│       └── fsutils.go
└── tools
    └── readmegen.go
~~~

## Contributing
Contributions are welcome! Please fork the repository and submit pull requests for any enhancements, bug fixes, or suggestions.

**Please Note:** The application is stable and ready-to-be used, but it is not currently covered by tests.
As a result, accepting new features may take longer than expected until we have at least some basic test coverage
in place. We appreciate your patience and contributions as we work towards enhancing
the project's quality and reliability.

## License
This project is licensed under the [MIT License](LICENSE).

Happy coding! :heart:

<!--
# AI Codebase Overview

The following part of readme.md is generated specifically for AI agents to get
full context of the codebase in a single place.
It contains the entire source code,
making additional repo scanning unnecessary.

The file ends with the delimiter: `--- END OF CODEBASE ---`
AI agents should rely solely on this file for analysis.

--- START OF CODEBASE ---
Project Tree [15 files, 16.06 MB total, max 15.99 MB] (* - for binary files)
.
├── LICENSE
├── Makefile
├── README.md
├── build
│   └── aictx *
├── cmd
│   └── aictx
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── aictx
│   │   ├── app.go
│   │   ├── core_ignores.go
│   │   ├── git.go
│   │   ├── git_test.go
│   │   └── helpers.go
│   └── fsutils
│       └── fsutils.go
└── tools
    ├── readme.template.md
    └── readmegen.go

Project Source [13 files, 50.97 KB total, max 22.61 KB]
==========================[01/13]===========================
File: LICENSE
Size: 1.05 KB
------------------------------------------------------------
MIT License

Copyright (c) 2025 amberpixels (Eugene M)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

==========================[02/13]===========================
File: Makefile
Size: 1.11 KB
------------------------------------------------------------
# Variables
GOLANGCI_LINT := $(shell which golangci-lint)

BUILD_DIR := build
CMD_DIR = ./cmd/aictx
MAIN_FILE := $(CMD_DIR)/main.go

BINARY_NAME := aictx
INSTALL_DIR := $(shell go env GOPATH)/bin

# Default target
all: build

# Build the binary
build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# Run the binary
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Tidy: format and vet the code
tidy:
	@go fmt $$(go list ./...)
	@go vet $$(go list ./...)
	@go mod tidy

# Install golangci-lint only if it's not already installed
lint-install:
	@if ! [ -x "$(GOLANGCI_LINT)" ]; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi

# Lint the code using golangci-lint
# todo reuse var if possible
lint: lint-install
	$(shell which golangci-lint) run

# Install the binary globally with aliases
install:
	@go install $(CMD_DIR)

# Uninstall the binary and remove the alias
uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

readme: build tools/readme.template.md
	@go run tools/readmegen.go

# Phony targets
.PHONY: all build run tidy lint-install lint install uninstall readme

==========================[03/13]===========================
File: README.md
Size: 5.64 KB
------------------------------------------------------------
<h1 align="center">
    AI Ctx
</h1>

🤖 **aictx** is a powerful CLI tool that transforms your entire codebase into a single, AI-ready text file.<br>
It's perfect for feeding context to LLMs like ChatGPT, Claude, and more. It supports both local directories and remote Git repositories.<br>
With flexible filtering options — including global and mode-specific include/exclude patterns — you control exactly which files are included,
creating a seamless narrative of your project.

## Features

- **📁 Local Directory Support**:
  Processes given local directory and its subdirectories (if needed, including hidden content as well).

- **🌐 Git Repository Support**:
  Accepts a Git repository shorthand (e.g., `github.com/amberpixels/aictx`) to process its codebase.

- **🗃️ Tree Mode**:
  Displays a structured tree view of the input with a summary (total file count, cumulative size, and largest file size).

- **📜 Source Mode**:
  Outputs the contents of allowed source files with informative headers including file number and size.
  Files exceeding a configurable size threshold are skipped.

- **🛠️ Flexible Filtering**:
  - Apply global and mode-specific glob patterns (supports comma-separated lists) to include or exclude files.
  - Automatically respects `.gitignore` if exists (can be disabled).
  - Automatically ignores common unwanted files (e.g., `vendor`, `Thumbs.db`, `__pycache__`, `node_modules`) (can be disabled).
  - Automatically ignores hidden and/or binary files (can be disabled).
  - Reads additional ignore patterns from a `.aictxignore` file in the input directory.
  - Automatically excludes the output file (default `output.txt` or a user-specified file) from processing.

## Installation
Ensure you have [Go](https://golang.org/) installed.

```bash
go install github.com/amberpixels/aictx@latest
```

## Usage and Options

```bash
Usage: aictx [<input-path>] [flags]

Arguments:
  [<input-path>]    Input directory (or git repo URL) to process

Flags:
  -h, --help                    Show context-sensitive help.
  -l, --local                   Treat inputPath arg as a local directory.
                                If inputPath is '.' it is automatically makes
                                local=true.
  -i, --include=""              Global include glob pattern (supports
                                comma-separated list)
  -x, --exclude=""              Global exclude glob pattern (supports
                                comma-separated list)
      --source.disabled         Disable source mode
      --source.include=""       Include glob pattern specific for source mode.
                                Global include is used if not specified.
      --source.exclude=""       Exclude glob pattern specific for source mode.
                                Global exclude is used if not specified.
      --source.threshold=0.1    Exclude sources for files >= threshold (Mb)
      --source.show-hidden      Show hidden files in source mode
      --tree.disabled           Disable tree mode
      --tree.include=""         Include glob pattern specific for tree mode.
                                Global include is used if not specified.
      --tree.exclude=""         Exclude glob pattern specific for tree mode.
                                Global exclude is used if not specified.
      --tree.show-hidden        Show hidden files in tree mode
  -o, --out="output.txt"        Output destination file ("stdout" for stdout)
  -v, --verbose                 Verbose mode
  -r, --raw                     Concatenate file contents in raw mode without
                                headers or summary
  -L, --list-core-ignores       List core ignore patterns and exit
      --no-core-ignores         Disable core ignore patterns
      --no-git-ignore           Disable respecting .gitignore file

```

## Examples

- **Display Both Tree and Source views for a Current Local Directory**

  ```bash
  aictx
  ```

- **Display Both Tree and Source views for a Github repo**

  ```bash
  aictx amberpixels/aictx

  # supported any type of github repo mention.
  # as well as gitlab's repos.
  aictx github.com/amberpixels/aictx
  ```

- **Include specific globs (for both Tree & Source mode) **

  ```bash
  aictx --i "*.go,go.md"
  ```
- **Process a Git Repository with a Custom Size Threshold**

  ```bash
  aictx github.com/amberpixels/aictx --source.threshold 0.2 --source.include="*.js" --out=stdout
  ```
- **List Core Ignore Patterns**

  ```bash
  aictx --list-core-ignores
  ```

- **Include everything: hidden files (and core-ignored)**

  ```bash
  aictx --no-core-ignores --show-hidden
  ```

## Project Structure

~~~text
Project Tree [9 files, 16.05 MB total, max 16.01 MB] (* - for binary files)
.
├── build
│   └── aictx *
├── cmd
│   └── aictx
│       └── main.go
├── internal
│   ├── aictx
│   │   ├── app.go
│   │   ├── core_ignores.go
│   │   ├── git.go
│   │   ├── git_test.go
│   │   └── helpers.go
│   └── fsutils
│       └── fsutils.go
└── tools
    └── readmegen.go
~~~

## Contributing
Contributions are welcome! Please fork the repository and submit pull requests for any enhancements, bug fixes, or suggestions.

**Please Note:** The application is stable and ready-to-be used, but it is not currently covered by tests.
As a result, accepting new features may take longer than expected until we have at least some basic test coverage
in place. We appreciate your patience and contributions as we work towards enhancing
the project's quality and reliability.

## License
This project is licensed under the [MIT License](LICENSE).

Happy coding! :heart:

==========================[04/13]===========================
File: cmd/aictx/main.go
Size: 3.80 KB
------------------------------------------------------------
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/charmbracelet/log"

	"github.com/alecthomas/kong"
	"github.com/amberpixels/aictx/internal/aictx"
)

type CliParams struct {
	InputPath string `arg:"" default:"." help:"Input directory (or git repo URL) to process"`
	Local     bool   `short:"l" help:"Treat inputPath arg as a local directory. If inputPath is '.' it is automatically makes local=true." default:"false"` //nolint:lll

	// Global include/exclude patterns will be applied to both source/tree modes unless overridden.
	Include string `short:"i" help:"Global include glob pattern (supports comma-separated list)" default:""`
	Exclude string `short:"x" help:"Global exclude glob pattern (supports comma-separated list)" default:""`

	Source struct {
		Disabled   bool    `help:"Disable source mode" default:"false"`
		Include    string  `help:"Include glob pattern specific for source mode. Global include is used if not specified." default:""` //nolint:lll
		Exclude    string  `help:"Exclude glob pattern specific for source mode. Global exclude is used if not specified." default:""` //nolint:lll
		Threshold  float64 `help:"Exclude sources for files >= threshold (Mb)" default:"0.1"`
		ShowHidden bool    `help:"Show hidden files in source mode" default:"false"`
	} `embed:"" prefix:"source."`

	Tree struct {
		Disabled   bool   `help:"Disable tree mode" default:"false"`
		Include    string `help:"Include glob pattern specific for tree mode. Global include is used if not specified." default:""` //nolint:lll
		Exclude    string `help:"Exclude glob pattern specific for tree mode. Global exclude is used if not specified." default:""` //nolint:lll
		ShowHidden bool   `help:"Show hidden files in tree mode" default:"false"`
	} `embed:"" prefix:"tree."`

	Out string `short:"o" help:"Output destination file (\"stdout\" for stdout)" default:"output.txt"`

	Verbose         bool `short:"v" help:"Verbose mode" default:"false"`
	Raw             bool `short:"r" help:"Concatenate file contents in raw mode without headers or summary" default:"false"` //nolint:lll
	ListCoreIgnores bool `short:"L" help:"List core ignore patterns and exit" default:"false"`
	NoCoreIgnores   bool `help:"Disable core ignore patterns" default:"false"`
	NoGitIgnore     bool `help:"Disable respecting .gitignore file" default:"false"`
}

func main() {
	var cli CliParams
	// Parse CLI arguments using Kong.
	kctx := kong.Parse(&cli)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: "aictx 🤖 ",
	})

	// If --list-core-excludes is set, show the core excludes and exit.
	if cli.ListCoreIgnores {
		aictx.PrintCoreIgnores(os.Stdout)
		return
	}

	app := &aictx.App{
		Lgr: logger,

		InputPath: cli.InputPath,
		Local:     cli.Local,

		// Global include/exclude patterns.
		Include: cli.Include,
		Exclude: cli.Exclude,

		// Source mode specific configuration.
		SourceEnabled:    !cli.Source.Disabled,
		SourceInclude:    cli.Source.Include,
		SourceExclude:    cli.Source.Exclude,
		SourceThreshold:  cli.Source.Threshold,
		SourceShowHidden: cli.Source.ShowHidden,

		TreeEnabled:    !cli.Tree.Disabled,
		TreeInclude:    cli.Tree.Include,
		TreeExclude:    cli.Tree.Exclude,
		TreeShowHidden: cli.Tree.ShowHidden,

		Raw:     cli.Raw,
		Verbose: cli.Verbose,

		NoCoreIgnores: cli.NoCoreIgnores,
		NoGitIgnore:   cli.NoGitIgnore,
	}

	if cli.Out == "stdout" || cli.Out == "std" || cli.Out == "-" {
		app.Out = os.Stdout
		// If we output to stdout we need to disable the verbose mode
		app.Verbose = false
	} else {
		f, err := os.Create(cli.Out)
		if err != nil {
			log.Printf("Error creating output file '%s': %v", cli.Out, err)
			return
		}
		defer f.Close()
		app.Out = f
		app.OutFilename = cli.Out
	}

	err := app.Run(ctx)
	kctx.FatalIfErrorf(err)
}

==========================[05/13]===========================
File: go.mod
Size: 1.95 KB
------------------------------------------------------------
module github.com/amberpixels/aictx

go 1.23.5

require (
	github.com/alecthomas/kong v1.8.1
	github.com/charmbracelet/log v0.4.0
	github.com/go-git/go-billy/v5 v5.6.2
	github.com/go-git/go-git/v5 v5.13.2
	github.com/stretchr/testify v1.10.0
	github.com/yarlson/pin v0.9.0
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v1.1.5 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/lipgloss v0.10.0 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/cyphar/filepath-securejoin v0.3.6 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/pjbgf/sha1cd v0.3.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 // indirect
	github.com/skeema/knownhosts v1.3.0 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/tools v0.23.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

==========================[06/13]===========================
File: internal/aictx/app.go
Size: 22.61 KB
------------------------------------------------------------
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

==========================[07/13]===========================
File: internal/aictx/core_ignores.go
Size: 1.16 KB
------------------------------------------------------------
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

==========================[08/13]===========================
File: internal/aictx/git.go
Size: 2.75 KB
------------------------------------------------------------
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

==========================[09/13]===========================
File: internal/aictx/git_test.go
Size: 2.35 KB
------------------------------------------------------------
package aictx_test

import (
	"testing"

	"github.com/amberpixels/aictx/internal/aictx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateGitRepoName(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedURL    string
		expectedBranch string
		expectError    bool
	}{
		{
			name:        "empty input",
			input:       "",
			expectError: true,
		},
		{
			name:        "dot input",
			input:       ".",
			expectError: true,
		},
		{
			name:        "HTTPS without prefix",
			input:       "foo/bar",
			expectedURL: "https://github.com/foo/bar.git",
		},
		{
			name:           "HTTPS with branch",
			input:          "foo/bar@dev",
			expectedURL:    "https://github.com/foo/bar.git",
			expectedBranch: "dev",
		},
		{
			name:        "HTTPS with github.com prefix",
			input:       "github.com/foo/bar",
			expectedURL: "https://github.com/foo/bar.git",
		},
		{
			name:           "HTTPS with github.com prefix and branch",
			input:          "github.com/foo/bar@feature",
			expectedURL:    "https://github.com/foo/bar.git",
			expectedBranch: "feature",
		},
		{
			name:        "HTTPS with https:// prefix",
			input:       "https://github.com/foo/bar.git",
			expectedURL: "https://github.com/foo/bar.git",
		},
		{
			name:        "HTTPS with http:// prefix",
			input:       "http://github.com/foo/bar",
			expectedURL: "https://github.com/foo/bar.git",
		},
		{
			name:        "Invalid format without slash",
			input:       "foobar",
			expectError: true,
		},
		{
			name:        "SSH URL without branch",
			input:       "git@github.com:user/repo.git",
			expectedURL: "git@github.com:user/repo.git",
		},
		{
			name:           "SSH URL with branch",
			input:          "git@github.com:user/repo.git@feature-branch",
			expectedURL:    "git@github.com:user/repo.git",
			expectedBranch: "feature-branch",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, branch, err := aictx.ValidateGitRepoName(tc.input)
			if tc.expectError {
				require.Error(t, err, "expected error for input: %s", tc.input)
			} else {
				require.NoError(t, err, "unexpected error for input: %s", tc.input)
				assert.Equal(t, tc.expectedURL, url, "URL mismatch for input: %s", tc.input)
				assert.Equal(t, tc.expectedBranch, branch, "branch mismatch for input: %s", tc.input)
			}
		})
	}
}

==========================[10/13]===========================
File: internal/aictx/helpers.go
Size: 2.22 KB
------------------------------------------------------------
package aictx

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	// KB is number of bytes in Kilobyte.
	KB = 1024
	// MB is number of bytes in Megabyte.
	MB = 1024 * KB
)

// fileHeader renders a header for each file.
// fileHeader renders a header for each file.
// It now includes a file counter (e.g. "[1/6]" or "[01/12]") inserted into a 60-char line.
func fileHeader(node *TreeNode, fileNum, totalFiles int) []byte {
	const totalLen = 60 // total characters (without the newline)
	var buf bytes.Buffer

	// Determine padding width: if totalFiles is single digit then width=1,
	// if <100 then width=2, if <1000 then width=3, etc.
	width := len(strconv.Itoa(totalFiles))
	// Create the number info string with padded file number.
	numInfo := fmt.Sprintf("[%0*d/%d]", width, fileNum, totalFiles)
	// Calculate remaining space and split evenly on left/right.
	rem := totalLen - len(numInfo)
	left := rem / 2 //nolint: mnd // 2 for half
	right := rem - left
	headerLine := strings.Repeat("=", left) + numInfo + strings.Repeat("=", right) + "\n"

	buf.WriteString(headerLine)
	buf.WriteString(fmt.Sprintf("File: %s\n", node.Path))
	if node.Size > 0 {
		buf.WriteString(fmt.Sprintf("Size: %s\n", formatSize(node.Size)))
	}
	buf.WriteString(strings.Repeat("-", totalLen) + "\n")
	return buf.Bytes()
}

// isHidden returns true if the provided file or folder name starts with a dot.
func isHidden(name string) bool {
	return len(name) > 0 && name[0] == '.'
}

// isBinary returns true if the provided data appears to be binary.
func isBinary(data []byte) bool {
	if len(data) == 0 {
		return false // empty files are considered text
	}
	for _, b := range data {
		if b == 0 {
			return true
		}
	}
	return !utf8.Valid(data)
}

// formatSize converts bytes to a human-friendly string.
func formatSize(bytes int64) string {
	if bytes >= MB {
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	} else if bytes >= KB {
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	}
	return fmt.Sprintf("%d B", bytes)
}

// exceedsThreshold returns true if the file size (in bytes) exceeds the threshold (in MB).
func exceedsThreshold(sizeBytes int64, thresholdMb float64) bool {
	mb := float64(sizeBytes) / (KB * KB)
	return mb > thresholdMb
}

==========================[11/13]===========================
File: internal/fsutils/fsutils.go
Size: 1.37 KB
------------------------------------------------------------
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

==========================[12/13]===========================
File: tools/readme.template.md
Size: 3.72 KB
------------------------------------------------------------
<h1 align="center">
    AI Ctx
</h1>

🤖 **aictx** is a powerful CLI tool that transforms your entire codebase into a single, AI-ready text file.
Perfect for feeding context to large language models like ChatGPT, Claude, and more, it supports both local directories and remote Git repositories.
With flexible filtering options—including global and mode-specific include/exclude patterns—you control exactly which files are included,
creating a seamless narrative of your project.

## Features

- **📁 Local Directory Support**:
  Processes given local directory and its subdirectories (if needed, including hidden content as well).

- **🌐 Git Repository Support**:
  Accepts a Git repository shorthand (e.g., `github.com/amberpixels/aictx`) to process its codebase.

- **🗃️ Tree Mode**:
  Displays a structured tree view of the input with a summary (total file count, cumulative size, and largest file size).

- **📜 Source Mode**:
  Outputs the contents of allowed source files with informative headers including file number and size.
  Files exceeding a configurable size threshold are skipped.

- **🛠️ Flexible Filtering**:
  - Apply global and mode-specific glob patterns (supports comma-separated lists) to include or exclude files.
  - Automatically respects `.gitignore` if exists (can be disabled).
  - Automatically ignores common unwanted files (e.g., `vendor`, `Thumbs.db`, `__pycache__`, `node_modules`) (can be disabled).
  - Automatically ignores hidden and/or binary files (can be disabled).
  - Reads additional ignore patterns from a `.aictxignore` file in the input directory.
  - Automatically excludes the output file (default `output.txt` or a user-specified file) from processing.

## Installation
Ensure you have [Go](https://golang.org/) installed.

```bash
go install github.com/amberpixels/aictx@latest
```

## Usage and Options

```bash
{{ .Usage }}
```

## Examples

- **Display Both Tree and Source views for a Current Local Directory**

  ```bash
  aictx
  ```

- **Display Both Tree and Source views for a Github repo**

  ```bash
  aictx amberpixels/aictx

  # supported any type of github repo mention.
  # as well as gitlab's repos.
  aictx github.com/amberpixels/aictx
  ```

- **Include specific globs (for both Tree & Source mode) **

  ```bash
  aictx --i "*.go,go.md"
  ```
- **Process a Git Repository with a Custom Size Threshold**

  ```bash
  aictx github.com/amberpixels/aictx --source.threshold 0.2 --source.include="*.js" --out=stdout
  ```
- **List Core Ignore Patterns**

  ```bash
  aictx --list-core-ignores
  ```

- **Include everything: hidden files (and core-ignored)**

  ```bash
  aictx --no-core-ignores --show-hidden
  ```

## Project Structure

~~~text
Project Tree [9 files, 16.05 MB total, max 16.01 MB] (* - for binary files)
.
├── build
│   └── aictx *
├── cmd
│   └── aictx
│       └── main.go
├── internal
│   ├── aictx
│   │   ├── app.go
│   │   ├── core_ignores.go
│   │   ├── git.go
│   │   ├── git_test.go
│   │   └── helpers.go
│   └── fsutils
│       └── fsutils.go
└── tools
    └── readmegen.go
~~~

## Contributing
Contributions are welcome! Please fork the repository and submit pull requests for any enhancements, bug fixes, or suggestions.

**Please Note:** The application is stable and ready-to-be used, but it is not currently covered by tests.
As a result, accepting new features may take longer than expected until we have at least some basic test coverage
in place. We appreciate your patience and contributions as we work towards enhancing
the project's quality and reliability.

## License
This project is licensed under the [MIT License](LICENSE).

Happy coding! :heart:

==========================[13/13]===========================
File: tools/readmegen.go
Size: 1.24 KB
------------------------------------------------------------
package main

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"

	"github.com/charmbracelet/log"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: "readmegen 📚 ",
	})

	logger.Info("Generating readme.md from template...")
	defer func() {
		logger.Info("readme.md generated successfully.")
	}()

	// Read the template file.
	tmplData, err := os.ReadFile("tools/readme.template.md")
	if err != nil {
		logger.Fatalf("Error reading readme.template.md: %v", err)
	}

	// Run "aictx --help" and capture its output.
	cmd := exec.Command("./build/aictx", "--help")
	helpOutput, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("Failed running './aictx --help': %v", err)
	}

	// Prepare data for the template.
	data := struct {
		Usage string
	}{
		Usage: string(helpOutput),
	}

	// Parse the template.
	tmpl, err := template.New("readme").Parse(string(tmplData))
	if err != nil {
		logger.Fatalf("Failed parsing template: %v", err)
	}

	// Execute the template.
	var output bytes.Buffer
	if err := tmpl.Execute(&output, data); err != nil {
		logger.Fatalf("Failed executing template: %v", err)
	}

	if err := os.WriteFile("readme.md", output.Bytes(), 0600); err != nil {
		logger.Fatalf("Failed writing README.md: %v", err)
	}
}
--- END OF CODEBASE ---
-->
