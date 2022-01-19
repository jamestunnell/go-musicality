package pitch_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/stretchr/testify/assert"
)

func TestSetEmpty(t *testing.T) {
	s := pitch.NewSet()

	assert.False(t, s.Contains(pitch.C0))
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Remove(pitch.C0))
}

func TestSetUnion(t *testing.T) {
	s1 := pitch.NewSet(pitch.A1)
	s2 := pitch.NewSet(pitch.A0, pitch.A2)
	u := s1.Union(s2)

	assert.Len(t, u, 3)
	assert.ElementsMatch(t, u, []*pitch.Pitch{pitch.A0, pitch.A1, pitch.A2})
}
