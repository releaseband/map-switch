package internal

import (
	"errors"
	"go/ast"
)

var ErrCastFailed = errors.New("cast failed")

func CastCompositeLit(expr ast.Expr) (*ast.CompositeLit, error) {
	l, ok := expr.(*ast.CompositeLit)
	if !ok {
		return nil, ErrCastFailed
	}

	return l, nil
}

func CastBasicLit(expr ast.Expr) (*ast.BasicLit, error) {
	bl, ok := expr.(*ast.BasicLit)
	if !ok {
		return nil, ErrCastFailed
	}

	return bl, nil
}

func CastMapType(expr ast.Expr) (*ast.MapType, error) {
	m, ok := expr.(*ast.MapType)
	if !ok {
		return nil, ErrCastFailed
	}

	return m, nil
}

func CastTypeSpec(i interface{}) (*ast.TypeSpec, error) {
	s, ok := i.(*ast.TypeSpec)
	if !ok {
		return nil, ErrCastFailed
	}

	return s, nil
}

func CastAstIdent(expr ast.Expr) (*ast.Ident, error) {
	i, ok := expr.(*ast.Ident)
	if !ok {
		return nil, ErrCastFailed
	}

	return i, nil
}

func CastArrayType(expr ast.Expr) (*ast.ArrayType, error) {
	arr, ok := expr.(*ast.ArrayType)
	if !ok {
		return nil, ErrCastFailed
	}

	return arr, nil
}

func CastKeyValueExpr(expr ast.Expr) (*ast.KeyValueExpr, error) {
	e, ok := expr.(*ast.KeyValueExpr)
	if !ok {
		return nil, ErrCastFailed
	}

	return e, nil
}

func GetSelectorExpr(expr ast.Expr) (*ast.SelectorExpr, error) {
	s, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return nil, ErrCastFailed
	}

	return s, nil
}

func CastImportSpec(spec ast.Spec) (*ast.ImportSpec, error) {
	i, ok := spec.(*ast.ImportSpec)
	if !ok {
		return nil, ErrCastFailed
	}

	return i, nil
}

func CastSliceBasicList(expr []ast.Expr) ([]*ast.BasicLit, error) {
	bls := make([]*ast.BasicLit, len(expr))
	for i, e := range expr {
		bl, err := CastBasicLit(e)
		if err != nil {
			return nil, err
		}

		bls[i] = bl
	}

	return bls, nil
}

func CastValueSpec(s ast.Spec) (*ast.ValueSpec, error) {
	vc, ok := s.(*ast.ValueSpec)
	if !ok {
		return nil, ErrCastFailed
	}

	return vc, nil
}
