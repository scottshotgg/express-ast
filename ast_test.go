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

	myAdder := function(int a, float b) (float) { return a + b }

*/

func TestAST(t *testing.T) {
	a := ast.Program{
		Files: map[string]ast.File{
			"main.expr": ast.File{
				Statements: []ast.Statement{
					&ast.Assignment{
						Declaration: true,
						Inferred:    true,
						Ident: ast.Ident{
							Type: ast.Type{
								Type: ast.FunctionType,
							},
							Name: "myAdder",
						},
						Value: &ast.Function{
							Name: "myFunction",
							Arguments: &ast.Group{
								Elements: []ast.Expression{
									&ast.IntLiteral{
										Type: ast.Type{
											Type: ast.IntType,
										},
									},
									&ast.FloatLiteral{
										Type: ast.Type{
											Type: ast.FloatType,
										},
									},
								},
							},
							Returns: &ast.Group{
								Elements: []ast.Expression{
									// Not sure if this should be an anonymous ident with a name,
									// without a name, or if ast.Type should just implement Expression
									&ast.Ident{
										Type: ast.Type{
											Type: ast.FloatType,
										},
									},
								},
							},
							Body: ast.Block{
								Statements: []ast.Statement{
									&ast.Return{
										Value: []ast.Expression{
											&ast.BinaryOperation{
												Kind: ast.AdditionBinaryOp,
												LeftNode: &ast.Ident{
													Name: "a",
													Type: ast.Type{
														Type: ast.IntType,
													},
												},
												RightNode: &ast.Ident{
													Name: "b",
													Type: ast.Type{
														Type: ast.FloatType,
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
