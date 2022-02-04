package temperley_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/generation/temperley"
)

func TestPitchModel(t *testing.T) {
	pm, err := temperley.NewMajorPitchModel(0, 0)
	assert.Nil(t, err)
	assert.NotNil(t, pm)

	pitches := pm.MakePitches(16)

	d, err := json.Marshal(pitches)
	t.Log(string(d))
}
