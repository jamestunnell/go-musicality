package validation

import "github.com/jamestunnell/go-musicality/notation/rat"

func VerifyPositiveRat(name string, r rat.Rat) error {
	if !r.Positive() {
		return NewErrNonPositive(name, r.String())
	}

	return nil
}

func VerifyPositiveFloat(name string, f float64) error {
	if f <= 0.0 {
		return NewErrNonPositive(name, f)
	}

	return nil
}
