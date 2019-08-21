package builder

import (
	types "github.com/scottshotgg/express-ast/types"
	token "github.com/scottshotgg/express-token"
)

type Builder struct {
	tokens []token.Token
	index  int
}

func New(tokens []token.Token) *Builder {
	return &Builder{
		tokens: tokens,
	}
}

// BuildAST builds an AST from the tokens provided by the lexer
func (b *Builder) BuildAST() (*types.Program, error) {
	p := types.NewProgram()

	// FIXME: Spoof this name for now
	file := types.NewFile("main.expr")

	for {
		// We know that the file can only consist of statements
		stmt, err := b.GetStatement()
		if err != nil {
			return nil, err
		}

		file.AddStatement(stmt)

		b.index++

		if b.index > len(b.tokens)-1 {
			break
		}
	}

	p.AddFile(file)

	return p, nil
}
