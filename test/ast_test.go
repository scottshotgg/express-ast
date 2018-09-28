package ast_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/scottshotgg/express-ast"
)

var (
	c = &spew.ConfigState{
		Indent:                "\t",
		DisableMethods:        true,
		DisablePointerMethods: true,
		SortKeys:              true,
		SpewKeys:              true,
	}
)

// TODO: FIXME: make function to do all this shit like NewInt, etc
// take it from the token2 library

/*
	// The following ast represents a declared function from a file named "main.expr"

	myAdder := func(int a, float b) (float) { return a + b }

*/

func TestAST(t *testing.T) {
	a := ast.Program{
		Files: map[string]*ast.File{
			"main.expr": &ast.File{
				Statements: []ast.Statement{
					&ast.Assignment{
						Declaration: true,
						Inferred:    true,
						Lhs: &ast.Ident{
							Type: ast.NewFloatType(),
							Name: "myAdder",
						},
						Rhs: &ast.Function{
							Name: "myFunction",
							Arguments: &ast.Group{
								Elements: []ast.Expression{
									ast.NewInt(ast.Token{}, 0),
									ast.NewFloat(ast.Token{}, 0),
								},
							},
							Returns: &ast.Group{
								Elements: []ast.Expression{
									// Not sure if this should be an anonymous ident with a name,
									// without a name, or if ast.Type should just implement Expression
									&ast.Ident{
										Type: ast.NewFloatType(),
									},
								},
							},
							Body: ast.Block{
								Statements: []ast.Statement{
									&ast.Return{
										Value: &ast.BinaryOperation{
											Op: ast.AdditionBinaryOp,
											LeftNode: &ast.Ident{
												Name: "a",
												Type: ast.NewIntType(),
											},
											RightNode: &ast.Ident{
												Name: "b",
												Type: ast.NewFloatType(),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	c.Dump("AST:", a)

	astJSON, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("\n", string(astJSON))
}
