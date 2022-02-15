package change_test

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestSortChanges(t *testing.T) {
	o1 := rat.New(1, 3)
	o2 := rat.New(2, 3)
	o3 := rat.New(3, 3)
	changes := change.Changes{
		change.NewImmediate(o1, 1.1),
		change.NewImmediate(o3, 3.3),
		change.NewImmediate(o2, 2.2),
	}
	r := &change.MinExclRange{Min: 0.0}

	assert.Nil(t, changes.Validate(r))

	r.Min = 2.0

	assert.NotNil(t, changes.Validate(r))

	sort.Sort(changes)

	assert.True(t, changes[0].Offset.Equal(o1))
	assert.True(t, changes[1].Offset.Equal(o2))
	assert.True(t, changes[2].Offset.Equal(o3))
}

func TestChangeMapUnmarshalJSON(t *testing.T) {
	jsonStr := `[
		{
			"offset": "0",
			"endValue": 0.5,
			"duration": "2/1"
		}
	]`

	var cm change.Changes

	err := json.Unmarshal([]byte(jsonStr), &cm)

	require.NoError(t, err)
}
