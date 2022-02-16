package validation

import (
	"fmt"
)

type ErrNonPositive struct {
	Name  string
	Value interface{}
}

func NewErrNonPositive(name string, val interface{}) *ErrNonPositive {
	return &ErrNonPositive{Name: name, Value: val}
}

func (e *ErrNonPositive) Error() string {
	return fmt.Sprintf("%s %v is not positive", e.Name, e.Value)
}
