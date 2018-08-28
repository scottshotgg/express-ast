package ast

type UnaryOp interface {
	Node
	Operand() *Expression
}
