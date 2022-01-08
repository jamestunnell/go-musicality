package validation

import "fmt"

type ErrNotLessEqualOne struct {
	Name  string
	Value float64
}

func NewErrNotLessEqualOne(name string, val float64) *ErrNotLessEqualOne {
	return &ErrNotLessEqualOne{
		Name:  name,
		Value: val,
	}
}

func (err *ErrNotLessEqualOne) Error() string {
	return fmt.Sprintf("%s %e is not <= 1", err.Name, err.Value)
}
