package main

import (
	"github.com/planetscale/go-logkeycheck/internal/logkeycheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
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
	singlechecker.Main(logkeycheck.Analyzer)
}
