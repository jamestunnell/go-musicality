package function

import (
	"math"
	"math/big"
)

func DomainMin() *big.Rat {
	return big.NewRat(-math.MaxInt64, 1)
}

func DomainMax() *big.Rat {
	return big.NewRat(math.MaxInt64, 1)
}

func DomainAll() Range {
	return NewRange(DomainMin(), DomainMax())
}

type Function interface {
	At(x *big.Rat) float64
	Domain() Range
}
