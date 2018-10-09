package ast

import (
	"fmt"

	"github.com/scottshotgg/express-token"
)

// If we want to allow objects to contain statements in Express objects,
// then we would need to implement Blocks as an ExpressionStatement

// TODO: try simplifying this and make one top level block per file and have a program effectively be a block

// Block statement represents the following form:
// `{` [ statement ]* `}`
type Block struct {
	Token      token.Token
	Statements []Statement

	// TODO: Need to solve where this goes. I think it should go in the parser information,
	// but if that's the case then it'll be a bit hard to link an object property
	Scope map[string]Expression
}

// TODO: implement expression
func (b *Block) statementNode()     {}
func (b *Block) expressionNode()    {}
func (b *Block) elseStatementNode() {}

// TokenLiteral returns the literal value of the token
func (b *Block) TokenLiteral() token.Token { return b.Token }

// Length returns the length of statments in the block
func (b *Block) Length() int { return len(b.Statements) }

func (b *Block) Kind() NodeType { return BlockNode }

func (b *Block) String() string {
	// FIXME: just doing this to get it to compile
	return fmt.Sprintf("%+v", *b)
}
