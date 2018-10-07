package ast

import "github.com/scottshotgg/express-token"

// FIXME: need to think about this more

// Call represents the following form:
// [ ident ] [ group ]
type Call struct {
	Token     token.Token
	Ident     *Ident
	Arguments *Group
	Returns   []Expression
}

func (c *Call) expressionNode() {}
func (c *Call) statementNode()  {}

// TokenLiteral returns the literal value of the token
func (c *Call) TokenLiteral() token.Token { return c.Token }

func (c *Call) Kind() NodeType { return CallNode }
