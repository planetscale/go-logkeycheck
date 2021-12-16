package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/planetscale/go-logkeycheck/internal/logkeycheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var (
	version = "development"
	date    = "development"
)

type analyzerPlugin struct{}

// AnalyzerPlugin exposes the required interface for a golangci-lint plugin.
// https://golangci-lint.run/contributing/new-linters/#create-a-plugin
var AnalyzerPlugin analyzerPlugin //nolint deadcode:unused

// GetAnalyzers implements golangci-lint's plugin interface.
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{logkeycheck.Analyzer}
}

func main() {
	// h/t to @fatih for this:
	// this is a small hack to implement the -V flag that is part of
	// go/analysis framework. It'll allow us to print the version with -V, but
	// the --help message will print the flags of the analyzer
	ff := flag.NewFlagSet("logkeycheck", flag.ContinueOnError)
	v := ff.Bool("V", false, "print version and exit")
	ff.Usage = func() {}
	ff.SetOutput(io.Discard)

	ff.Parse(os.Args[1:]) // nolint: errcheck
	if *v {
		fmt.Printf("%s (%s)\n", version, date)
		os.Exit(0)
	}

	singlechecker.Main(logkeycheck.Analyzer)
}
