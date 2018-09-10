package ast

// If we want to allow objects to contain statements in Express objects,
// then we would need to implement Blocks as an ExpressionStatement

// TODO: try simplifying this and make one top level block per file and have a program effectively be a block

// ExpressionBlock statement represents the following form:
// `{` [ statement ]* `}`
type ExpressionBlock struct {
	Token       Token
	Expressions []Expression
}

// TODO: implement expression
func (eb *ExpressionBlock) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (eb *ExpressionBlock) TokenLiteral() string { return eb.Token.Literal }

// Length returns the length of statments in the block
func (eb *ExpressionBlock) Length() int {
	return len(eb.Expressions)
}
