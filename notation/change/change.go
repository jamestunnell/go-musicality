package change

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/validation"
)

type Change struct {
	Offset   *big.Rat `json:"offset"`
	EndValue float64  `json:"endValue"`
	Duration *big.Rat `json:"duration"`
}

func New(offset *big.Rat, endVal float64, dur *big.Rat) *Change {
	return &Change{
		Offset:   offset,
		EndValue: endVal,
		Duration: dur,
	}
}

func NewImmediate(offset *big.Rat, endVal float64) *Change {
	return New(offset, endVal, new(big.Rat))
}

func (c *Change) Equal(c2 *Change) bool {
	return rat.IsEqual(c.Offset, c2.Offset) &&
		(c.EndValue == c2.EndValue) &&
		rat.IsEqual(c.Duration, c2.Duration)
}

func (c *Change) Validate(r ValueRange) *validation.Result {
	errs := []error{}

	if rat.IsNegative(c.Duration) {
		errs = append(errs, fmt.Errorf("duration %v is negative", c.Duration))
	}

	if !r.Includes(c.EndValue) {
		errs = append(errs, fmt.Errorf("endVal %v not in range %s", c.EndValue, r.String()))
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
