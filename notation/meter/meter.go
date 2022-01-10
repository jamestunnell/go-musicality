package meter

import (
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

func (m *Meter) Validate() *validation.Result {
	errs := []error{}

	if m.Numerator == 0 {
		errs = append(errs, validation.NewErrNonPositiveUInt("Numerator", 0))
	}

	if m.Denominator == 0 {
		errs = append(errs, validation.NewErrNonPositiveUInt("Denominator", 0))
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
