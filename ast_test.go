package ast_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	ast "github.com/scottshotgg/express-ast"
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

func TestAST(t *testing.T) {
	a := ast.Program{
		Files: []ast.File{
			ast.File{
				Statements: []ast.Statement{
					&ast.Block{},
					&ast.Function{
						Name: "myFunction",
						Arguments: []ast.Expression{
							&ast.IntLiteral{
								Type:  ast.IntType,
								Value: 444,
							},
							&ast.FloatLiteral{
								Type:  ast.FloatType,
								Value: 1.0,
							},
						},
						Returns: []ast.Expression{
							&ast.BinaryOperation{
								Kind: ast.AdditionBinaryOp,
								LeftNode: &ast.IntLiteral{
									Type:  ast.IntType,
									Value: 2,
								},
								RightNode: &ast.FloatLiteral{
									Type:  ast.FloatType,
									Value: 1.0,
								},
							},
						},
						Body: ast.Block{
							Statements: []ast.Statement{
								&ast.Return{
									Value: []ast.Expression{
										&ast.BinaryOperation{
											Kind: ast.AdditionBinaryOp,
											LeftNode: &ast.IntLiteral{
												Type:  ast.IntType,
												Value: 2,
											},
											RightNode: &ast.FloatLiteral{
												Type:  ast.FloatType,
												Value: 1.0,
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
