package ast

// C blocks fundamentally contain statements that cannot be checked
// at compile time by the Express compiler. These statements will be
// directly injected into the generated C++ source code that will
// then be checked by Clang.

// CBlock statement represents the following form:
// `c {` [ c_statement ]* `}`
type CBlock struct {
	Token Token
	Body  Block
}

// Implement Node and Statement

func (c *CBlock) statementNode() {}

// TokenLiteral returns the literal value of the token
func (c *CBlock) TokenLiteral() string { return c.Token.Literal }

func (c *CBlock) Kind() NodeType { return CBlockNode }
