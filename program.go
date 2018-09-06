package ast

// Program represents the following form:
// [ statement ]*
type Program struct {
	// We'll wanna put 'File' in between Program and Block at some point
	Files []File
}

// Length returns the length of files in the program
func (p *Program) Length() int {
	return len(p.Files)
}
