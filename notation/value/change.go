package value

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/validation"
)

var zero = big.NewRat(0, 1)

type Change struct {
	Offset   *big.Rat
	EndValue float64
	Duration *big.Rat
}

func (c *Change) Validate() *validation.Result {
	errs := []error{}

	if c.Offset.Cmp(zero) == -1 {
		errs = append(errs, fmt.Errorf("offset %v is negative", c.Offset))
	}

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
