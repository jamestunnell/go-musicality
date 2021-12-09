package util

import (
	"fmt"
	"math/big"
)

type NonPositiveError struct {
	ValueStr string
}

func NewNonPositiveFloatError(val float64) *NonPositiveError {
	return &NonPositiveError{ValueStr: fmt.Sprintf("%e", val)}
}

func NewNonPositiveRatError(rat *big.Rat) *NonPositiveError {
	return &NonPositiveError{ValueStr: rat.String()}
}

func (e *NonPositiveError) Error() string {
	return fmt.Sprintf("value %s is not positive", e.ValueStr)
}
