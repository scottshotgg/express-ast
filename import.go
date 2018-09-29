package ast

import "github.com/scottshotgg/express-token"

// Import is an import statement in the form of:
// `import` [ string_lit ]
type Import struct {
	Token token.Token
	Name  *Ident
	Path  string
}

// Implement Node and Statement
func (i *Import) statementNode() {}

// TokenLiteral returns the literal value of the token
func (i *Import) TokenLiteral() token.Token { return i.Token }

func (i *Import) Kind() NodeType { return ImportNode }
