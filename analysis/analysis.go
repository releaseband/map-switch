package analysis

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/releaseband/map-switch/analysis/internal/helpers"

	"github.com/releaseband/map-switch/analysis/internal"
)

func parseImports(d *ast.GenDecl) ([]string, error) {
	imports := make([]string, len(d.Specs))
	for i, imp := range d.Specs {
		importStr, err := internal.ParseImport(imp)
		if err != nil {
			return nil, err
		}

		imports[i] = importStr
	}

	return imports, nil
}

func parseResult(vcs *ast.ValueSpec, countType string, imports []string) (*internal.Result, error) {
	result := internal.NewResult(countType)

	for _, sv := range vcs.Values {
		lit, err := internal.CastCompositeLit(sv)
		if err != nil && errors.Is(err, internal.ErrCastFailed) {
			continue
		}

		mapType, err := internal.CastMapType(lit.Type)
		if err != nil {
			return nil, err
		}

		mapKeyVal, err := internal.ParseKeyValueTypeFromMapType(mapType)
		if err != nil {
			return nil, err
		}

		v, err := internal.ParseMapValues(lit)
		if err != nil {
			return nil, err
		}

		md := internal.NewMapData(*mapKeyVal, v)
		result.SetMapData(md)

		imps := helpers.GetNeedImports(*mapKeyVal, imports)
		result.SetImports(imps)
	}

	return result, nil
}

func analysisSingleVar(decl *ast.GenDecl, imports []string) ([]internal.Result, error) {
	results := make([]internal.Result, 0, 2)

	comment, err := internal.ParseCommentFromSingleValParam(decl)
	if err != nil {
		return nil, err
	}

	if comment == nil {
		return nil, nil
	}

	result := internal.NewResult(comment.CountType)
	result.FuncName = comment.StructName

	for _, spec := range decl.Specs {
		vc, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		result, err := parseResult(vc, comment.CountType, imports)
		if err != nil {
			return nil, err
		}

		result.FuncName = comment.StructName
		results = append(results, *result)
	}

	return results, nil
}

func analysisMultiVar(decl *ast.GenDecl, imports []string) ([]internal.Result, error) {
	results := make([]internal.Result, 0, 2)

	comments, vcs, err := internal.ParseCommentsFromMultiValParam(decl)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, nil
	}

	for i, comment := range comments {
		result, err := parseResult(vcs[i], comment.CountType, imports)
		if err != nil {
			return nil, err
		}

		result.FuncName = comment.StructName
		results = append(results, *result)
	}

	return results, nil
}

func analysisFileByMap(mapData MapParams) (*internal.FileResult, error) {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, mapData.FilePath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("ParseFile failed: %w", err)
	}

	var (
		imports    []string
		results    []internal.Result
		mResults   []internal.Result
		allResults []internal.Result
	)

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.IMPORT:
				imports, err = parseImports(decl)
				if err != nil {
					return nil, err
				}

			case token.VAR:
				mResults, err = analysisMultiVar(decl, imports)
				if err != nil {
					return nil, err
				}

				results, err = analysisSingleVar(decl, imports)
				if err != nil {
					return nil, err
				}

				if len(results) > 0 {
					allResults = append(allResults, results...)
				}

				if len(mResults) > 0 {
					allResults = append(allResults, mResults...)
				}
			}
		}
	}

	return internal.NewFileResult(f.Name.Name, allResults), nil
}
