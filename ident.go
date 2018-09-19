package ast

import (
	"errors"

	ast "github.com/scottshotgg/express-ast"
)

// Ident represents the following form:
// [ name ]
type Ident struct {
	Token Token
	Type  Type
	Name  string
}

func (i *Ident) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (i *Ident) TokenLiteral() string { return i.Token.Literal }

// Might need to make specific type-functions
// But I don't think identifiers here need to have a type, that's NOT what the AST is for; keep track of that in the parser, etc

func NewIdent(t Token, it Type, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  it,
		Name:  n,
	}, nil
}

func NewIntIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  ast.NewIntType(),
		Name:  n,
	}, nil
}

func NewBoolIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  ast.NewBoolType(),
		Name:  n,
	}, nil
}

func NewFloatIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  ast.NewFloatType(),
		Name:  n,
	}, nil
}

func NewCharIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  ast.NewCharType(),
		Name:  n,
	}, nil
}

func NewStringIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  ast.NewStringType(),
		Name:  n,
	}, nil
}

// func NewStructIdent(t Token, n string) (*Ident, error) {
// 	if n == "" {
// 		return nil, errors.New("Cannot use empty string as identifier name")
// 	}

// 	return &Ident{
// 		Token: t,
// 		Type:  ast.NewStructType(),
// 		Name:  n,
// 	}, nil
// }

func NewObjectIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  ast.NewObjectType(),
		Name:  n,
	}, nil
}

func NewFunctionIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  ast.NewFunctionType(),
		Name:  n,
	}, nil
}

// func NewVarIdent(t Token, n string) (*Ident, error) {
// 	if n == "" {
// 		return nil, errors.New("Cannot use empty string as identifier name")
// 	}

// 	return &Ident{
// 		Token: t,
// 		Type:  ast.NewVarType(),
// 		Name:  n,
// 	}, nil
// }
