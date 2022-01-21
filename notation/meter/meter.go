package meter

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/validation"
)

type Meter struct {
	Numerator, Denominator uint
}

func New(num, denom uint) *Meter {
	return &Meter{
		Numerator:   num,
		Denominator: denom,
	}
}

func (m *Meter) String() string {
	return fmt.Sprintf("%d/%d", m.Numerator, m.Denominator)
}

func (m *Meter) Equal(other *Meter) bool {
	return m.Numerator == other.Numerator && m.Denominator == other.Denominator
}

func (m *Meter) Validate() *validation.Result {
	errs := []error{}

	if err := validation.VerifyNonZeroUInt("numerator", m.Numerator); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyNonZeroUInt("denominator", m.Denominator); err != nil {
		errs = append(errs, err)
	}

	if len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "meter",
		Errors:     errs,
		SubResults: []*validation.Result{},
	}
}
