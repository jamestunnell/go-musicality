package note_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func TestNoteNonPositiveDuration(t *testing.T) {
	durs := []*big.Rat{big.NewRat(0, 1), big.NewRat(-1, 16)}
	for _, dur := range durs {
		note := note.New(dur)

		assert.Error(t, note.Validate())
	}
}

func TestNoteIsRestIsMonophonic(t *testing.T) {
	rest, mono, poly := testNoteSetup(t)

	assert.True(t, rest.IsRest())
	assert.False(t, rest.IsMonophonic())

	assert.False(t, mono.IsRest())
	assert.True(t, mono.IsMonophonic())

	assert.False(t, poly.IsRest())
	assert.False(t, poly.IsMonophonic())
}

func TestNoteMarshalUnmarshalJSON(t *testing.T) {
	dur := big.NewRat(3, 2)
	p1 := pitch.New(3, 2, 0)
	p2 := pitch.New(4, 0, 0)
	rest := note.New(dur)

	require.NoError(t, rest.Validate())

	mono := note.New(dur, p1)

	require.NoError(t, mono.Validate())

	poly := note.New(dur, p1, p2)

	require.NoError(t, poly.Validate())

	for _, n := range []*note.Note{rest, mono, poly} {
		d, err := json.Marshal(n)

		if !assert.Nil(t, err) {
			continue
		}

		t.Log(string(d))

		assert.Greater(t, len(d), 0)

		var n2 note.Note
		err = json.Unmarshal(d, &n2)

		if !assert.Nil(t, err) {
			continue
		}

		compareNotes(t, n, &n2)
	}
}

func testNoteSetup(t *testing.T) (rest, mono, poly *note.Note) {
	dur := big.NewRat(3, 2)
	p1 := pitch.New(3, 2, 0)
	p2 := pitch.New(4, 0, 0)
	rest = note.New(dur)

	require.NoError(t, rest.Validate())

	mono = note.New(dur, p1)

	require.NoError(t, mono.Validate())

	poly = note.New(dur, p1, p2)

	require.NoError(t, poly.Validate())

	return rest, mono, poly
}

func compareNotes(t *testing.T, n1, n2 *note.Note) bool {
	ok := assert.Equal(t, len(n1.Pitches), len(n2.Pitches))
	if ok {
		for i, p := range n2.Pitches {
			if !assert.Equal(t, *n1.Pitches[i], *p) {
				ok = false
			}
		}
	}

	return assert.Equal(t, n1.Duration, n2.Duration) && ok
}
