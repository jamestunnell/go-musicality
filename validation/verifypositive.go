package validation

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
)

func VerifyPositiveRat(name string, r *big.Rat) error {
	if !rat.IsPositive(r) {
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
