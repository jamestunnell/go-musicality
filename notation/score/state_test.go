package score_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/stretchr/testify/assert"
)

func TestStateValid(t *testing.T) {
	s := &score.State{
		Tempo:   120.0,
		Dynamic: 0.0,
	}

	assert.Nil(t, s.Validate())
}

func TestStateInvalid(t *testing.T) {
	testStateInvalid(t, "zero tempo", 0.0, 0.5)
	testStateInvalid(t, "neg. tempo", -1.0, 0.5)
	testStateInvalid(t, "volume < -1", 120.0, -1.1)
	testStateInvalid(t, "volume > 1", 120.0, 1.01)
}

func testStateInvalid(t *testing.T, name string, tempo, vol float64) {
	t.Run(name, func(t *testing.T) {
		s := &score.State{Tempo: tempo, Dynamic: vol}

		assert.NotNil(t, s.Validate())
	})
}
