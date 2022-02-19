package pitchgen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/composition/pitchgen"
)

func TestTemperleyGenerator(t *testing.T) {
	g, err := pitchgen.NewMajorTemperleyGenerator(0, 0)
	assert.Nil(t, err)
	assert.NotNil(t, g)

	pitches := pitchgen.MakePitches(16, g)

	for _, p := range pitches {
		assert.NotNil(t, p)
	}
}
