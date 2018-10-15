package ast

import (
	"fmt"

	"github.com/scottshotgg/express-token"
)

// type ArrayType int

// const (
// 	_ ArrayType = iota

// 	Homogenous

// 	Heterogeneous
// )

// Array represents array type data structures
type Array struct {
	Token token.Token
	// How will this act with `var` elements?
	TypeOf     *Type
	Length     int
	Elements   []Expression
	Homogenous bool
}

// TODO: this should implement iterable.... no?

func (a *Array) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (a *Array) TokenLiteral() token.Token { return a.Token }

// Type implements Literal
func (a *Array) Type() *Type { return a.TypeOf }

func (a *Array) Kind() NodeType { return ArrayNode }

func (a *Array) String() string {
	// FIXME: just doing this to get it to compile
	return fmt.Sprintf("%+v", *a)
}

func NewArray(t token.Token, elements []Expression) *Array {
	homogenous := true

	var typeOf *Type
	if len(elements) > 0 {
		typeOf = elements[0].Type()
		for _, e := range elements[1:] {
			// TODO: should actually do a comparison for upgradable types here...
			if e.Type().Type != typeOf.Type {
				homogenous = false
				break
			}
		}
	}

	return &Array{
		Token:      t,
		TypeOf:     typeOf,
		Length:     len(elements),
		Elements:   elements,
		Homogenous: homogenous,
	}
}
