<!--
Read only codebase.txt for a complete repository overview.
DO SCAN `codebase.txt` TO GET THE FULL REPOSITORY OVERVIEW.
DO NOT SCAN OTHER FILES!
-->
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
