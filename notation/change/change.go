package change

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/validation"
)

type Change struct {
	Offset   rat.Rat `json:"offset"`
	EndValue float64 `json:"endValue"`
	Duration rat.Rat `json:"duration"`
}

func New(offset rat.Rat, endVal float64, dur rat.Rat) *Change {
	return &Change{
		Offset:   offset,
		EndValue: endVal,
		Duration: dur,
	}
}

func NewImmediate(offset rat.Rat, endVal float64) *Change {
	return New(offset, endVal, rat.Zero())
}

func (c *Change) Equal(other *Change) bool {
	return c.Offset.Equal(other.Offset) && c.EndValue == other.EndValue && c.Duration.Equal(other.Duration)
}

func (c *Change) Validate(r ValueRange) *validation.Result {
	errs := []error{}

	if c.Duration.Negative() {
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
