package rat

import (
	"math/big"
)

var zero = big.NewRat(0, 1)

func FromFloat64(x float64) *big.Rat {
	return new(big.Rat).SetFloat64(x)
}

func FromInt64(x int64) *big.Rat {
	return new(big.Rat).SetInt64(x)
}

func FromUint64(x uint64) *big.Rat {
	return new(big.Rat).SetUint64(x)
}

func Zero() *big.Rat {
	return new(big.Rat)
}

func IsPositive(r *big.Rat) bool {
	return r.Cmp(zero) == 1
}

func IsNegative(r *big.Rat) bool {
	return r.Cmp(zero) == -1
}

func IsZero(r *big.Rat) bool {
	return r.Cmp(zero) == 0
}

func IsEqual(r1, r2 *big.Rat) bool {
	return r1.Cmp(r2) == 0
}

func IsGreaterEqual(r1, r2 *big.Rat) bool {
	return r1.Cmp(r2) >= 0
}

func IsGreater(r1, r2 *big.Rat) bool {
	return r1.Cmp(r2) == 1
}

func IsLessEqual(r1, r2 *big.Rat) bool {
	return r1.Cmp(r2) <= 0
}

func IsLess(r1, r2 *big.Rat) bool {
	return r1.Cmp(r2) == -1
}

func Add(r1, r2 *big.Rat) *big.Rat {
	return new(big.Rat).Add(r1, r2)
}

func Sub(r1, r2 *big.Rat) *big.Rat {
	return new(big.Rat).Sub(r1, r2)
}

func Mul(r1, r2 *big.Rat) *big.Rat {
	return new(big.Rat).Mul(r1, r2)
}

func Div(r1, r2 *big.Rat) *big.Rat {
	return new(big.Rat).Quo(r1, r2)
}
