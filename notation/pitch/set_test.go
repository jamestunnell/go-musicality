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
	assert.Empty(t, s.Pitches())
	assert.False(t, s.Remove(pitch.C0))
}

func TestSetNotEmpty(t *testing.T) {
	s := pitch.NewSet(pitch.D3, pitch.F3)

	assert.True(t, s.Contains(pitch.D3))
	assert.True(t, s.Contains(pitch.F3))
	assert.Equal(t, 2, s.Len())
	assert.Len(t, s.Pitches(), 2)
	assert.True(t, s.Remove(pitch.D3))
	assert.False(t, s.Remove(pitch.D3))
	assert.False(t, s.Contains(pitch.D3))
}

func TestSetEqual(t *testing.T) {
	s1 := pitch.NewSet()
	s2 := pitch.NewSet(pitch.D3, pitch.F3)
	s3 := pitch.NewSet(pitch.A3, pitch.F3)

	assert.True(t, s1.Equal(s1))
	assert.False(t, s1.Equal(s2))
	assert.False(t, s1.Equal(s3))

	assert.True(t, s2.Equal(s2))
	assert.False(t, s2.Equal(s3))
}

func TestSetOperations(t *testing.T) {
	s1 := pitch.NewSet(pitch.A1, pitch.A2, pitch.A3)
	s2 := pitch.NewSet(pitch.A0, pitch.A2, pitch.A3)
	union := s1.Union(s2)
	unionExpected := pitch.NewSet(pitch.A0, pitch.A1, pitch.A2, pitch.A3)
	intersect := s1.Intersect(s2)
	intersectExpected := pitch.NewSet(pitch.A2, pitch.A3)
	diff := s1.Diff(s2)
	diffExpected := pitch.NewSet(pitch.A1)

	assert.True(t, union.Equal(unionExpected))
	assert.True(t, intersect.Equal(intersectExpected))
	assert.True(t, diff.Equal(diffExpected))
}
