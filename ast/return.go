package ast

// `return` [ expression ]
type Return struct {
	Token Token
	Value Expression
}

func (_ *Return) statmentNode()        {}
func (r *Return) TokenLiteral() string { return r.Token.Literal }
