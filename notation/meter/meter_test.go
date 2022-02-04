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
	assert.Equal(t, uint64(3), m.Numerator)
	assert.Equal(t, uint64(4), m.Denominator)
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
