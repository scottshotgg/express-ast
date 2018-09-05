package ast

// `return` [ expression ]
type Return struct {
	Token Token
	Value []Expression
}

func (_ *Return) statementNode()       {}
func (r *Return) TokenLiteral() string { return r.Token.Literal }
