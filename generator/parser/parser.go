package parser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/releaseband/map-switch/generator/parser/models"
)

const mapGenCommentTag = "//map_gen:"

func parseImports(d *ast.GenDecl) ([]string, error) {
	imports := make([]string, len(d.Specs))
	for i, spec := range d.Specs {
		imp, ok := spec.(*ast.ImportSpec)
		if !ok {
			return nil, errors.New("cast import failed")
		}

		imports[i] = imp.Path.Value
	}

	return imports, nil
}

func getName(commentText string) (string, bool) {
	if strings.HasPrefix(commentText, mapGenCommentTag) {
		elements := strings.Split(strings.Replace(commentText, mapGenCommentTag, "", 1), ";")

		for _, e := range elements {
			if strings.HasPrefix(e, "name=") {
				return strings.Replace(e, "name=", "", 1), true
			}
		}
	}

	return "", false
}

func isSingleVariable(decl *ast.GenDecl) bool {
	return decl.Doc != nil && len(decl.Doc.List) > 0
}

func parseSingleVariantData(decl *ast.GenDecl) (*models.Variant, error) {
	var (
		name string
		ok   bool
	)

	if decl.Doc != nil {
		for _, c := range decl.Doc.List {
			name, ok = getName(c.Text)
			if ok {
				break
			}
		}
	}

	if !ok {
		return nil, nil
	}

	mapData, err := searchAndParseMap(decl.Specs)
	if err != nil {
		return nil, fmt.Errorf("searchAndParseMap: %w", err)
	}

	return models.NewVariant(name, mapData), nil
}

func parseVar(decl *ast.GenDecl) ([]models.Variant, error) {
	var variants []models.Variant

	if isSingleVariable(decl) {
		variant, err := parseSingleVariantData(decl)
		if err != nil {
			return nil, fmt.Errorf("parseSingleVariantData: %w", err)
		}

		if variant != nil {
			variants = append(variants, *variant)
		}
	}

	return variants, nil
}

func ParseFile(path string) (*models.FileDeclaration, error) {
	set, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %w", err)
	}

	fileDecl := models.NewFileDeclaration(set.Name.Name)

	for _, d := range set.Decls {
		decl, ok := d.(*ast.GenDecl)
		if ok {
			switch decl.Tok {
			case token.VAR:
				variants, err := parseVar(decl)
				if err != nil {
					return nil, err
				}

				if len(variants) > 0 {
					fileDecl.AddVariants(variants)
				}
			}
		}
	}

	return fileDecl, nil
}
