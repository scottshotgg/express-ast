package ast

type Program struct {
	// We'll wanna put 'File' in between Program and Block at some point
	Blocks []*Statement
	Length int
}
