package ast

// Program represents the following form:
// [ statement ]*
type Program struct {
	Files []File
}

// Length returns the length of files in the program
func (p *Program) Length() int {
	return len(p.Files)
}
