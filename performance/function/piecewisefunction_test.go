package function_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/performance/function"
)

var (
	negHundredth = new(big.Rat).SetFloat64(-0.01)
	hundredth    = new(big.Rat).SetFloat64(0.01)
)

func TestPiecewiseFunctionLimitedDomain(t *testing.T) {
	pairs := []function.SubdomainFunctionPair{
		{Subdomain: function.NewRange(negTwo, zero), Function: function.NewConstantFunction(-1)},
		{Subdomain: function.NewRange(zero, two), Function: function.NewConstantFunction(1)},
	}
	f, err := function.NewPiecewiseFunction(pairs)

	assert.Nil(t, err)
	assert.NotNil(t, f)

	testFunctionAt(t, f, negTwo, -1)
	testFunctionAt(t, f, negHundredth, -1)
	testFunctionAt(t, f, zero, 1)
	testFunctionAt(t, f, hundredth, 1)
	testFunctionAt(t, f, two, 1)

	_, err = function.At(f, new(big.Rat).SetFloat64(-2.01))

	assert.NotNil(t, err)

	_, err = function.At(f, new(big.Rat).SetFloat64(2.01))

	assert.NotNil(t, err)
}

func TestPiecewiseFunctionUnlimitedDomain(t *testing.T) {
	fA := function.NewLinearFunction(7, 5)
	fB := function.NewLinearFunction(-2.5, -10)
	xBoundary := new(big.Rat).SetFloat64(-5.6789)
	pairs := []function.SubdomainFunctionPair{
		{Subdomain: function.NewRange(function.DomainMin(), xBoundary), Function: fA},
		{Subdomain: function.NewRange(xBoundary, function.DomainMax()), Function: fB},
	}
	f, err := function.NewPiecewiseFunction(pairs)

	assert.Nil(t, err)
	assert.NotNil(t, f)

	beforeXBoundary1 := new(big.Rat).Sub(xBoundary, two)
	beforeXBoundary2 := new(big.Rat).Sub(xBoundary, hundredth)
	afterXBoundary1 := new(big.Rat).Add(xBoundary, hundredth)
	afterXBoundary2 := new(big.Rat).Add(xBoundary, two)

	testFunctionAt(t, f, beforeXBoundary1, fA.At(beforeXBoundary1))
	testFunctionAt(t, f, beforeXBoundary2, fA.At(beforeXBoundary2))
	testFunctionAt(t, f, xBoundary, fB.At(xBoundary))
	testFunctionAt(t, f, afterXBoundary1, fB.At(afterXBoundary1))
	testFunctionAt(t, f, afterXBoundary2, fB.At(afterXBoundary2))
}
