package function

import (
	"math/big"

	"gonum.org/v1/gonum/stat/distuv"
)

type RandomFunction struct {
	r distuv.Rander
}

func NewRandomFunction(r distuv.Rander) *RandomFunction {
	return &RandomFunction{r: r}
}

func (s *RandomFunction) Domain() Range {
	return DomainAll()
}

func (s *RandomFunction) At(x *big.Rat) float64 {
	return s.r.Rand()
}
