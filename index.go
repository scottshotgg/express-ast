package ast

// Index is the action represented by the square brackets ([ expression ] `[` [ expression ] `]`)
// that allows the internals of an object, array, or map to be utilized
type Index struct {
	Token    Token
	Name     string
	Indicies []Expression
}

func (i *Index) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (i *Index) TokenLiteral() string { return i.Token.Literal }

func (i *Index) Kind() NodeType { return IndexNode }
