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

type Int struct {
	Token Token
	Type  LiteralType
	Value int
}

func (_ *IntLiteral) expressionNode()      {}
func (l *IntLiteral) TokenLiteral() string { return il.Token.Literal }

type Float struct {
	Token Token
	Type  LiteralType
	Value float64
}

func (_ *FloatLiteral) expressionNode()       {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }

type Char struct {
	Token Token
	Type  LiteralType
	Value [1]rune
}

func (_ *CharLiteral) expressionNode()       {}
func (cl *CharLiteral) TokenLiteral() string { return cl.Token.Literal }

type String struct {
	Token Token
	Type  LiteralType
	Value string
}

func (_ *StringLiteral) expressionNode()       {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

type Bool struct {
	Token Token
	Type  LiteralType
	Value bool
}

func (_ *BoolLiteral) expressionNode()       {}
func (bl *BoolLiteral) TokenLiteral() string { return bl.Token.Literal }
