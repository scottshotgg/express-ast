package ast

// FIXME: need to think about this one some more

// `fn` [ group ] { group } [ block ] [ group ]
type Fn struct {
	Token Token
}

func (_ *Fn) expressionNode()      {}
func (f *Fn) TokenLiteral() string { return f.Token.Literal }
