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
