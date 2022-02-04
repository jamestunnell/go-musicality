package function_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/performance/function"
)

var (
	pZeroZero = function.NewPoint(0.0, 0.0)
	pOneOne   = function.NewPoint(1.0, 1.0)
)

func TestSigmoid(t *testing.T) {
	f := function.NewSigmoidFunction(pZeroZero, pOneOne)

	testFunctionAt(t, f, 0.0, 0.0)
	testFunctionAt(t, f, 0.5, 0.5)
	testFunctionAt(t, f, 1.0, 1.0)

	testFunctionSample(t, f, function.NewRange(0.0, 1.0), 0.5, []float64{0.0, 0.5, 1.0})
}

func TestSigmoidOutOfDomain(t *testing.T) {
	f := function.NewSigmoidFunction(pZeroZero, pOneOne)

	_, err := function.At(f, -0.01)

	assert.NotNil(t, err)

	_, err = function.At(f, 1.01)

	assert.NotNil(t, err)
}
