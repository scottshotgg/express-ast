package ast

type AssignmentType int

const (
	// =
	Equals AssignmentType = iota + 1

	// :
	Set

	// :=
	Init
)

// [ ident ] [ assign_op ] [ expression ]
type Assignment struct {
	Token Token
	Type  AssignmentType
	Ident Ident
	Value Expression
}

func (_ *Assignment) statmentNode() {}
func (a *Assignment) TokenLiteral() { return a.Token.Literal }
