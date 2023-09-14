package function_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
)

var (
	negFive = rat.FromInt64(-5)
	five    = rat.FromInt64(5)
	zero    = rat.Zero()
	one     = rat.FromInt64(1)
)

func TestSampleUnlimitedDomain(t *testing.T) {
	f := &echo{domain: function.DomainAll()}
	xRange := function.NewRange(zero, big.NewRat(10, 1))
	xStep := two
	expected := []float64{0.0, 2.0, 4.0, 6.0, 8.0, 10.0}

	testFunctionSample(t, f, xRange, xStep, expected)

	xRange = function.NewRange(negHalf, half)
	xStep = half
	expected = []float64{-0.5, 0.0, 0.5}

	testFunctionSample(t, f, xRange, xStep, expected)
}

func TestSampleLimitedDomain(t *testing.T) {
	f := &echo{domain: function.NewRange(negFive, five)}

	testFunctionSample(t, f, function.NewRange(negFive, five), five, []float64{-5, 0, 5})
	testFunctionSample(t, f, function.NewRange(negFive, zero), one, []float64{-5, -4, -3, -2, -1, 0})
	testFunctionSample(t, f, function.NewRange(zero, five), one, []float64{0, 1, 2, 3, 4, 5})

	testFunctionSampleBadRange(t, f, function.NewRange(big.NewRat(-10, 1), zero), one)
	testFunctionSampleBadRange(t, f, function.NewRange(zero, big.NewRat(10, 1)), one)
	testFunctionSampleBadRange(t, f, function.NewRange(big.NewRat(-10, 1), big.NewRat(-8, 1)), one)
	testFunctionSampleBadRange(t, f, function.NewRange(big.NewRat(8, 1), big.NewRat(10, 1)), one)
}

func TestSampleNonPositiveXStep(t *testing.T) {
	f := &echo{domain: function.NewRange(negFive, five)}

	_, err := function.Sample(f, f.domain, zero)

	assert.NotNil(t, err)

	_, err = function.Sample(f, f.domain, negOne)

	assert.NotNil(t, err)
}

func testFunctionSample(
	t *testing.T,
	f function.Function,
	xrange function.Range,
	xstep *big.Rat, expected []float64) {
	samples, err := function.Sample(f, xrange, xstep)

	assert.Nil(t, err)

	if !assert.Equal(t, len(expected), len(samples)) {
		return
	}

	for i, sample := range samples {
		assert.InDelta(t, expected[i], sample, 1.0e-5)
	}
}

func testFunctionSampleBadRange(t *testing.T, f function.Function, xrange function.Range, xstep *big.Rat) {
	_, err := function.Sample(f, xrange, xstep)

	assert.NotNil(t, err)
}
