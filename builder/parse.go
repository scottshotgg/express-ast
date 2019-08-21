package builder

import (
	"fmt"

	"github.com/pkg/errors"
	types "github.com/scottshotgg/express-ast/types"
	token "github.com/scottshotgg/express-token"
)

var (
	ErrNoEqualsFoundAfterIdent = errors.New("No equals found after ident in assignment")
)

// TODO: types are expressions, keep it that way

func (b *Builder) ParseAssignmentStatement() (*types.Assignment, error) {
	// into: [expr] = [expr]
	// Check that the next token is an ident
	if b.tokens[b.index].Type != token.Ident {
		return nil, b.AppendTokenToError("Could not get assignment statement without ident")
	}

	ident, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	var a = types.Assignment{
		Token: b.tokens[b.index],
		Left:  ident,
	}

	if b.index > len(b.tokens)-1 {
		return &a, nil
	}

	// Increment over the ident token
	b.index++

	// TODO: come back to this
	// if b.tokens[b.index].Type == token.Set {
	// 	return b.ParseSet(ident)
	// }

	// Check for the equals token
	if b.tokens[b.index].Type != token.Assign {
		if ident.Type == "call" {
			return ident, nil
		}

		// TODO: this is where we need to check for `:`

		// return nil, b.AppendTokenToError(fmt.Sprintf("No equals found after ident in assignment: %+v", b.tokens[b.index]))
		// This need to return the token in case the parse needs to be recovered! Look at ParseEnumBlock for an example of parse recovery
		return ident, ErrNoEqualsFoundAfterIdent
	}

	// Increment over the equals
	b.index++

	// Parse the right hand side
	expr, err := b.ParseExpression()
	if err != nil {
		return nil, err
	}

	// Increment over the first part of the expression
	b.index++

	var node = &types.Node{
		Type:  "assignment",
		Left:  ident,
		Right: expr,
	}

	// Do one pass for declarations, and check that the assignments
	// and usages corraborate in the type checker
	// return node, b.ScopeTree.Assign(node)
	return node, nil
}

// TODO: what if types were expressions ...

// ParseStatement ** does ** not look ahead
// We need to make ParseDeclarationStatement last and just accept the type until we get to the semantic parser
func (b *Builder) ParseStatement() (types.Statement, error) {
	switch b.tokens[b.index].Type {

	// case token.Launch:
	// 	return b.ParseLaunchStatement()

	// case token.Defer:
	// 	return b.ParseDeferStatement()

	// case token.Enum:
	// 	return b.ParseEnumBlockStatement()

	// case token.Map:
	// 	return b.ParseMapStatement()

	// case token.PriOp:
	// 	return b.ParseDerefStatement()

	// case token.Package:
	// 	return b.ParsePackageStatement()

	// case token.Import:
	// 	return b.ParseImportStatement()

	// // case token.Use:
	// // 	return b.ParseUseStatement()

	// case token.Include:
	// 	return b.ParseIncludeStatement()

	// case token.TypeDef:
	// 	return b.ParseTypeDeclarationStatement()

	// case token.Struct:
	// 	return b.ParseStructStatement()

	// case token.Object:
	// 	return b.ParseObjectStatement()

	// // case token.C:
	// // 	return b.ParseCBlock()

	// case token.Type:
	// 	// // Struct is a keyword and a type so if we get it as a type statment
	// 	// // then we need to divert the parsing
	// 	// if b.tokens[b.index].Value.String == token.StructType {
	// 	// 	return b.ParseStructStatement()
	// 	// }

	// 	return b.ParseDeclarationStatement(nil)

	// // For literal and idents, we will need to figure out what
	// // kind of statement it is
	// case token.Literal:
	// 	return b.ParseLiteralStatement()

	case token.Ident:
		var expr, err = b.ParseExpression()
		fmt.Println("expr, err", expr, err)

		return b.ParseAssignmentStatement()

	// case token.Function:
	// 	return b.ParseFunctionStatement()

	// case token.LBrace:
	// 	return b.ParseBlockStatement()

	// case token.Let:
	// 	return b.ParseLetStatement()

	// case token.If:
	// 	return b.ParseIfStatement()

	// case token.For:
	// 	return b.ParseForStatement()

	// case token.Return:
	// 	return b.ParseReturnStatement()

	default:
		return nil, b.AppendTokenToError(fmt.Sprintf("Could not create statement from: %+v", b.tokens[b.index].Type))
	}
}

// func (b *Builder) ParseLaunchStatement() (*types.Launch, error) {
// 	// Check ourselves ...
// 	if b.tokens[b.index].Type != token.Launch {
// 		return nil, b.AppendTokenToError("Could not get launch statement")
// 	}

// 	// Step over the import token
// 	b.index++

// 	// Might need to make this an explicit function call later
// 	expr, err := b.ParseStatement()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &types.Node{
// 		Type: "launch",
// 		Left: expr,
// 	}, nil
// }

// func addDeferDeclarationToBlock(n *types.Node) *types.Node {
// 	var stmts, ok = n.Value.([]*builder.types.Node)
// 	stmts = append([]*builder.types.Node(&types.Node{
// 		Type: "defer"
// 	}, stmts...))
// }

// func (b *Builder) parseFileImport(filename string) (*types.Node, *ScopeTree, error) {
// 	// var path, err = os.Getwd()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	source, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	fmt.Println("source", string(source))

