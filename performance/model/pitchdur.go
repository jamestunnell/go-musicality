package sequence

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/validation"
)

type Element struct {
	Duration *big.Rat
	Pitch    *pitch.Pitch
	Attack   float64
}

func (e *Element) Validate() *validation.Result {
	errs := []error{}

	if err := validation.VerifyPositiveRat("duration", e.Duration); err != nil {
		errs = append(errs, err)
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
