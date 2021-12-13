package logkeycheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

// docs for analysistest helpers: https://pkg.go.dev/golang.org/x/tools/go/analysis/analysistest#Run

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
