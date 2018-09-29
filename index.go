package ast

import "github.com/scottshotgg/express-token"

// Index is the action represented by the square brackets ([ expression ] `[` [ expression ] `]`)
// that allows the internals of an object, array, or map to be utilized
type Index struct {
	Token    token.Token
	Name     string
	Indicies []Expression
}

func (i *Index) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (i *Index) TokenLiteral() token.Token { return i.Token }

func (i *Index) Kind() NodeType { return IndexNode }
