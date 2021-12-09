package duration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/duration"
)

func TestZero(t *testing.T) {
	d := duration.Zero()

	assert.False(t, d.Positive())
	assert.True(t, d.Zero())
}

func TestNegative(t *testing.T) {
	d := duration.New(-1, 2)

	assert.False(t, d.Positive())
	assert.False(t, d.Zero())
}

func TestPositive(t *testing.T) {
	testPositive(t, duration.New(1, 2))
	testPositive(t, duration.New(3, 2))
	testPositive(t, duration.New(100, 1))
}

func testPositive(t *testing.T, d *duration.Duration) {
	t.Run(d.String(), func(t *testing.T) {
		assert.True(t, d.Positive())
		assert.False(t, d.Zero())
	})
}
