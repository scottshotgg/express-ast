package ast

type LiteralType int

const (
	IntType LiteralType = iota + 1
	FloatType
	CharType
	StringType
	BoolType
)

// FIXME: gotta implement literal on all of these
type Literal interface {
	Type() LiteralType
}

type IntLiteral struct {
	Token Token
	Type  LiteralType
	Value int
}

func (_ *IntLiteral) expressionNode()       {}
func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }

type FloatLiteral struct {
	Token Token
	Type  LiteralType
	Value float64
}

func (_ *FloatLiteral) expressionNode()       {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }

type CharLiteral struct {
	Token Token
	Type  LiteralType
	Value [1]rune
}

func (_ *CharLiteral) expressionNode()       {}
func (cl *CharLiteral) TokenLiteral() string { return cl.Token.Literal }

type StringLiteral struct {
	Token Token
	Type  LiteralType
	Value string
}

func (_ *StringLiteral) expressionNode()       {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

type BoolLiteral struct {
	Token Token
	Type  LiteralType
	Value bool
}

func (_ *BoolLiteral) expressionNode()       {}
func (bl *BoolLiteral) TokenLiteral() string { return bl.Token.Literal }
