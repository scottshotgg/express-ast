package ast

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express-token"
)

type ASTBuilder struct {
	Tokens []token.Token
	Index  int
}

func (a *ASTBuilder) GetFactor() (Expression, error) {
	currentToken := a.Tokens[a.Index]

	switch currentToken.Type {

	case token.Literal:
		switch currentToken.Value.Type {

		case token.IntType:
			return NewInt(currentToken, currentToken.Value.True.(int)), nil

		case token.BoolType:
			return NewBool(currentToken, currentToken.Value.True.(bool)), nil

		case token.CharType:
			return NewChar(currentToken, currentToken.Value.True.(rune)), nil

		case token.StringType:
			return NewString(currentToken, currentToken.Value.True.(string)), nil
		}

	case token.Ident:
		return NewIdent(currentToken, "")
	}

	return nil, errors.Errorf("Could not parse factor from token: %+v", currentToken)
}

func (a *ASTBuilder) GetTerm() (Expression, error) {
	factor, err := a.GetFactor()
	if err != nil {
		return nil, err
	}

	fmt.Println("factor", factor)

	// FIXME: ideally, we should check for a `UNARY` operator class
	nextToken := a.Tokens[a.Index+1]
	if nextToken.Type == token.Increment {
		a.Index++
		factor = &UnaryOp{
			Token: a.Tokens[a.Index],
			Op:    Increment,
			Value: factor,
		}
		a.Index++
	}

	// FIXME: ideally, these should have a `CONDITIONAL` class that we can
	// test and then do a NewConditional with the token
	nextToken = a.Tokens[a.Index+1]
	if nextToken.Type == token.LThan || nextToken.Type == token.GThan {
		a.Index++
		a.Index++
		factor2, err := a.GetExpression()
		if err != nil {
			return nil, err
		}

		var ct ConditionType
		switch nextToken.Type {
		case token.LThan:
			ct = LessThan

		case token.GThan:
			ct = GreaterThan

		default:
			return nil, errors.New("Could not deduce condition type")
		}

		fmt.Println("therese a conditional")
		return &Condition{
			Type:  ct,
			Left:  factor,
			Right: factor2,
		}, nil
	}

	if a.Index+1 < len(a.Tokens)-1 {
		for a.Tokens[a.Index+1].Type == token.PriOp {
			a.Index++

			operand := a.Tokens[a.Index]

			a.Index++
			factor2, err := a.GetExpression()
			if err != nil {
				return nil, err
			}

			fmt.Println("factor2", factor2, operand.Value.String)

			factor, err = NewBinaryOperation(operand, operand.Value.String, factor, factor2)
			if err != nil {
				return nil, err
			}

			if a.Index > len(a.Tokens)-1 {
				break
			}
		}
	}

	fmt.Println("returning")

	return factor, nil
}

func (a *ASTBuilder) GetExpression() (Expression, error) {
	term, err := a.GetTerm()
	if err != nil {
		return nil, err
	}

	// FIXME: ideally, these should have a `CONDITIONAL` class that we can
	// test and then do a NewConditional with the token
	nextToken := a.Tokens[a.Index+1]
	if nextToken.Type == token.LThan || nextToken.Type == token.GThan {
		a.Index++
		a.Index++
		term2, err := a.GetExpression()
		if err != nil {
			return nil, err
		}

		var ct ConditionType
		switch nextToken.Type {
		case token.LThan:
			ct = LessThan

		case token.GThan:
			ct = GreaterThan

		default:
			return nil, errors.New("Could not deduce condition type")
		}

		fmt.Println("therese a conditional")
		return &Condition{
			Type:  ct,
			Left:  term,
			Right: term2,
		}, nil
	}

	if a.Index+1 < len(a.Tokens)-1 {
		for a.Tokens[a.Index+1].Type == token.SecOp {
			a.Index++

			operand := a.Tokens[a.Index].Value.String

			a.Index++
			term2, err := a.GetExpression()
			if err != nil {
				return nil, err
			}

			fmt.Println("term2", term2, operand)

			term, err = NewBinaryOperation(a.Tokens[a.Index], operand, term, term2)
			if err != nil {
				return nil, err
			}

			if a.Index >= len(a.Tokens)-1 {
				break
			}
		}
	}

	// FIXME: should probably check for secondary operations right here

	return term, nil
}

func (a *ASTBuilder) GetGroup() (*Group, error) {
	// `(` [ expression ]* `)`
	fmt.Println("getting group", a.Tokens[a.Index])

	if a.Tokens[a.Index].Type != token.LParen {
		return nil, errors.New("Function declaration requires left paren after function identifier")
	}

	groupToken := a.Tokens[a.Index]

	elements := []Expression{}

	a.Index++
	for a.Tokens[a.Index].Type != token.RParen {
		expr, err := a.GetExpression()
		if err != nil {
			return nil, err
		}

		fmt.Println("hey its me", expr, err)
		elements = append(elements, expr)

		a.Index++
	}

	return &Group{
		// Not sure if we should create a `group` token
		Token: groupToken,
		// Not sure what type to put, maybe make a group type?
		// TypeOf: Type(0),
		Elements: elements,
	}, nil
}

