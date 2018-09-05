package ast

// `function` [ ident ] [ group ] { group } [ block ]
type Function struct {
	Token     Token
	Name      string
	Arguments []Expression
	Returns   []Expression
	Body      Block
}

func (_ *Function) statementNode()       {}
func (f *Function) TokenLiteral() string { return f.Token.Literal }
