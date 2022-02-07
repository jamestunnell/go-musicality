package function

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-musicality/notation/rat"
)

func Sample(f Function, xrange Range, xstep rat.Rat) ([]float64, error) {
	if !xrange.IsValid() {
		return []float64{}, fmt.Errorf("x-range %v is not valid", xrange)
	}

	if !xstep.Positive() {
		return []float64{}, fmt.Errorf("x-step %v is not positive", xstep)
	}

	d := f.Domain()

	if !d.IncludesRange(xrange) {
		err := fmt.Errorf("x-range %v is not included in domain %v", xrange, d)
		return []float64{}, err
	}

	n := 1 + int(math.Floor(xrange.Span().Div(xstep).Float64()))
	samples := make([]float64, n)

	x := xrange.Start.Clone()
	for i := 0; i < n; i++ {
		samples[i] = f.At(x)

		x.Accum(xstep)
	}

	return samples, nil
}
