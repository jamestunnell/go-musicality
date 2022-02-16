package pitch_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func TestPitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C3, pitch.C2}

	sort.Sort(ps)

	assert.True(t, ps[0].Equal(pitch.C1))
	assert.True(t, ps[1].Equal(pitch.C2))
	assert.True(t, ps[2].Equal(pitch.C3))

	pStrings := ps.Strings()

	assert.Equal(t, []string{"C1", "C2", "C3"}, pStrings)
}

func TestPitchesStrings(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C3, pitch.C2}

	sort.Sort(ps)

	assert.True(t, ps[0].Equal(pitch.C1))
	assert.True(t, ps[1].Equal(pitch.C2))
	assert.True(t, ps[2].Equal(pitch.C3))
}
