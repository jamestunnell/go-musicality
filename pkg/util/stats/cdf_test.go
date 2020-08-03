package stats_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/pkg/util/stats"
	"github.com/stretchr/testify/assert"
)

func TestNewCDFNoValues(t *testing.T) {
	cdf, err := stats.NewCDF([]float64{})

	assert.Nil(t, cdf)
	assert.NotNil(t, err)
}

func TestNewCDFValuesDontSumToOne(t *testing.T) {
	testCases := [][]float64{
		[]float64{0.2, 0.8 + stats.UnityTolerance + 1e-5},
		[]float64{0.2, 0.81},
		[]float64{0.2, 0.79},
		[]float64{0.2, 0.8 - (stats.UnityTolerance + 1e-5)},
	}

	for _, vals := range testCases {
		cdf, err := stats.NewCDF(vals)

		assert.Nil(t, cdf)
		assert.NotNil(t, err)
	}
}

func TestNewCDFValuesSumVeryCloseToOne(t *testing.T) {
	testCases := [][]float64{
		[]float64{0.5, 0.5 + stats.UnityTolerance*0.99},
		[]float64{0.5, 0.5 + stats.UnityTolerance/2},
		[]float64{0.5, 0.5 - stats.UnityTolerance/2},
		[]float64{0.5, 0.5 - stats.UnityTolerance*0.99},
	}

	for _, vals := range testCases {
		cdf, err := stats.NewCDF(vals)

		assert.NotNil(t, cdf)
		assert.Nil(t, err)
	}
}

func TestCDFSize(t *testing.T) {
	probs := []float64{0.2, 0.2, 0.25, 0.05, 0.3}
	cdf, err := stats.NewCDF(probs)

	assert.NotNil(t, cdf)
	assert.Nil(t, err)

	assert.Equal(t, len(probs), cdf.Size())
}

func TestCDFRand1(t *testing.T) {
	probs := []float64{1.0}
	cdf, err := stats.NewCDF(probs)

	assert.NotNil(t, cdf)
	assert.Nil(t, err)

	for i := 0; i < 100; i++ {
		assert.Equal(t, 0, cdf.Rand())
	}
}

func TestCDFRand5050(t *testing.T) {
	probs := []float64{0.5, 0.5}
	cdf, err := stats.NewCDF(probs)

	assert.NotNil(t, cdf)
	assert.Nil(t, err)

	for i := 0; i < 100; i++ {
		n := cdf.Rand()
		assert.True(t, n == 0 || n == 1)
	}
}

func TestCDFRandCounts(t *testing.T) {
	probs := []float64{0.1, 0.2, 0.3, 0.4}
	cdf, err := stats.NewCDF(probs)

	assert.NotNil(t, cdf)
	assert.Nil(t, err)

	counts := map[int]uint{}

	for i := 0; i < 1000; i++ {
		n := cdf.Rand()

		assert.True(t, n >= 0 && n <= 3)

		if count, found := counts[n]; found {
			counts[n] = count + 1
		} else {
			counts[n] = 1
		}
	}

	for i := 1; i < cdf.Size(); i++ {
		assert.GreaterOrEqual(t, counts[i], counts[i-1])
	}
}
