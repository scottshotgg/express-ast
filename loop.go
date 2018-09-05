package ast

type LoopType int

const (
	StdFor LoopType = iota + 1
	ForIn
	ForOf
	ForOver
	While
	Until
)

type Loop struct {
	Token Token
	Type  LoopType
	Start int
	End   int
	Step  int
	// Iter  *Iterable
}