func (a *ASTBuilder) GetIf() (*IfElse, error) {
	var err error

	ifElse := IfElse{
		Token: a.Tokens[a.Index],
	}

	a.Index++
	// FIXME: make a GetCondition() function later m8
	ifElse.Condition, err = a.GetExpression()
	if err != nil {
		return nil, err
	}

	a.Index++
	ifElse.Body, err = a.GetBlock()
	if err != nil {
		return nil, err
	}
	fmt.Println("ifElse", ifElse)
	// Check for an else branch
	elseToken := a.Tokens[a.Index+1]
	if elseToken.Type == token.Else {
		a.Index++

		// Check whether there is an if or just another block
		switch a.Tokens[a.Index+1].Type {
		case token.If:
			ifElse.Else, err = a.GetIf()
			if err != nil {
				return nil, err
			}

		case token.LBrace:
			a.Index++

			ifElse.Else = &IfElse{
				Token: elseToken,
			}

			ifElse.Else.Body, err = a.GetBlock()
			if err != nil {
				return nil, err
			}

		default:
			return nil, errors.New("Empty else branch")
		}
	}

	return &ifElse, nil
}

func (a *ASTBuilder) GetBlock() (*Block, error) {
	if a.Tokens[a.Index].Type != token.LBrace {
		return nil, errors.New("Could not find block opening")
	}

	lb := a.Tokens[a.Index]

	statements := []Statement{}

	a.Index++
	for a.Tokens[a.Index].Type != token.RBrace {
		stmt, err := a.GetStatement()
		if err != nil {
			return nil, err
		}

		statements = append(statements, stmt)

		a.Index++
	}

	return &Block{
		Token:      lb,
		Statements: statements,
		// Scope: //FIXME: I don't think this should be here
	}, nil
}

// GetStatement needs to switch and capture these:
//	- assignment
//		- type
//		- ident
//	- block
//	- call
//		- ident
//	- func / fn
//	- if/else
//	- loop
//	- return
func (a *ASTBuilder) GetStatement() (Statement, error) {
	typeOf := ""
	currentToken := a.Tokens[a.Index]

	switch currentToken.Type {
	case token.Separator:
		// TODO: just skip the separator for now
		a.Index++
		return a.GetStatement()

	case token.Type:
		// Look for an ident as the next thing for now
		// fallthrough to the next block for now
		typeOf = currentToken.Value.String
		a.Index++

		// Expect an ident to always follow a token for now
		fallthrough

	case token.Ident:
		// Here we will want to look at what is next and handle it
		// If it is an assignment statment then we are looking for an expression afterwards

		// FIXME: need to implement Type() so that we can get the var type
		ident, err := NewIdent(a.Tokens[a.Index], typeOf)
		if err != nil {
			return nil, err
		}

		a.Index++

		switch a.Tokens[a.Index].Type {
		case token.Assign:
			assignmentToken := a.Tokens[a.Index]

			a.Index++
			expr, err := a.GetExpression()
			if err != nil {
				return nil, err
			}
			fmt.Println("expr", expr)

			// TODO: figure out why i put this here
			if expr == nil {
				return nil, nil
			}

			// TODO: could make a new boolean assignment here?
			as, err := NewAssignment(assignmentToken, ident, expr)
			if err != nil {
				return nil, err
			}

			if typeOf != "" {
				as.SetDeclaration(true)
			}

			// TODO: add statement here later
			return as, nil

		// This is a function call as a statement
		// TODO: implement function call expressions
		case token.LParen:
			// Get the group for the args
			args, err := a.GetGroup()
			if err != nil {
				return nil, err
			}

			return &Call{
				Token:     ident.Token,
				Ident:     ident,
				Arguments: args,
			}, nil
		}

		return nil, errors.Errorf("Expected assignment token, got %+v", a.Tokens[a.Index])

	case token.LBrace:
		// Here we will want to recursively call GetStatement()
		// however, a block should be able to be parsed for an expression as well
		return a.GetBlock()

	// 	// This one will have to be figured out when parsing the ident
	// case token.Call:

	// TODO: break this out into the individual keywords
	// - switch, etc
	// case token.Keyword:
	// 	// switch

	case token.Function:
		// Next things we look for after the Function token is:
		//	[ ident ] [ group ] { group } [ block ]
		fmt.Println("Found a function token")

		// Get the function token
		functionToken := a.Tokens[a.Index]

		// Get the ident token
		a.Index++
		identToken := a.Tokens[a.Index]

		// Get the group for the args
		a.Index++
		args, err := a.GetGroup()
		if err != nil {
			return nil, err
		}

		// FIXME: skip getting the returns for now

		// Get the body of the function
		// the body is essentially just a list of statements
		// this is the exact same as a file in our definition

		a.Index++
		block, err := a.GetBlock()
		if err != nil {
			return nil, err
		}

		return NewFunction(functionToken, identToken, args, block)

	// 	// TODO: create this token
	case token.If:
		// TODO:
		// look for a conditional/expression
		// get a block
		// check for an else
		// if theres an else, look for another if or a block

		return a.GetIf()

	// // FIXME: maybe this needs to switch to token.Loop later on
	case token.For:
		fmt.Println("we found a for loop")
		// We need to be able to parse different types of loops here:
		// - standard loops
		// - preposition loops

		// Save the `for` token
		forToken := a.Tokens[a.Index]

		a.Index++

		// Figure out what type of loop it is by the next token
		switch a.Tokens[a.Index].Type {
		// support declaring static typed variables as well
		case token.Type:
			typeOf = a.Tokens[a.Index].Value.String
			a.Index++
			fallthrough

		case token.Ident:
			ident, err := NewIdent(a.Tokens[a.Index], typeOf)
			if err != nil {
				return nil, err
			}

			// Look ahead one token to determine what type of loop it is
			switch a.Tokens[a.Index+1].Type {

			// For now just keep it like this:
			// Later we can change it to actually get specific nodes:
			// like:
			//	- GetAssignmentStatement()
			//	- GetConditionalExpression()
			//	- GetArithmeticExpression()

			case token.Assign:
				stmt, err := a.GetStatement()
				if err != nil {
					return nil, err
				}

				a.Index++

				// For now just check for the separator here
				if a.Tokens[a.Index].Type == token.Separator {
					a.Index++
				}

				expr, err := a.GetExpression()
				if err != nil {
					return nil, err
				}

				a.Index++

				// For now just check for the separator here
				if a.Tokens[a.Index].Type == token.Separator {
					a.Index++
				}

				expr2, err := a.GetExpression()
				if err != nil {
					return nil, err
				}

				body, err := a.GetBlock()
				if err != nil {
					return nil, err
				}

				// FIXME: should make a new function for this
				return &Loop{
					Token: forToken,
					Type:  StdFor,
					Init:  stmt,
					Cond:  expr,
					Post:  expr2,
					Body:  body,
				}, nil

			case token.Keyword:
				a.Index++
				preposition := a.Tokens[a.Index]

				a.Index++
				expr, err := a.GetExpression()
				if err != nil {
					return nil, err
				}
				fmt.Println("expr me", expr)

				iter, err := NewIterable(forToken, preposition, ident, expr)
				if err != nil {
					return nil, err
				}

				a.Index++

				body, err := a.GetBlock()
				if err != nil {
					return nil, err
				}

				// FIXME: should make a new function for this
				return &Loop{
					Token: forToken,
					Type:  StdFor,
					Iter:  iter,
					Body:  body,
				}, nil
			}
		}

		return nil, errors.New("Could not parse loop")

	case token.Return:
		// For now just look for a single expression afterwards
		a.Index++
		expr, err := a.GetExpression()
		if err != nil {
			return nil, err
		}

		fmt.Println("return return")

		return NewReturn(token.Token{}, expr), nil

	default:
		return nil, errors.Errorf("Could not get statement from token: %+v", currentToken)
	}

	return nil, errors.Errorf("Could not deduce statement starting at: %+v", a.Tokens[a.Index])
}

