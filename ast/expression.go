package ast

type Expression interface {
	Node

	// This is just something to force the interface
	expressionNode()
}
