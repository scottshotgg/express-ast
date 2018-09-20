package ast

type Switch struct {
	Token      Token
	Expression Expression
	Cases      []*Case
	Default    Statement
}

type Case struct {
	Token      Token
	Expression Expression
	Body       Statement
}

// Implement Node and Statement

func (s *Switch) statementNode() {}

// TokenLiteral returns the literal value of the token
func (s *Switch) TokenLiteral() string { return s.Token.Literal }
