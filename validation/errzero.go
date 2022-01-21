package validation

import (
	"fmt"
)

type ErrZero struct {
	Name string
}

func NewErrZero(name string) *ErrZero {
	return &ErrZero{
		Name: name,
	}
}

func (e *ErrZero) Error() string {
	return fmt.Sprintf("%s is 0", e.Name)
}
