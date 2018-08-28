package ast

type Statement interface {
	Node

	// This is just something to force the interface
	statementNode()
}
