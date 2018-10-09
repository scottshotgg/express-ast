package ast

import (
	"fmt"

	"github.com/scottshotgg/express-token"
)

// Group is an abstract type that is used in the grammar of the form:
// `(` { element }* `)`
type Group struct {
	Token    token.Token
	TypeOf   Type
	Elements []Expression
}

func (g *Group) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (g *Group) TokenLiteral() token.Token { return g.Token }

// Type implements Literal
func (g *Group) Type() Type { return g.TypeOf }

func (g *Group) Kind() NodeType { return GroupNode }

func (g *Group) String() string {
	// FIXME: just doing this to get it to compile
	return fmt.Sprintf("%+v", *g)
}
