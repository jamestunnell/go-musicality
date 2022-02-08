package meter_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/meter"
)

func TestNew(t *testing.T) {
	m := meter.New(3, 4)

	assert.Nil(t, m.Validate())
	assert.Equal(t, uint64(3), m.Numerator)
	assert.Equal(t, uint64(4), m.Denominator)
}

func TestString(t *testing.T) {
	assert.Equal(t, "3/4", meter.New(3, 4).String())
	assert.Equal(t, "4/4", meter.New(4, 4).String())
	assert.Equal(t, "6/8", meter.New(6, 8).String())
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

func TestMarshalUnmarshal(t *testing.T) {
	m := meter.New(4, 4)

	d, err := json.Marshal(m)

	require.NoError(t, err)

	var m2 meter.Meter

	err = json.Unmarshal(d, &m2)

	require.NoError(t, err)

	assert.True(t, m2.Equal(m))
}

func TestUnmarshalWrongType(t *testing.T) {
	var m2 meter.Meter

	err := json.Unmarshal([]byte(`"-4/4"`), &m2)

	assert.Error(t, err)
}
