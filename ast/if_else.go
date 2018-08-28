package ast

// if [ condition ] [ block ] { [ else ] [ statement ] }
type IfElse struct {
	Token         Token
	IfCondition   *Condition
	If            *Block
	ElseCondition *Condition
	// TODO: Hmmm this is supposed to only be a block or another if statement
	// but should we try to bound it?
	Else *Statement
}

func (_ *IfElse) statementNode()        {}
func (ie *IfElse) TokenLiteral() string { return ie.Token.Literal }
