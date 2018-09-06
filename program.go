package ast

// Program represents the following form:
// [ statement ]*
type Program struct {
	// We'll wanna put 'File' in between Program and Block at some point
	Statements []Statement
	Length     int
}
