package meter_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestConvenience(t *testing.T) {
	assert.Nil(t, meter.ThreeFour().Validate())
	assert.Nil(t, meter.TwoFour().Validate())
	assert.Nil(t, meter.FourFour().Validate())
	assert.Nil(t, meter.SixEight().Validate())
	assert.Nil(t, meter.TwoTwo().Validate())
}

func TestInvalid(t *testing.T) {
	m := meter.New(0, rat.New(1, 4))
	results := m.Validate()

	require.NotNil(t, results)
	assert.Len(t, results.Errors, 1)

	m = meter.New(4, rat.Zero())
	results = m.Validate()

	require.NotNil(t, results)
	assert.Len(t, results.Errors, 1)

	m = meter.New(0, rat.Zero())
	results = m.Validate()

	require.NotNil(t, results)
	assert.Len(t, results.Errors, 2)
}

func TestMarshalUnmarshal(t *testing.T) {
	m := meter.New(4, rat.New(1, 4))

	d, err := json.Marshal(m)

	require.NoError(t, err)

	var m2 meter.Meter

	err = json.Unmarshal(d, &m2)

	require.NoError(t, err)

	assert.True(t, m2.Equal(m))
}
