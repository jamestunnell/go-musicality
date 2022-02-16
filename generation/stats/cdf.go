package stats

import (
	"errors"
	"log"
	"math"
	"math/rand"
)

type CDF struct {
	values []float64
}

const UnityTolerance = 1e-5

// Convert probability density/mass function values, which should have already been
// normalized to sum to 1, into cumulative distribution function values.
func NewCDF(probs []float64) (*CDF, error) {
	n := len(probs)
	if n == 0 {
		return nil, errors.New("no probability values given")
	}

	cdfValues := make([]float64, n)

	cdfValues[0] = probs[0]

	for i := 1; i < n; i++ {
		cdfValues[i] = cdfValues[i-1] + probs[i]
	}

	if math.Abs(cdfValues[n-1]-1.0) > UnityTolerance {
		return nil, errors.New(
			"last CDF value not close to 1, input probabilities may not have been normalized")
	}

	cdfValues[n-1] = 1.0

	return &CDF{values: cdfValues}, nil
}

func (cdf *CDF) Size() int {
	return len(cdf.values)
}

// Use rand.Float64 to randomly select one index based on the CDF values.
func (cdf *CDF) Rand() int {
	x := rand.Float64()

	for i := 0; i < len(cdf.values); i++ {
		if x < cdf.values[i] {
			return i
		}
	}

	log.Fatal("failed to select a pitch semitone")

	return -1
}
