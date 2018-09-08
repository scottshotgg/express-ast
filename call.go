package ast

// FIXME: need to think about this more

// Call represents the following form:
// [ ident ] [ group ]
type Call struct {
	Token     Token
	Ident     *Ident
	Arguments []*Expression
	Returns   []*Expression
}

// TODO: implement statement
func (c *Call) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (c *Call) TokenLiteral() string { return c.Token.Literal }
