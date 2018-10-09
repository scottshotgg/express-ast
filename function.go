package ast

import (
	"fmt"

	"github.com/scottshotgg/express-token"
)

// Function represents the following form:
// [ `func` | `fn` ] [ ident ] [ group ] { group } [ block ]
type Function struct {
	Lambda    bool
	Async     bool
	Token     token.Token
	Ident     *Ident
	Arguments *Group
	Returns   *Group
	Body      *Block
}

// Implement statement
func (f *Function) statementNode() {}

// Implement expression
func (f *Function) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (f *Function) TokenLiteral() token.Token { return f.Token }

// Type implements literal so that functions can be assigned to idents
func (f *Function) Type() LiteralType { return FunctionType }

func (f *Function) Kind() NodeType { return FunctionNode }

func (f *Function) String() string {
	// FIXME: just doing this to get it to compile
	return fmt.Sprintf("%+v", *f)
}

func NewFunction(ft, it token.Token, args *Group, body *Block) (*Function, error) {
	ident, err := NewIdent(it, "")
	if err != nil {
		return nil, err
	}

	return &Function{
		Token:     ft,
		Ident:     ident,
		Arguments: args,
		Body:      body,
	}, nil
}
