package change_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestNewImmediate(t *testing.T) {
	c := change.NewImmediate(rat.Zero(), 2.5)

	assert.True(t, c.Duration.Zero())
	assert.True(t, c.Offset.Zero())
}

func TestChangeEqual(t *testing.T) {
	const testValue = 2.2
	c1 := change.NewImmediate(rat.Zero(), testValue)
	c2 := change.NewImmediate(rat.Zero(), testValue+0.1)
	c3 := change.NewImmediate(rat.New(1, 100), testValue)
	c4 := change.New(rat.Zero(), testValue, rat.New(1, 2))

	assert.True(t, c1.Equal(c1))
	assert.False(t, c1.Equal(c2))
	assert.False(t, c1.Equal(c3))
	assert.False(t, c1.Equal(c4))
}

func TestValidateDuration(t *testing.T) {
	c := change.New(rat.Zero(), 0.0, rat.FromInt64(0))
	r := &change.MinMaxInclRange{Min: 0.0, Max: 1.0}

	assert.Nil(t, c.Validate(r))

	c.Duration = rat.FromInt64(-1)

	assert.NotNil(t, c.Validate(r))
}

func TestValidateEndValue(t *testing.T) {
	c := change.New(rat.Zero(), 0.0, rat.FromInt64(1))
	r1 := &change.MinMaxInclRange{Min: 0.0, Max: 1.0}
	r2 := &change.MinExclRange{Min: 0.0}

	assert.Nil(t, c.Validate(r1))
	assert.NotNil(t, c.Validate(r2))

	c.EndValue = 1.0

	assert.Nil(t, c.Validate(r1))
	assert.Nil(t, c.Validate(r2))

	c.EndValue = -0.01

	assert.NotNil(t, c.Validate(r1))
	assert.NotNil(t, c.Validate(r2))

	c.EndValue = 1.01

	assert.NotNil(t, c.Validate(r1))
	assert.Nil(t, c.Validate(r2))
}
