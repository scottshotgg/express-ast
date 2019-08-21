package builder_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/express-ast/builder"
	lex "github.com/scottshotgg/express-lex"
)

func TestNew(t *testing.T) {
	var lexer, err = lex.NewFromFile("../test/test.expr")
	if err != nil {
		t.Fatalf("error %+v", err)
	}

	tokens, err := lexer.Lex()
	if err != nil {
		t.Fatalf("error %+v", err)
	}

	var b = builder.New(tokens)
	program, err := b.BuildAST()
	if err != nil {
		t.Fatalf("error %+v", err)
	}

	fmt.Println("program", program)
}

func TestParseAssignmentStatement(t *testing.T) {
	var lexer, err = lex.NewFromFile("../test/test.expr")
	if err != nil {
		t.Fatalf("error %+v", err)
	}

	tokens, err := lexer.Lex()
	if err != nil {
		t.Fatalf("error %+v", err)
	}

	var b = builder.New(tokens)
	program, err := b.BuildAST()
	if err != nil {
		t.Fatalf("error %+v", err)
	}

	fmt.Println("program", program)

	b.ParseAssignmentStatement()
}
