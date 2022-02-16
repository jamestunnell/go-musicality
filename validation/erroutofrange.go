package validation

import (
	"fmt"
)

type ErrOutOfRange struct {
	Name            string
	Value, Min, Max interface{}
}

func NewErrOutOfRange(name string, val, min, max interface{}) *ErrOutOfRange {
	return &ErrOutOfRange{
		Name:  name,
		Value: val,
		Min:   min,
		Max:   max,
	}
}

func (e *ErrOutOfRange) Error() string {
	return fmt.Sprintf("%s %v is not in range [%v, %v]", e.Name, e.Value, e.Min, e.Max)
}
