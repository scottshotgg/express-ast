package ast

// ConditionType encompasses all types of conditions
type ConditionType int

const (
	// IsEqualTo is the == operator
	IsEqualTo ConditionType = iota + 1

	// LessThan is the < operator
	LessThan

	// GreaterThan is the > operator
	GreaterThan

	// Not is the ! operator
	Not

	// Or is the || operator
	Or

	// And is the && operator
	And
)

// Condition represents the following form:
// [ expression ] [ condition_op ] [ expression ]
type Condition struct {
	Token Token
	Type  ConditionType
	Left  *Expression
	Right *Expression
	Value bool
}

func (c *Condition) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (c *Condition) TokenLiteral() string { return c.Token.Literal }
