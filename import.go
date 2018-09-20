package ast

type Import struct {
	Token Token
	Name *Ident
	Path string
}

// Implement Node and Statement

func (i *Import) statementNode() {}

// TokenLiteral returns the literal value of the token
func (i *Import) TokenLiteral() string { return s.Token.Literal }