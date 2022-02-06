package function_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/performance/function"
)

type echo struct {
	domain function.Range
}

func (f *echo) At(x *big.Rat) float64 {
	xFlt, _ := x.Float64()

	return xFlt
}

func (f *echo) Domain() function.Range {
	return f.domain
}

func TestAtUnlimitedDomain(t *testing.T) {
	f := &echo{domain: function.DomainAll()}

	testFunctionAt(t, f, big.NewRat(0, 1), 0.0)
	testFunctionAt(t, f, big.NewRat(-1, 5), -0.2)
	testFunctionAt(t, f, big.NewRat(75, 1000), 0.075)
}

func TestAtLimitedDomain(t *testing.T) {
	min := new(big.Rat).SetInt64(-2)
	max := new(big.Rat).SetInt64(2)
	f := &echo{domain: function.NewRange(min, max)}

	testFunctionAt(t, f, big.NewRat(0, 1), 0.0)
	testFunctionAt(t, f, min, -2)
	testFunctionAt(t, f, max, 2)

	_, err := function.At(f, new(big.Rat).SetFloat64(-2.1))

	assert.NotNil(t, err)

	_, err = function.At(f, new(big.Rat).SetFloat64(2.1))

	assert.NotNil(t, err)
}

func testFunctionAt(t *testing.T, f function.Function, x *big.Rat, yExpected float64) {
	y, err := function.At(f, x)

	assert.Nil(t, err)
	assert.InDelta(t, yExpected, y, 1.0e-5)
}
