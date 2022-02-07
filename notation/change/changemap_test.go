package change_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/change"
)

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
