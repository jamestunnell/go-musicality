package function_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/performance/function"
)

var (
	negTwo = big.NewRat(-2, 1)
	negOne = big.NewRat(-1, 1)
	half   = new(big.Rat).SetFloat64(0.5)
	two    = big.NewRat(2, 1)
	three  = big.NewRat(3, 1)
	four   = big.NewRat(4, 1)
)

func TestRangeIsValid(t *testing.T) {
	assert.False(t, function.NewRange(two, two).IsValid())
	assert.False(t, function.NewRange(three, two).IsValid())
	assert.True(t, function.NewRange(two, three).IsValid())
}

func TestRangeSpan(t *testing.T) {
	assert.Equal(t, one, function.NewRange(two, three).Span())
	assert.Equal(t, four, function.NewRange(negTwo, two).Span())
}

func TestRangeIncludesValue(t *testing.T) {
	r := function.NewRange(zero, two)

	assert.False(t, r.IncludesValue(new(big.Rat).SetFloat64(-.01)))
	assert.True(t, r.IncludesValue(r.Start))
	assert.True(t, r.IncludesValue(r.End))
	assert.True(t, r.IncludesValue(new(big.Rat).SetFloat64(0.1)))
	assert.True(t, r.IncludesValue(half))
	assert.False(t, r.IncludesValue(new(big.Rat).SetFloat64(2.01)))
}

func TestRangeIncludesRange(t *testing.T) {
	r := function.NewRange(zero, two)

	assert.True(t, r.IncludesRange(r))
	assert.True(t, r.IncludesRange(function.NewRange(r.Start, half)))
	assert.True(t, r.IncludesRange(function.NewRange(half, r.End)))
	assert.True(t, r.IncludesRange(function.NewRange(half, one)))

	assert.False(t, r.IncludesRange(function.NewRange(new(big.Rat).SetFloat64(-.01), one)))
	assert.False(t, r.IncludesRange(function.NewRange(half, new(big.Rat).SetFloat64(2.01))))
	assert.False(t, r.IncludesRange(function.NewRange(negTwo, negOne)))
	assert.False(t, r.IncludesRange(function.NewRange(three, four)))
}
