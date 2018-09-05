package ast

type MultiStatement interface {
	Node

	// This is just something to force the interface
	statementNode()

	// statement crud here
	// This should be used for block, etc
}
