package ast

type BinaryOpType int

const (
	AdditionBinaryOp BinaryOpType = iota + 1
	SubtractionBinaryOp
	MultiplicationBinaryOp
	DivisionBinaryOp
)

// type BinaryOp interface {
// 	Expression

// 	Type() BinaryOpType
// 	Right() *Expression
// 	Left() *Expression
// 	Evaluate() *Literal
// }

// [ expression ] [ bin_op ] [ expression ]
type BinaryOperation struct {
	Token     Token
	Kind      BinaryOpType
	LeftNode  Expression
	RightNode Expression
	Value     Literal
}

// func (b *BinaryOperation) Type() *Expression     { return b.Kind }
// func (b *BinaryOperation) Right() *Expression    { return b.RightExpr }
// func (b *BinaryOperation) Left() *Expression     { return b.LeftExpr }
// func (b *BinaryOperation) Evaluate() *Expression { return b.Value }

func (_ *BinaryOperation) expressionNode()      {}
func (b *BinaryOperation) TokenLiteral() string { return b.Token.Literal }
