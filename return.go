package ast

import "github.com/scottshotgg/express-token"

// Return represents the following form:
// `return` [ expression ]
type Return struct {
	Token token.Token
	Value Expression
}

func (r *Return) statementNode() {}

// TokenLiteral returns the literal value of the token
func (r *Return) TokenLiteral() token.Token { return r.Token }

func (r *Return) Kind() NodeType { return ReturnNode }

func NewReturn(t token.Token, e Expression) *Return {
	return &Return{
		Token: t,
		Value: e,
	}
}
