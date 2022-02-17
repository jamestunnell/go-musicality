package function_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/function"
	"github.com/jamestunnell/go-musicality/common/rat"
)

type echo struct {
	domain function.Range
}

func (f *echo) At(x rat.Rat) float64 {
	return x.Float64()
}

func (f *echo) Domain() function.Range {
	return f.domain
}

func TestAtUnlimitedDomain(t *testing.T) {
	f := &echo{domain: function.DomainAll()}

	testFunctionAt(t, f, zero, 0.0)
	testFunctionAt(t, f, rat.New(-1, 5), -0.2)
	testFunctionAt(t, f, rat.New(75, 1000), 0.075)
}

func TestAtLimitedDomain(t *testing.T) {
	min := rat.FromInt64(-2)
	max := rat.FromInt64(2)
	f := &echo{domain: function.NewRange(min, max)}

	testFunctionAt(t, f, zero, 0.0)
	testFunctionAt(t, f, min, -2)
	testFunctionAt(t, f, max, 2)

	_, err := function.At(f, rat.FromFloat64(-2.1))

	assert.NotNil(t, err)

	_, err = function.At(f, rat.FromFloat64(2.1))

	assert.NotNil(t, err)
}

func testFunctionAt(t *testing.T, f function.Function, x rat.Rat, yExpected float64) {
	y, err := function.At(f, x)

	assert.Nil(t, err)
	assert.InDelta(t, yExpected, y, 1.0e-5)
}
