package change_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestChangeMap(t *testing.T) {
	o1 := rat.New(1, 3)
	o2 := rat.New(2, 3)
	o3 := rat.New(3, 3)
	cm := change.Map{
		rat.New(3, 3): change.NewImmediate(1.1),
		rat.New(1, 3): change.NewImmediate(3.3),
		rat.New(2, 3): change.NewImmediate(2.2),
	}
	r := &change.MinExclRange{Min: 0.0}

	assert.Nil(t, cm.Validate(r))

	r.Min = 2.0

	assert.NotNil(t, cm.Validate(r))

	offsets := cm.SortedOffsets()

	require.Len(t, offsets, 3)

	assert.True(t, offsets[0].Equal(o1))
	assert.True(t, offsets[1].Equal(o2))
	assert.True(t, offsets[2].Equal(o3))
}

func TestChangeMapUnmarshalJSON(t *testing.T) {
	jsonStr := `{
		"0": {
			"endValue": 0.5,
			"duration": "2/1"
		}
	}`

	var cm change.Map

	err := json.Unmarshal([]byte(jsonStr), &cm)

	require.NoError(t, err)
}
