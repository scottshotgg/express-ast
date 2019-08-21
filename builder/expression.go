package builder

import (
	"fmt"

	types "github.com/scottshotgg/express-ast/types"
	token "github.com/scottshotgg/express-token"
)

func (b *Builder) ParseExpression() (types.Expression, error) {
	term, err := b.ParseTerm()
	if err != nil {
		return term, err
	}

	var (
		ok     bool
		opFunc opCallbackFn
	)

	// LOOKAHEAD performed to figure out whether the expression is done
	for b.index < len(b.tokens)-1 {
		// Look for a tier2 operator in the func map
		opFunc, ok = b.OpFuncMap[1][b.tokens[b.index+1].Type]
		if !ok {
			break
		}

		// Step over the factor
		b.index++

		term, err = opFunc(term)
		if err != nil {
			return term, err
		}
	}

	return term, nil
}

// TODO: Maybe even define types for Term and Factor
func (b *Builder) ParseTerm() (types.Expression, error) {
	factor, err := b.ParseFactor()
	if err != nil {
		return factor, err
	}

	var (
		ok     bool
		opFunc opCallbackFn
	)

	// LOOKAHEAD performed to figure out whether the expression is done
	for b.index < len(b.tokens)-1 {

		// Look for a tier1 operator in the func map
		opFunc, ok = b.OpFuncMap[0][b.tokens[b.index+1].Type]
		if !ok {
			break
		}
		fmt.Println("OPFUNC", b.tokens[b.index+1])

		// Step over the factor
		b.index++

		factor, err = opFunc(factor)
		if err != nil {
			// if err == ErrOutOfTokens {
			// 	return factor,
			// }

			return factor, err
		}
	}

	return factor, nil
}

func (b *Builder) ParseFactor() (types.Expression, error) {
	// Here we will switch on the type and determine whether we have:
	// - literal
	// - ident
	// - call
	// - index operation
	// - selection operation
	// - block
	// - array
	// - nil

	if b.index > len(b.tokens)-1 {
		return nil, ErrOutOfTokens
		// return nil, nil
	}

	switch b.tokens[b.index].Type {
	// Any literal value
	case token.Literal:
		return &types.Literal{
			Value: b.tokens[b.index].Value.True,
		}, nil

	// Variable identifier
	case token.Ident:
		// // Check the scope map for the variable, if we already have a variable declared then use that
		// var node = b.ScopeTree.Get(b.tokens[b.index].Value.String)
		// if node != nil {
		// 	// TODO: might need to fix this
		// 	return node, nil
		// }

		return &types.Ident{
			Token: b.tokens[b.index],
			Name:  b.tokens[b.index].Value.String,
		}, nil

		// // Deref operator
		// case token.PriOp:
		// 	return b.ParseDerefExpression()

		// // Ref operator
		// case token.Ampersand:
		// 	return b.ParseRefExpression()

		// // Nested expression
		// case token.LParen:
		// 	return b.ParseNestedExpression()

		// // Array expression
		// case token.LBracket:
		// 	return b.ParseArrayExpression()

		// // Named block
		// case token.LBrace:
		// 	var a, c = b.ParseBlockStatement()
		// 	// If this is an expression, then whatever called ParseExpression
		// 	// is going to increment the index again ...
		// 	b.index--
		// 	return a, c
	}

	return nil, b.AppendTokenToError("Could not parse expression from token")
}

// func (b *Builder) ParseNestedExpression() (*types.Node, error) {
// 	// Check ourselves
// 	if b.tokens[b.index].Type != token.LParen {
// 		return nil, b.AppendTokenToError("Could not get nested expression")
// 	}

// 	// Skip over the left paren
// 	b.index++

// 	expr, err := b.ParseExpression()
// 	if err != nil {
// 		return expr, err
// 	}

// 	// Skip over the expression
// 	b.index++

