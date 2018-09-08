package ast

// Function represents the following form:
// `function` [ ident ] [ group ] { group } [ block ]
type Function struct {
	Lambda    bool
	Async     bool
	Token     Token
	Name      string
	Arguments []Expression
	Returns   []Expression
	Body      Block
}

func (f *Function) statementNode()  {}
func (f *Function) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (f *Function) TokenLiteral() string { return f.Token.Literal }
