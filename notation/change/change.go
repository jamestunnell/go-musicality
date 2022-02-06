package change

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/validation"
)

var zero = big.NewRat(0, 1)

type Change struct {
	EndValue float64
	Duration *big.Rat
}

func New(endVal float64, dur *big.Rat) *Change {
	return &Change{
		EndValue: endVal,
		Duration: dur,
	}
}

func NewImmediate(endVal float64) *Change {
	return New(endVal, big.NewRat(0, 1))
}

func (c *Change) Validate() *validation.Result {
	errs := []error{}

	if c.Duration.Cmp(zero) == -1 {
		errs = append(errs, fmt.Errorf("duration %v is negative", c.Duration))
	}

	if len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "change",
		Errors:     errs,
		SubResults: []*validation.Result{},
	}
}
