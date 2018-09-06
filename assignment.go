package ast

type AssignmentType int

const (
	// Equals =
	Equals AssignmentType = iota + 1

	// Set :
	Set

	// Init :=
	Init
)

// [ ident ] [ assign_op ] [ expression ]
type Assignment struct {
	Token Token
	Type  AssignmentType
	Ident Ident
	Value Expression
}

func (_ *Assignment) statmentNode()        {}
func (a *Assignment) TokenLiteral() string { return a.Token.Literal }