// BuildAST builds an AST from the tokens provided by the lexer
func (a *ASTBuilder) BuildAST() (*Program, error) {
	p := NewProgram()

	// FIXME: Spoof this name for now
	file := NewFile("main.expr")

	for {
		// We know that the file can only consist of statements
		stmt, err := a.GetStatement()
		if err != nil {
			return nil, err
		}

		file.AddStatement(stmt)

		a.Index++

		if a.Index > len(a.Tokens)-1 {
			break
		}
	}

	p.AddFile(file)

	return p, nil
}

func CompressTokens(lexTokens []token.Token) ([]token.Token, error) {
	compressedTokens := []token.Token{}

	alreadyChecked := false

	for i := 0; i < len(lexTokens)-1; i++ {
		fmt.Println("i", lexTokens[i])

		// This needs to be simplified
		if lexTokens[i].Type == "ASSIGN" || lexTokens[i].Type == "SEC_OP" || lexTokens[i].Type == "PRI_OP" && lexTokens[i+1].Type == "ASSIGN" || lexTokens[i+1].Type == "SEC_OP" || lexTokens[i+1].Type == "PRI_OP" {
			compressedToken, ok := token.TokenMap[lexTokens[i].Value.String+lexTokens[i+1].Value.String]
			fmt.Println("added \"" + lexTokens[i].Value.String + lexTokens[i+1].Value.String + "\"")
			if ok {
				compressedTokens = append(compressedTokens, compressedToken)
				i++

				// If we were able to combine the last two tokens and make a new one, mark it
				if i == len(lexTokens)-1 {
					alreadyChecked = true
				}

				continue
			}
		}

		// Filter out the white space
		if lexTokens[i].Type == "WS" {
			continue
		}

		compressedTokens = append(compressedTokens, lexTokens[i])
	}

	// If it hasn't been already checked and the last token is not a white space, then append it
	if !alreadyChecked && lexTokens[len(lexTokens)-1].Type != "WS" {
		compressedTokens = append(compressedTokens, lexTokens[len(lexTokens)-1])
	}

	return compressedTokens, nil
}
