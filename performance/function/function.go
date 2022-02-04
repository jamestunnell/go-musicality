package function

import (
	"math"
)

var (
	DomainAllFloat64 = NewRange(-math.MaxFloat64, math.MaxFloat64)
)

type Function interface {
	At(x float64) float64
	Domain() Range
}
