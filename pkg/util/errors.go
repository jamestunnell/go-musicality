package util

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/duration"
)

type NonPositiveError struct {
	ValueStr string
}

func NewNonPositiveFloatError(val float64) *NonPositiveError {
	return &NonPositiveError{ValueStr: fmt.Sprintf("%e", val)}
}

func NewNonPositiveDurationError(d *duration.Duration) *NonPositiveError {
	return &NonPositiveError{ValueStr: d.String()}
}

func (e *NonPositiveError) Error() string {
	return fmt.Sprintf("value %s is not positive", e.ValueStr)
}
