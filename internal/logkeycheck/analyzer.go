package logkeycheck

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	snakeCaseRE = regexp.MustCompile("^[a-z]+(_[a-z]+)*$")

	// List of all packages containing funcs to analyze.
	pkgs = map[string]struct{}{
		"github.com/planetscale/log": {},
		"go.uber.org/zap":            {},
	}

	// TODO: convert to struct
	// list of all zap funcs we are interested in analyzing.
	funcs = map[string]struct{}{
		"Any":         {},
		"Array":       {},
		"Binary":      {},
		"Bool":        {},
		"Boolp":       {},
		"Bools":       {},
		"ByteString":  {},
		"ByteStrings": {},
		"Complex128":  {},
		"Complex128p": {},
		"Complex128s": {},
		"Complex64":   {},
		"Complex64p":  {},
		"Complex64s":  {},
		"Duration":    {},
		"Durationp":   {},
		"Durations":   {},
		"Errors":      {},
		"Float32":     {},
		"Float32p":    {},
		"Float32s":    {},
		"Float64":     {},
		"Float64p":    {},
		"Float64s":    {},
		"Inline":      {},
		"Int":         {},
		"Int16":       {},
		"Int16p":      {},
		"Int16s":      {},
		"Int32":       {},
		"Int32p":      {},
		"Int32s":      {},
		"Int64":       {},
		"Int64p":      {},
		"Int64s":      {},
		"Int8":        {},
		"Int8p":       {},
		"Int8s":       {},
		"Intp":        {},
		"Ints":        {},
		"NamedError":  {},
		"Namespace":   {},
		"Object":      {},
		"Reflect":     {},
		"Skip":        {},
		"Stack":       {},
		"StackSkip":   {},
		"String":      {},
		"Stringer":    {},
		"Stringp":     {},
		"Strings":     {},
		"Time":        {},
		"Timep":       {},
		"Times":       {},
		"Uint":        {},
		"Uint16":      {},
		"Uint16p":     {},
		"Uint16s":     {},
		"Uint32":      {},
		"Uint32p":     {},
		"Uint32s":     {},
		"Uint64":      {},
		"Uint64p":     {},
		"Uint64s":     {},
		"Uint8":       {},
		"Uint8p":      {},
		"Uint8s":      {},
		"Uintp":       {},
		"Uintptr":     {},
		"Uintptrp":    {},
		"Uintptrs":    {},
		"Uints":       {},
	}
)

var Analyzer = &analysis.Analyzer{
	Name:     "logkeycheck",
	Doc:      "Checks that log keys are properly formatted",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// pass.ResultOf[inspect.Analyzer] will be set if we've added inspect.Analyzer to Requires.
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr)

		// all of the funcs we're interested in have at least 2 args
		if len(call.Args) < 2 {
			return
		}

		fun, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		pkg := pass.TypesInfo.Uses[fun.Sel].Pkg()

		// only interested in funcs from these packages:
		// Use .Path() to match against the fully qualified package name in case the import has been aliased.
		if _, ok := pkgs[pkg.Path()]; !ok {
			return
		}

		// only interested in these funcs:
		if _, ok := funcs[fun.Sel.Name]; !ok {
			return
		}

		// the first argument must be a string
		firstArg, ok := call.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}
		if firstArg.Kind != token.STRING {
			return
		}

		// remove double quotes around the arg string
		trimmed := firstArg.Value[1 : len(firstArg.Value)-1]
		if trimmed == "" {
			return
		}

		if !snakeCaseRE.MatchString(trimmed) {
			pass.Report(analysis.Diagnostic{
				Pos:            firstArg.Pos(),
				End:            firstArg.End(),
				Category:       "logging",
				Message:        fmt.Sprintf("log key '%s' should be snake_case.", trimmed),
				SuggestedFixes: nil,
			})
		}
	})

	return nil, nil
}
