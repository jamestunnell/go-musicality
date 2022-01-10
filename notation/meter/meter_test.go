package meter_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	m := meter.New(3, 4)

	assert.Nil(t, m.Validate())
	assert.Equal(t, m.Numerator, uint(3))
	assert.Equal(t, m.Denominator, uint(4))
}

func TestInvalid(t *testing.T) {
	m := meter.New(0, 4)
	results := m.Validate()

	require.NotNil(t, results)
	assert.Len(t, results.Errors, 1)

	m = meter.New(4, 0)
	results = m.Validate()

	require.NotNil(t, results)
	assert.Len(t, results.Errors, 1)

	m = meter.New(0, 0)
	results = m.Validate()

	require.NotNil(t, results)
	assert.Len(t, results.Errors, 2)
}
