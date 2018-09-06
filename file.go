package ast

// File represents a file that is being compiled
type File struct {
	Statements []Statement
}

// Length returns the list of statements in the file
func (f *File) Length() int {
	// TODO: this will have to do something to recurse down the chain and figure out blocks and add that to the total
	// return len(f.Statements)

	// for _, stmt := range f.Statements {
	// 	// TODO: statement should define a .Length() function that will return the length of the statement node
	// }

	return -1
}
