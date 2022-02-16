package function_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/function"
)

var (
	negTwo = rat.FromInt64(-2)
	negOne = rat.FromInt64(-1)
	half   = rat.FromFloat64(0.5)
	two    = rat.FromInt64(2)
	three  = rat.FromInt64(3)
	four   = rat.FromInt64(4)
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

	assert.False(t, r.IncludesValue(rat.FromFloat64(-.01)))
	assert.True(t, r.IncludesValue(r.Start))
	assert.True(t, r.IncludesValue(r.End))
	assert.True(t, r.IncludesValue(rat.FromFloat64(0.1)))
	assert.True(t, r.IncludesValue(half))
	assert.False(t, r.IncludesValue(rat.FromFloat64(2.01)))
}

func TestRangeIncludesRange(t *testing.T) {
	r := function.NewRange(zero, two)

	assert.True(t, r.IncludesRange(r))
	assert.True(t, r.IncludesRange(function.NewRange(r.Start, half)))
	assert.True(t, r.IncludesRange(function.NewRange(half, r.End)))
	assert.True(t, r.IncludesRange(function.NewRange(half, one)))

	assert.False(t, r.IncludesRange(function.NewRange(rat.FromFloat64(-.01), one)))
	assert.False(t, r.IncludesRange(function.NewRange(half, rat.FromFloat64(2.01))))
	assert.False(t, r.IncludesRange(function.NewRange(negTwo, negOne)))
	assert.False(t, r.IncludesRange(function.NewRange(three, four)))
}