// 	if b.tokens[b.index].Type != token.RParen {
// 		return nil, b.AppendTokenToError("No right paren found at end of nested expression")
// 	}

// 	// Skip over the right paren
// 	b.index++

// 	return expr, nil
// }

// func (b *Builder) ParseArrayExpression() (*types.Node, error) {
// 	// Check ourselves
// 	if b.tokens[b.index].Type != token.LBracket {
// 		return nil, b.AppendTokenToError("Could not get array expression")
// 	}

// 	// Skip over the left bracket token
// 	b.index++

// 	var (
// 		expr  *types.Node
// 		exprs []*types.Node
// 		err   error
// 	)

// 	for b.index < len(b.tokens) && b.tokens[b.index].Type != token.RBracket {
// 		expr, err = b.ParseExpression()
// 		if err != nil {
// 			return expr, err
// 		}

// 		b.index++

// 		exprs = append(exprs, expr)

// 		// Check and skip over the separator
// 		if b.tokens[b.index].Type == token.Separator {
// 			b.index++
// 		}
// 	}

// 	// // Step over the right bracket token
// 	// b.index++

// 	return &types.Node{
// 		Type:  "array",
// 		Value: exprs,
// 	}, nil
// }

// func (b *Builder) ParseGroupOfExpressions() (*types.Node, error) {
// 	// Check ourselves
// 	if b.tokens[b.index].Type != token.LParen {
// 		return nil, b.AppendTokenToError("Could not get group of expressions")
// 	}

// 	// Skip over the left paren token
// 	b.index++

// 	var (
// 		expr  *types.Node
// 		exprs []*types.Node
// 		err   error
// 	)

// 	for b.tokens[b.index].Type != token.RParen {
// 		expr, err = b.ParseExpression()
// 		if err != nil {
// 			return expr, err
// 		}

// 		// Step over the expression token
// 		b.index++

// 		exprs = append(exprs, expr)

// 		// Check and skip over the separator
// 		if b.tokens[b.index].Type == token.Separator {
// 			b.index++
// 		}
// 	}

// 	// // Step over the right paren token
// 	// b.index++

// 	return &types.Node{
// 		Type:  "egroup",
// 		Value: exprs,
// 	}, nil
// }

// func (b *Builder) ParseDerefExpression() (*types.Expression, error) {
// 	// Check ourselves ...
// 	if b.tokens[b.index].Type != token.PriOp &&
// 		b.tokens[b.index].Value.String == "*" {
// 		return nil, b.AppendTokenToError("Could not get deref")
// 	}

// 	// Look ahead and make sure it is an ident;you can't deref just anything...
// 	if b.tokens[b.index+1].Type != token.Ident {
// 		return nil, b.AppendTokenToError("Could not get ident to deref")
// 	}

// 	// Step over the deref
// 	b.index++

// 	ident, err := b.ParseExpression()
// 	if err != nil {
// 		return ident, err
// 	}

// 	return &types.Node{
// 		Type: "deref",
// 		Left: ident,
// 	}, nil
// }

// func (b *Builder) ParseRefExpression() (*types.Node, error) {
// 	// Check ourselves ...
// 	if b.tokens[b.index].Type != token.Ampersand &&
// 		b.tokens[b.index].Value.String == "&" {
// 		return nil, b.AppendTokenToError("Could not get ref")
// 	}

// 	// Look ahead and make sure it is an ident; you can't ref anything...
// 	if b.tokens[b.index+1].Type != token.Ident {
// 		return nil, b.AppendTokenToError("Could not get ident to ref")
// 	}

// 	// Step over the deref
// 	b.index++

// 	// Will probably have to change this to just parse the ident instead
// 	// so we don't have problems with operator precedence
// 	ident, err := b.ParseExpression()
// 	if err != nil {
// 		return ident, err
// 	}

// 	return &types.Node{
// 		Type: "ref",
// 		Left: ident,
// 	}, nil
// }
