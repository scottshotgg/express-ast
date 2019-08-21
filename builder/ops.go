package builder

import (
	types "github.com/scottshotgg/express-ast/types"
	token "github.com/scottshotgg/express-token"
)

func (b *Builder) ParseSet(n *types.Node) (*types.Node, error) {
	// This will be encountered when we have:
	// <expr> `:` <expr>

	// Step over the set token
	b.index++

	right, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	// Step over the Expression
	b.index++

	return &types.Node{
		Type:  "kv",
		Left:  n,
		Right: right,
	}, nil
}

func (b *Builder) ParseBinOp(n *types.Node) (*types.Node, error) {
	var op = b.tokens[b.index].Value.String

	// Step over the operator token
	b.index++

	right, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	return &types.Node{
		Type:  "binop",
		Value: op,
		Left:  n,
		Right: right,
	}, nil
}

func (b *Builder) ParseLessThanExpression(n *types.Node) (*types.Node, error) {
	// Step over the conditional operator token
	b.index++

	right, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	return &types.Node{
		Type:  "comp",
		Value: "<",
		Left:  n,
		Right: right,
	}, nil
}

func (b *Builder) ParseGreaterThanExpression(n *types.Node) (*types.Node, error) {
	// Step over the conditional operator token
	b.index++

	right, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	return &types.Node{
		Type:  "comp",
		Value: ">",
		Left:  n,
		Right: right,
	}, nil
}

func (b *Builder) ParseLessOrEqualThanExpression(n *types.Node) (*types.Node, error) {
	// Step over the conditional operator token
	b.index++

	right, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	return &types.Node{
		Type:  "comp",
		Value: "<=",
		Left:  n,
		Right: right,
	}, nil
}

func (b *Builder) ParseGreaterOrEqualThanExpression(n *types.Node) (*types.Node, error) {
	// Step over the conditional operator token
	b.index++

	right, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	return &types.Node{
		Type:  "comp",
		Value: ">=",
		Left:  n,
		Right: right,
	}, nil
}

func (b *Builder) ParseEqualityExpression(n *types.Node) (*types.Node, error) {
	// Step over the conditional operator token
	b.index++

	right, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	return &types.Node{
		Type:  "comp",
		Value: "==",
		Left:  n,
		Right: right,
	}, nil
}

func (b *Builder) ParseIncrement(n *types.Node) (*types.Node, error) {
	return &types.Node{
		Type: "inc",
		Left: n,
	}, nil
}

func (b *Builder) ParseCall(n *types.Node) (*types.Node, error) {
	// Check ourselves ...
	if b.tokens[b.index].Type != token.LParen {
		return nil, b.AppendTokenToError("Could not get left paren")
	}

	// We are not allowing for named arguments right now
	args, err := b.ParseGroupOfExpressions()
	if err != nil {
		return nil, err
	}

	return &types.Node{
		Type:  "call",
		Value: n,
		Metadata: map[string]interface{}{
			"args": args,
		},
	}, nil
}

func (b *Builder) ParseIndexExpression(n *types.Node) (*types.Node, error) {
	if b.index > len(b.tokens)-1 {
		return nil, ErrOutOfTokens
	}

	if b.tokens[b.index].Type != token.LBracket {
		return nil, b.AppendTokenToError("Could not get left bracket")
	}

	b.index++

	expr, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	// Step over the expression
	b.index++

	return &types.Node{
		Type: "index",
		// Value: n,
		Left:  n,
		Right: expr,
	}, nil
}

func (b *Builder) ParseSelection(n *types.Node) (*types.Node, error) {
	if b.index > len(b.tokens)-1 {
		return nil, ErrOutOfTokens
	}

	if b.tokens[b.index].Type != token.Accessor {
		return nil, b.AppendTokenToError("Could not get selection operator")
	}

	// Step over the accessor
	b.index++

	expr, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	// b.index++

	return &types.Node{
		Type: "selection",
		// Value: n,
		Left:  n,
		Right: expr,
	}, nil
}
