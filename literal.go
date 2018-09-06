package ast

// LiteralType encompasses all types of literals
type LiteralType int

const (
	// IntType denotes an integer literal type
	IntType LiteralType = iota + 1

	// FloatType denotes a float literal type
	FloatType

	// CharType denotes a char literal type
	CharType

	// StringType denotes a string literal type
	StringType

	// BoolType denotes a bool literal type
	BoolType

	// VarType denotes a var literal type
	VarType
)

// Literal is an abstract types that represents a literal value, in constrast with a value-producer, such as an expression
type Literal interface {
	Type() LiteralType
	ActingType() LiteralType
}

// IntLiteral represents any non floating-point number
type IntLiteral struct {
	Token      Token
	Type       LiteralType
	ActingType LiteralType
	Value      int
}

func (il *IntLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }

// FloatLiteral represents any floating point number
type FloatLiteral struct {
	Token      Token
	Type       LiteralType
	ActingType LiteralType
	Value      float64
}

func (fl *FloatLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }

// CharLiteral represents a single-character capped string
type CharLiteral struct {
	Token      Token
	Type       LiteralType
	ActingType LiteralType
	Value      [1]rune
}

func (cl *CharLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (cl *CharLiteral) TokenLiteral() string { return cl.Token.Literal }

// StringLiteral respresents a quoted body of text in the form of:
// `"` [ _text_ ] `"`
type StringLiteral struct {
	Token      Token
	Type       LiteralType
	ActingType LiteralType
	Value      string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// BoolLiteral represents a variable that is restricted to either a true or false value
type BoolLiteral struct {
	Token      Token
	Type       LiteralType
	ActingType LiteralType
	Value      bool
}

func (bl *BoolLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (bl *BoolLiteral) TokenLiteral() string { return bl.Token.Literal }

// VarLiteral represents a variable that is restricted to either a true or false value
type VarLiteral struct {
	Token      Token
	Type       LiteralType
	ActingType LiteralType
	Value      bool
}

func (vl *VarLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (vl *VarLiteral) TokenLiteral() string { return vl.Token.Literal }
