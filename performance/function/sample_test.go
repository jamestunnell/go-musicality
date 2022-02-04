package function_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/performance/function"
)

func TestSampleUnlimitedDomain(t *testing.T) {
	f := &echo{domain: function.DomainAllFloat64}

	testFunctionSample(t, f, function.NewRange(0.0, 10.0), 2.0, []float64{0.0, 2.0, 4.0, 6.0, 8.0, 10.0})
	testFunctionSample(t, f, function.NewRange(-0.5e50, 0.5e50), 0.5e50, []float64{-0.5e50, 0.0, 0.5e50})
}

func TestSampleLimitedDomain(t *testing.T) {
	f := &echo{domain: function.NewRange(-5, 5)}

	testFunctionSample(t, f, function.NewRange(-5, 5), 5, []float64{-5, 0, 5})
	testFunctionSample(t, f, function.NewRange(-5, 0), 1, []float64{-5, -4, -3, -2, -1, 0})
	testFunctionSample(t, f, function.NewRange(0, 5), 1, []float64{0, 1, 2, 3, 4, 5})

	testFunctionSampleBadRange(t, f, function.NewRange(-10, 0), 1)
	testFunctionSampleBadRange(t, f, function.NewRange(0, 10), 1)
	testFunctionSampleBadRange(t, f, function.NewRange(-10, -8), 1)
	testFunctionSampleBadRange(t, f, function.NewRange(8, 10), 1)
}

func TestSampleNonPositiveXStep(t *testing.T) {
	f := &echo{domain: function.NewRange(-5, 5)}

	_, err := function.Sample(f, f.domain, 0.0)

	assert.NotNil(t, err)

	_, err = function.Sample(f, f.domain, -1.0)

	assert.NotNil(t, err)
}

func testFunctionSample(t *testing.T, f function.Function, xrange function.Range, xstep float64, expected []float64) {
	samples, err := function.Sample(f, xrange, xstep)

	assert.Nil(t, err)

	if !assert.Equal(t, len(expected), len(samples)) {
		return
	}

	for i, sample := range samples {
		assert.InDelta(t, expected[i], sample, 1.0e-5)
	}
}

func testFunctionSampleBadRange(t *testing.T, f function.Function, xrange function.Range, xstep float64) {
	_, err := function.Sample(f, xrange, xstep)

	assert.NotNil(t, err)
}
