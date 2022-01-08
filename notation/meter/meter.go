package meter

import (
	"github.com/jamestunnell/go-musicality/validation"
)

type Meter struct {
	numerator, denominator uint
}

func New(num, denom uint) *Meter {
	return &Meter{
		numerator:   num,
		denominator: denom,
	}
}

func (m *Meter) Numerator() uint {
	return m.numerator
}

func (m *Meter) Denominator() uint {
	return m.denominator
}

func (m *Meter) Validate() *validation.Result {
	errs := []error{}

	if m.numerator == 0 {
		errs = append(errs, validation.NewErrNonPositiveUInt("numerator", 0))
	}

	if m.denominator == 0 {
		errs = append(errs, validation.NewErrNonPositiveUInt("denominator", 0))
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
