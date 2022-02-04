package function

import (
	"fmt"
	"math"
)

func Sample(f Function, xrange Range, xstep float64) ([]float64, error) {
	if !xrange.IsValid() {
		return []float64{}, fmt.Errorf("x-range %v is not valid", xrange)
	}

	if xstep <= 0.0 {
		return []float64{}, fmt.Errorf("x-step %f is not positive", xstep)
	}

	d := f.Domain()

	if !d.IncludesRange(xrange) {
		err := fmt.Errorf("x-range %v is not included in domain %v", xrange, d)
		return []float64{}, err
	}

	n := 1 + int(math.Floor(xrange.Span()/xstep))
	samples := make([]float64, n)

	for i := 0; i < n; i++ {
		x := xrange.Start + (float64(i) * xstep)
		samples[i] = f.At(x)
	}

	return samples, nil
}
