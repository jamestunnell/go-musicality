package sequence

import (
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/validation"
)

type Element struct {
	Duration   float64
	Pitch      *pitch.Pitch
	Attack     float64
	Separation float64
}

func (e *Element) Validate() *validation.Result {
	errs := []error{}

	if e.Duration <= 0.0 {
		errs = append(errs, validation.NewErrNonPositiveFloat("duration", e.Duration))
	}

	if len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "element",
		Errors:     errs,
		SubResults: []*validation.Result{},
	}
}
