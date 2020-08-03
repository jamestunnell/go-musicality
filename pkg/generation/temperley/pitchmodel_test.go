package temperley_test

import (
	"encoding/json"
	"testing"

	"github.com/jamestunnell/go-musicality/pkg/generation/temperley"
	"github.com/stretchr/testify/assert"
)

func TestPitchModel(t *testing.T) {
	pm, err := temperley.NewMajorPitchModel(0)
	assert.Nil(t, err)
	assert.NotNil(t, pm)

	pitches := pm.MakePitches(16)

	d, err := json.Marshal(pitches)
	t.Log(string(d))
}
