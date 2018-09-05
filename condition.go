package ast

type ConditionType int

const (
	EqualsCondition ConditionType = iota + 1
	LessThan
	GreaterThan
	Not
	Or
	And
)

type Condition struct {
	Token Token
	Type  ConditionType
	Left  *Expression
	Right *Expression
	Value bool
}

func (_ *Condition) expressionNode()      {}
func (c *Condition) TokenLiteral() string { return c.Token.Literal }
