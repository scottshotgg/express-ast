package ast

// If we want to allow objects to contain statements in Express objects,
// then we would need to implement Blocks as an ExpressionStatement

// Block statement represents the following form:
// `{` [ statement ]* `}`
type Block struct {
	Token      Token
	Statements []Statement
	Length     int
}

func (b *Block) statementNode() {}

// TokenLiteral returns the literal value of the token
func (b *Block) TokenLiteral() string { return b.Token.Literal }
