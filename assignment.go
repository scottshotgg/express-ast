package ast

// AssignmentType encompasses the different types of assignment
type AssignmentType int

const (
	// Equals is the = operator
	Equals AssignmentType = iota + 1

	// Set is the : operator
	Set

	// Init is the := operator
	Init
)

// Assignment statement represents the following form:
// { type } [ ident ] [ assign_op ] [ expression ]
type Assignment struct {
	Declaration bool
	Inferred    bool
	Token       Token
	Type        AssignmentType
	Ident       Ident
	Value       *Expression
}

func (a *Assignment) statmentNode() {}

// TokenLiteral returns the literal value of the token
func (a *Assignment) TokenLiteral() string { return a.Token.Literal }
