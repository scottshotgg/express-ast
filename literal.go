package ast

// Literal is an abstract types that represents a literal value, in constrast with a value-producer, such as an expression
type Literal interface {
	Type() LiteralType
}

// Literals should have acting types and acting values that get set when the value is set

// IntLiteral represents any non floating-point number
type IntLiteral struct {
	Token Token
	Type  Type
	Value int
}

func (il *IntLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }

// FloatLiteral represents any floating point number
type FloatLiteral struct {
	Token Token
	Type  Type
	Value float64
}

func (fl *FloatLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }

// CharLiteral represents a single-character capped string:
// `'` [ _single_character_ ] `'`
type CharLiteral struct {
	Token Token
	Type  Type
	Value [1]rune
}

func (cl *CharLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (cl *CharLiteral) TokenLiteral() string { return cl.Token.Literal }

// StringLiteral represents a double quoted body of text:
// TODO: how to do a backtick quoted body of text
// `"` [ _text_ ] `"`
type StringLiteral struct {
	Token Token
	Type  Type
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// BoolLiteral represents a variable that is restricted to either a true or false value
type BoolLiteral struct {
	Token Token
	Type  Type
	Value bool
}

func (bl *BoolLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (bl *BoolLiteral) TokenLiteral() string { return bl.Token.Literal }

// VarLiteral represents a dynamically typed variable; it can hold anything
type VarLiteral struct {
	Token Token
	Type  Type
	// TODO: could either do it this way or this can reference another literal type
	Value interface{}
}

func (vl *VarLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (vl *VarLiteral) TokenLiteral() string { return vl.Token.Literal }

// ObjectLiteral represents a named block : this produces a variable
type ObjectLiteral struct {
	Token Token
	Type  Type
	// TODO: could either do it this way or make block implement literal and then it can be directly used as a literal
	Value Block
}

func (ol *ObjectLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (ol *ObjectLiteral) TokenLiteral() string { return ol.Token.Literal }

// StructLiteral represents a named object : this produces a type
// TODO: this might need to be moved to the type.go file
type StructLiteral struct {
	Token Token
	Type  Type
	// TODO: could either do it this way or make block implement literal and then it can be directly used as a literal
	Value Block
}

func (sl *StructLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (sl *StructLiteral) TokenLiteral() string { return sl.Token.Literal }

// FunctionLiteral represents a named object : this produces a type
type FunctionLiteral struct {
	Token Token
	Type  Type
	// TODO: could either do it this way or make block implement literal and then it can be directly used as a literal
	Value Block
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
