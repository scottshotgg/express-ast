package builder

import (
	"github.com/pkg/errors"
)

const (
	formatString = "; %+v"
)

var (
	ErrNotImplemented = errors.New("Not implemented")
	ErrMultDimArrInit = errors.New("Cannot use multiple expression inside array type initializer")
	ErrOutOfTokens    = errors.New("Out of tokens")
)

func (b *Builder) AppendTokenToError(errText string) error {
	if b.index < len(b.tokens)-1 {
		return errors.Errorf(errText+formatString, b.tokens[b.index])
	}

	return errors.New(errText + "; No token to print")
}
