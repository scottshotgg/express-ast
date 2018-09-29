package ast

import "github.com/scottshotgg/express-token"

// Array represents array type data structures
type Array struct {
	Token token.Token
	// How will this act with `var` elements?
	TypeOf LiteralType
	Length int
}

func (a *Array) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (a *Array) TokenLiteral() token.Token { return a.Token }

// Type implements Literal
func (a *Array) Type() LiteralType { return a.TypeOf }

func (a *Array) Kind() NodeType { return ArrayNode }
