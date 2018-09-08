package ast

// If we want to allow objects to contain statements in Express objects,
// then we would need to implement Blocks as an ExpressionStatement

// TODO: try simplifying this and make one top level block per file and have a program effectively be a block

// Block statement represents the following form:
// `{` [ statement ]* `}`
type Block struct {
	Token      Token
	Statements []Statement
}

// TODO: implement expression
func (b *Block) statementNode()  {}
func (b *Block) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (b *Block) TokenLiteral() string { return b.Token.Literal }

// Length returns the length of statments in the block
func (b *Block) Length() int {
	return len(b.Statements)
}
