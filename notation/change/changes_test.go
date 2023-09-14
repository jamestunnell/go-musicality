package change_test

import (
	"encoding/json"
	"math/big"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/change"
)

func TestSortChanges(t *testing.T) {
	o1 := big.NewRat(1, 3)
	o2 := big.NewRat(2, 3)
	o3 := big.NewRat(3, 3)
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

	assert.True(t, rat.IsEqual(changes[0].Offset, o1))
	assert.True(t, rat.IsEqual(changes[1].Offset, o2))
	assert.True(t, rat.IsEqual(changes[2].Offset, o3))
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
