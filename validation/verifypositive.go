package validation

import "math/big"

func VerifyPositiveRat(name string, r *big.Rat) error {
	if r.Cmp(big.NewRat(0, 1)) != 1 {
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
