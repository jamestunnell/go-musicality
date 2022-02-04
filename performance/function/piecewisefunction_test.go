package function_test

import (
	"math"
	"testing"

	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/stretchr/testify/assert"
)

func TestPiecewiseFunctionLimitedDomain(t *testing.T) {
	pairs := []function.SubdomainFunctionPair{
		{Subdomain: function.NewRange(-2, 0), Function: function.NewConstantFunction(-1)},
		{Subdomain: function.NewRange(0, 2), Function: function.NewConstantFunction(1)},
	}
	f, err := function.NewPiecewiseFunction(pairs)

	assert.Nil(t, err)
	assert.NotNil(t, f)

	testFunctionAt(t, f, -2, -1)
	testFunctionAt(t, f, -0.01, -1)
	testFunctionAt(t, f, 0, 1)
	testFunctionAt(t, f, 0.01, 1)
	testFunctionAt(t, f, 2, 1)

	_, err = function.At(f, -2.01)

	assert.NotNil(t, err)

	_, err = function.At(f, 2.01)

	assert.NotNil(t, err)
}

func TestPiecewiseFunctionUnlimitedDomain(t *testing.T) {
	fA := function.NewLinearFunction(7, 5)
	fB := function.NewLinearFunction(-2.5, -10)
	xBoundary := -5.6789
	pairs := []function.SubdomainFunctionPair{
		{Subdomain: function.NewRange(-math.MaxFloat64, xBoundary), Function: fA},
		{Subdomain: function.NewRange(xBoundary, math.MaxFloat64), Function: fB},
	}
	f, err := function.NewPiecewiseFunction(pairs)

	assert.Nil(t, err)
	assert.NotNil(t, f)

	testFunctionAt(t, f, -2.4e55, fA.At(-2.4e55))
	testFunctionAt(t, f, xBoundary-0.01, fA.At(xBoundary-0.01))
	testFunctionAt(t, f, xBoundary, fB.At(xBoundary))
	testFunctionAt(t, f, xBoundary+0.01, fB.At(xBoundary+0.01))
	testFunctionAt(t, f, 66.3e34, fB.At(66.3e34))
}
