package ast

// FIXME: need to think about this more

// [ ident ] [ group ]
type Call struct {
	Token     Token
	Ident     *Ident
	Arguments []*Expression
	Returns   []*Expression
}

func (_ *Call) expressionNode()      {}
func (c *Call) TokenLiteral() string { return c.Token.Literal }
