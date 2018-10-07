package ast

import "github.com/scottshotgg/express-token"

type UnaryType int

const (
	_ UnaryType = iota
	Increment
	Decrement
)

type UnaryOp struct {
	Token token.Token
	Op    UnaryType
	Value Expression
}

func (u *UnaryOp) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (u *UnaryOp) TokenLiteral() token.Token { return u.Token }

func (u *UnaryOp) Kind() NodeType { return ConditionNode }
