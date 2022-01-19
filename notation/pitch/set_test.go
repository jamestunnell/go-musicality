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

func TestSetOperations(t *testing.T) {
	s1 := pitch.NewSet(pitch.A1, pitch.A2, pitch.A3)
	s2 := pitch.NewSet(pitch.A0, pitch.A2, pitch.A3)
	union := s1.Union(s2)
	intersect := s1.Intersect(s2)
	diff := s1.Diff(s2)

	assert.Equal(t, 4, union.Len())
	assert.ElementsMatch(t, union.Pitches(), pitch.Pitches{pitch.A0, pitch.A1, pitch.A2, pitch.A3})

	assert.Equal(t, 2, intersect.Len())
	assert.ElementsMatch(t, intersect.Pitches(), pitch.Pitches{pitch.A2, pitch.A3})

	assert.Equal(t, 1, diff.Len())
	assert.ElementsMatch(t, diff.Pitches(), pitch.Pitches{pitch.A1})

	diff = s2.Diff(s1)

	assert.Equal(t, 1, diff.Len())
	assert.ElementsMatch(t, diff.Pitches(), pitch.Pitches{pitch.A0})
}
