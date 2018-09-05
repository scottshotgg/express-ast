package ast

// If we want to allow objects to contain statements in Express objects,
// then we would need to implement Blocks as an ExpressionStatement

// `{` [ statement ]* `}`
type Block struct {
	Token      Token
	Statements []Statement
	Length     int
}

func (_ *Block) statementNode()       {}
func (b *Block) TokenLiteral() string { return b.Token.Literal }
