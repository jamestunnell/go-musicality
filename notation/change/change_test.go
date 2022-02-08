package change_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestNewImmediate(t *testing.T) {
	c := change.NewImmediate(2.5)

	assert.True(t, c.Duration.Zero())
}

func TestValidateDuration(t *testing.T) {
	c := change.New(0.0, rat.FromInt64(0))
	r := &change.MinMaxInclRange{Min: 0.0, Max: 1.0}

	assert.Nil(t, c.Validate(r))

	c.Duration = rat.FromInt64(-1)

	assert.NotNil(t, c.Validate(r))
}

func TestValidateEndValue(t *testing.T) {
	c := change.New(0.0, rat.FromInt64(1))
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
