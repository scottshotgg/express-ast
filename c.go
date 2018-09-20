package ast

type CBlock struct {
	Token Token
	Body  Block
}

// Implement Node and Statement

func (c *CBlock) statementNode() {}

// TokenLiteral returns the literal value of the token
func (c *CBlock) TokenLiteral() string { return c.Token.Literal }
