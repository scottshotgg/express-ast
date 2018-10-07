package ast

import (
	"errors"

	"github.com/scottshotgg/express-token"
)

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
	Token       token.Token
	Type        AssignmentType
	// For now put Expression here but I think this should be a type `Assignable` where an Expression implements an `Assignable` property
	LHS Expression
	RHS Expression
}

// When going through the logic for this:
//	- if Declaration is already set to true when Inferred is being set to true -> error
//	- if the variable is already declared and Declared or Inferred is being set to true -> error

func (a *Assignment) statementNode() {}

// TokenLiteral returns the literal value of the token
func (a *Assignment) TokenLiteral() token.Token { return a.Token }

func (a *Assignment) Kind() NodeType { return AssignmentNode }

// TODO: dont think I wanna do this yet
// func NewAssignmentStatement() Assignment {
// 	return &Assignment{

// 	}
// }

// NewAssignment returns a new assignment statement and determines whether it is inferred
func NewAssignment(t token.Token, i *Ident, e Expression) (*Assignment, error) {
	if e == nil {
		return nil, errors.New("Expression value cannot by nil")
	}

	as := Assignment{
		Token: t,
		LHS:   i,
		RHS:   e,
	}

	switch t.Value.String {
	case ":=":
		as.Type = Init
		as.Declaration = true
		as.Inferred = true

	case ":":
		as.Type = Set
		as.Declaration = true

	case "=":
		as.Type = Equals

	default:
		return nil, errors.New("Could not detect assingment type from token")
	}

	return &as, nil
}

// SetDeclaration changes the assignment type to a declaration
func (a *Assignment) SetDeclaration(declaration bool) {
	a.Declaration = declaration
}

// SetInferred changes the assignment type to infer the lvalue type
func (a *Assignment) SetInferred(inferred bool) {
	a.Inferred = inferred
}

// func (a *Assignment) ExpressionType() {
// 	if a.Value != nil {
// 		return
// 	}
// }
