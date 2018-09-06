package ast

// Ident represents the following form:
// [ type ] [ name ]
type Ident struct {
	Token Token
	Type  string
	Value string
}

func (i *Ident) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (i *Ident) TokenLiteral() string { return i.Token.Literal }
