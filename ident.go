package ast

// [ type ] [ name ]
type Ident struct {
	Token Token
	Type  string
	Value string
}

func (_ *Ident) expressionNode()      {}
func (i *Ident) TokenLiteral() string { return i.Token.Literal }
