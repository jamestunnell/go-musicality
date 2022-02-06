package function

import (
	"fmt"
	"math"
	"math/big"
)

func Sample(f Function, xrange Range, xstep *big.Rat) ([]float64, error) {
	if !xrange.IsValid() {
		return []float64{}, fmt.Errorf("x-range %v is not valid", xrange)
	}

	if xstep.Cmp(big.NewRat(0, 1)) <= 0 {
		return []float64{}, fmt.Errorf("x-step %v is not positive", xstep)
	}

	d := f.Domain()

	if !d.IncludesRange(xrange) {
		err := fmt.Errorf("x-range %v is not included in domain %v", xrange, d)
		return []float64{}, err
	}

	z, _ := new(big.Rat).Quo(xrange.Span(), xstep).Float64()
	n := 1 + int(math.Floor(z))
	samples := make([]float64, n)

	for i := 0; i < n; i++ {
		offset := new(big.Rat).Mul(big.NewRat(int64(i), 1), xstep)
		x := new(big.Rat).Add(xrange.Start, offset)
		samples[i] = f.At(x)
	}

	return samples, nil
}
