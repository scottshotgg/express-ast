package ast

// // DeclarationType encompasses the different types of assignment
// type DeclarationType int

// const (
// 	// Equals is the = operator
// 	Equals DeclarationType = iota + 1
// )

// // Declaration statement represents the following form:
// // [ type ] [ ident ] { [ assign_op ] [ expression ] }
// type Declaration struct {
// 	Token Token
// 	Type  DeclarationType
// 	Ident Ident
// 	Value Expression
// }

// func (d *Declaration) statmentNode() {}

// // TokenLiteral returns the literal value of the token
// func (d *Declaration) TokenLiteral() string { return d.Token.Literal }