// 	// Lex and tokenize the source code
// 	tokens, err := lex.New(string(source)).Lex()
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// Compress certain tokens;
// 	tokens, err = ast.CompressTokens(tokens)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// Build the AST
// 	b2 := New(tokens)
// 	ast, err := b2.BuildAST()
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// fmt.Printf("ast %+v\n", ast.Value.([]*types.Node)[0].Left.Value.(string))

// 	// TODO: extremely unsafe, fix this
// 	return ast, b2.ScopeTree, nil
// }

// TODO: change this to ParseKeyValueStatement
// func (b *Builder) ParseLiteralStatement() (*types.Node, error) {
// 	// Parse an expession
// 	// check the next token for a `:`
// 	// Parse another expression
// 	// Return a key-value pair

// 	// Get the expression
// 	var left, err = b.ParseExpression()
// 	if err != nil {
// 		return nil, err
// 	}

// 	b.index++

// 	switch b.tokens[b.index].Type {
// 	case token.Set:
// 		return b.ParseSet(left)

// 	default:
// 		return nil, errors.Errorf("ParseLiteralStatement not implemented for: %+v", b.tokens[b.index].Type)
// 	}
// }

// func (b *Builder) ParseObjectStatement() (*types.Object, error) {
// 	// Check ourselves ...
// 	if b.tokens[b.index].Type != token.Object {
// 		return nil, b.AppendTokenToError("Could not get object declaration statement")
// 	}

// 	// Skip over the `struct` token
// 	b.index++

// 	// Create the ident
// 	ident, err := b.ParseExpression()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Increment over the ident token
// 	b.index++

// 	// Create a new child scope for the function
// 	b.ScopeTree, err = b.ScopeTree.NewChildScope(ident.Value.(string))
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check for the equals token
// 	if b.tokens[b.index].Type != token.Assign {
// 		return nil, b.AppendTokenToError("No equals found after ident in object def")
// 	}

// 	// Increment over the equals
// 	b.index++

// 	// Parse the right hand side
// 	body, err := b.ParseBlockStatement()
// 	if err != nil {
// 		return nil, err
// 	}

// 	body.Kind = "object"

// 	// _, err = b.AddStructured(ident.Value.(string), body)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// Object does not get a type ... yet
// 	// var v = &TypeValue{
// 	// 	Composite: true,
// 	// 	Type:      StruturedValue,
// 	// 	Kind:      body.Kind,
// 	// }
// 	// v.Props, err = b.extractPropsFromComposite(body)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// // Increment over the first part of the expression
// 	// b.index++

// 	// Assign our scope back to the current one
// 	b.ScopeTree, err = b.ScopeTree.Leave()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Again about the object not creating a type ...
// 	// err = b.ScopeTree.NewType(ident.Value.(string), v)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	return &types.Node{
// 		Type:  "object",
// 		Left:  ident,
// 		Right: body,
// 	}, nil
// }

// func (b *Builder) ParseDeclarationStatement() (*types.Declaration, error) {
// 	var typeOf, err = b.ParseType()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("typeOf outside", typeOf)

// 	// Check that the next token is an ident
// 	if b.tokens[b.index].Type != token.Ident {
// 		return nil, b.AppendTokenToError("Could not get ident in declaration statement")
// 	}

// 	// Create the ident
// 	ident, err := b.ParseExpression()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var typeString = typeOf.Value.(string)
// 	if typeString == "map" || typeString == "object" || typeString == "struct" {
// 		b.ScopeTree, err = b.ScopeTree.NewChildScope(ident.Value.(string))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	// // Check the scope map to make sure this hasn't been declared for the current scope
// 	// var node = b.ScopeTree.Local(ident.Value.(string))

// 	// // If the return value isn't nil then that means we found something in the local scope
// 	// if node != nil {
// 	// 	return nil, errors.Errorf("Variable already declared: %+v\n", node)
// 	// }

// 	// err = b.ScopeTree.Declare(ident)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// Increment over the ident token
// 	b.index++

// 	// Check for the equals token
// 	if b.tokens[b.index].Type != token.Assign {
// 		return &types.Node{
// 			Type:  "decl",
// 			Value: typeOf,
// 			Left:  ident,
// 		}, nil

// 		// return nil, errors.New("No equals found after ident")
// 	}

// 	// Increment over the equals
// 	b.index++

// 	// Parse the right hand side
// 	expr, err := b.ParseExpression()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Increment over the first part of the expression
// 	b.index++

// 	// Leave the scope if we entered it above
// 	if typeString == "map" || typeString == "object" || typeString == "struct" {
// 		// Assign our scope back to the current one
// 		b.ScopeTree, err = b.ScopeTree.Leave()
// 		if err != nil {
// 			return nil, err
// 		}

// 		if typeString == "struct" {
// 			var v = &TypeValue{
// 				Composite: true,
// 				Type:      StruturedValue,
// 				Kind:      expr.Kind,
// 			}

// 			v.Props, err = b.extractPropsFromComposite(expr)
// 			if err != nil {
// 				return nil, err
// 			}

// 			err = b.ScopeTree.NewType(ident.Value.(string), v)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		// Could defer this and then exit when we error?
// 	}

// 	var node = &types.Node{
// 		Type:  "decl",
// 		Value: typeOf,
// 		Left:  ident,
// 		Right: expr,
// 	}

// 	return node, b.ScopeTree.Declare(node)
// }
