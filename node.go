package ast

type Position struct {
	Line   int
	Column int
}

type Location struct {
	Start *Position
	End   *Position
}

type Node interface {
	// TODO: this will just be a string for now until I rework the lexer
	TokenLiteral() string

	// Location() map[string]*Location
}
