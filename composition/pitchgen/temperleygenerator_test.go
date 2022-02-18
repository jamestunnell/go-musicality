package pitchgen_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/composition/pitchgen"
)

func TestPitchModel(t *testing.T) {
	g, err := pitchgen.NewMajorTemperleyGenerator(0, 0)
	assert.Nil(t, err)
	assert.NotNil(t, g)

	pitches := g.MakePitches(16)

	d, err := json.Marshal(pitches)
	t.Log(string(d))
}
