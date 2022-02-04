package function_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/stretchr/testify/assert"
)

type echo struct {
	domain function.Range
}

func (f *echo) At(x float64) float64 {
	return x
}

func (f *echo) Domain() function.Range {
	return f.domain
}

func TestAtUnlimitedDomain(t *testing.T) {
	f := &echo{domain: function.DomainAllFloat64}

	testFunctionAt(t, f, 0.0, 0.0)
	testFunctionAt(t, f, -0.22, -0.22)
	testFunctionAt(t, f, 50000.2, 50000.2)
	testFunctionAt(t, f, 107.5e22, 107.5e22)
}

func TestAtLimitedDomain(t *testing.T) {
	f := &echo{domain: function.NewRange(-2.5, 2.5)}

	testFunctionAt(t, f, 0.0, 0.0)
	testFunctionAt(t, f, -2.5, -2.5)
	testFunctionAt(t, f, 2.5, 2.5)

	_, err := function.At(f, -2.51)

	assert.NotNil(t, err)

	_, err = function.At(f, 2.51)

	assert.NotNil(t, err)
}

func testFunctionAt(t *testing.T, f function.Function, x, yExpected float64) {
	y, err := function.At(f, x)

	assert.Nil(t, err)
	assert.InDelta(t, yExpected, y, 1.0e-5)
}
