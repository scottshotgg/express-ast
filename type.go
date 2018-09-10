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

	// ObjectType denotes an object literal type
	ObjectType

	// StructType denotes a struct literal type
	StructType

	// FunctionType denotes a function literal type
	FunctionType

	// UserDefinedType denotes a type user defined type
	UserDefinedType
)

// Type is used to specify a variable type
type Type struct {
	Name       string
	Type       LiteralType
	ShadowType *LiteralType
	UpgradesTo *LiteralType
}

var (
	nameToUserDefinedTypeMap = map[string]*Type{}
	idToUserDefinedTypeMap   = map[LiteralType]*Type{}

	// Any type ID greater than 99 is a user defined type
	typeIndex LiteralType = 99

	// UpgradableTypesMap allows definitions of upgradable types
	UpgradableTypesMap = map[LiteralType]LiteralType{
		IntType:    FloatType,
		CharType:   StringType,
		StructType: ObjectType,
	}
)

// DeclareUserDefinedType declares a user defined type in the
// type map and returns a type ID
func DeclareUserDefinedType(udt *Type) LiteralType {
	typeIndex++

	udt.Type = typeIndex

	nameToUserDefinedTypeMap[udt.Name] = udt
	idToUserDefinedTypeMap[udt.Type] = udt

	return typeIndex
}
