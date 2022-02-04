package function_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/performance/function"
)

func TestRangeIsValid(t *testing.T) {
	assert.False(t, function.NewRange(2, 2).IsValid())
	assert.False(t, function.NewRange(3, 2).IsValid())
	assert.True(t, function.NewRange(2, 3).IsValid())
}

func TestRangeSpan(t *testing.T) {
	assert.Equal(t, 1.0, function.NewRange(2, 3).Span())
	assert.Equal(t, 4.0, function.NewRange(-2, 2).Span())
}

func TestRangeIncludesValue(t *testing.T) {
	r := function.NewRange(0.0, 2.2)

	assert.False(t, r.IncludesValue(-0.01))
	assert.True(t, r.IncludesValue(r.Start))
	assert.True(t, r.IncludesValue(r.End))
	assert.True(t, r.IncludesValue(0.1))
	assert.True(t, r.IncludesValue(0.5))
	assert.True(t, r.IncludesValue(2.1))
	assert.False(t, r.IncludesValue(2.201))
}

func TestRangeIncludesRange(t *testing.T) {
	r := function.NewRange(0.0, 2.2)

	assert.True(t, r.IncludesRange(r))
	assert.True(t, r.IncludesRange(function.NewRange(r.Start, 0.5)))
	assert.True(t, r.IncludesRange(function.NewRange(0.5, r.End)))
	assert.True(t, r.IncludesRange(function.NewRange(0.5, 1.0)))

	assert.False(t, r.IncludesRange(function.NewRange(r.Start-0.01, 1.0)))
	assert.False(t, r.IncludesRange(function.NewRange(0.5, r.End+0.01)))
	assert.False(t, r.IncludesRange(function.NewRange(-2.0, -1.0)))
	assert.False(t, r.IncludesRange(function.NewRange(3.0, 4.0)))
}
