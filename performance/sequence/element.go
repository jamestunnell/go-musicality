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

	if e.Duration.Cmp(big.NewRat(0, 1)) < 1 {
		errs = append(errs, validation.NewErrNonPositiveRat("duration", e.Duration))
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
