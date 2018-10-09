package ast

import (
	"fmt"

	"github.com/scottshotgg/express-token"
)

// this file might be better served from the actual parser or whatever is working with the AST

// DefaultLiteral is a type of Literal token
type DefaultLiteral struct {
	Token  token.Token
	TypeOf Type
	Value  string
}

func (d *DefaultLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (d *DefaultLiteral) TokenLiteral() token.Token { return d.Token }

// Type implements literal
func (d *DefaultLiteral) Type() Type { return d.TypeOf }

func (d *DefaultLiteral) Kind() NodeType { return LiteralNode }

func (d *DefaultLiteral) String() string {
	// FIXME: just doing this to get it to compile
	return fmt.Sprintf("%+v", *d)
}

// NewDefault returns a new int literal
func NewDefault(t token.Token) *DefaultLiteral {
	return &DefaultLiteral{
		Token: t,
		Value: t.Value.String,
	}
}
