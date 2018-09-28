package ast

// Access is the action ([ expression ] `.` [ expression ]) represented by the
// dot operator that allows the internals of a struct to be utilized
type Access struct {
	Token   Token
	Name    string
	Parents []string
}

func (a *Access) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (a *Access) TokenLiteral() string { return a.Token.Literal }

func (a *Access) Kind() NodeType { return AccessNode }
