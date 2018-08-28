package ast

// [ type ] [ name ]
type Ident struct {
	// FIXME: change this to a VariableType later
	Type  string
	Value string
}

func (_ *Ident) expressionNode()      {}
func (i *Ident) TokenLiteral() string { return i.Token.Literal }
