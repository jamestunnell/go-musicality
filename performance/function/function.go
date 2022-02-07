package function

import (
	"math"

	"github.com/jamestunnell/go-musicality/notation/rat"
)

func DomainMin() rat.Rat {
	return rat.New(-math.MaxInt64, 1)
}

func DomainMax() rat.Rat {
	return rat.New(math.MaxInt64, 1)
}

func DomainAll() Range {
	return NewRange(DomainMin(), DomainMax())
}

type Function interface {
	At(x rat.Rat) float64
	Domain() Range
}
