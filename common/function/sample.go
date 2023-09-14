package function

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
)

func Sample(f Function, xrange Range, xstep *big.Rat) ([]float64, error) {
	if !xrange.IsValid() {
		return []float64{}, fmt.Errorf("x-range %v is not valid", xrange)
	}

	if !rat.IsPositive(xstep) {
		return []float64{}, fmt.Errorf("x-step %v is not positive", xstep)
	}

	d := f.Domain()

	if !d.IncludesRange(xrange) {
		err := fmt.Errorf("x-range %v is not included in domain %v", xrange, d)
		return []float64{}, err
	}

	samples := []float64{}

	x := xrange.Start
	for rat.IsLessEqual(x, xrange.End) {
		samples = append(samples, f.At(x))

		x = rat.Add(x, xstep)
	}

	return samples, nil
}
