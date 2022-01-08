package validation

import (
	"fmt"
	"math/big"
)

type ErrNonPositive struct {
	Name, Value string
}

func NewErrNonPositiveInt(name string, val int) *ErrNonPositive {
	return &ErrNonPositive{Name: name, Value: fmt.Sprintf("%d", val)}
}

func NewErrNonPositiveUInt(name string, val uint) *ErrNonPositive {
	return &ErrNonPositive{Name: name, Value: fmt.Sprintf("%d", val)}
}

func NewErrNonPositiveFloat(name string, val float64) *ErrNonPositive {
	return &ErrNonPositive{Name: name, Value: fmt.Sprintf("%e", val)}
}

func NewErrNonPositiveRat(name string, rat *big.Rat) *ErrNonPositive {
	return &ErrNonPositive{Name: name, Value: rat.String()}
}

func (e *ErrNonPositive) Error() string {
	return fmt.Sprintf("%s %s is not positive", e.Name, e.Value)
}
