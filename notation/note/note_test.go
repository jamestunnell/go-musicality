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

func TestNewNonPositiveDuration(t *testing.T) {
	durs := []*big.Rat{big.NewRat(0, 1), big.NewRat(-1, 16)}
	for _, dur := range durs {
		note := note.New(dur)

		assert.Error(t, note.Validate())
	}
}

func TestNewPositiveDuration(t *testing.T) {
	durs := []*big.Rat{big.NewRat(1, 2), big.NewRat(1, 16)}
	for _, dur := range durs {
		note := note.New(dur)

		assert.NoError(t, note.Validate())
	}
}

func TestNoteArticulation(t *testing.T) {
	n := note.New(big.NewRat(1, 2), pitch.D6)

	assert.Equal(t, "", n.Articulation)

	n.Articulation = "unknown"

	assert.Error(t, n.Validate())

	n.Articulation = note.Accent

	assert.NoError(t, n.Validate())
}

func TestNoteIsRestIsMonophonic(t *testing.T) {
	dur := big.NewRat(3, 2)
	p1 := pitch.New(3, 2, 0)
	p2 := pitch.New(4, 0, 0)
	rest := note.New(dur)
	mono := note.New(dur, p1)
	poly := note.New(dur, p1, p2)

	require.NoError(t, rest.Validate())
	assert.True(t, rest.IsRest())
	assert.False(t, rest.IsMonophonic())

	require.NoError(t, mono.Validate())
	assert.False(t, mono.IsRest())
	assert.True(t, mono.IsMonophonic())

	require.NoError(t, poly.Validate())
	assert.False(t, poly.IsRest())
	assert.False(t, poly.IsMonophonic())
}

func TestNoteMarshalUnmarshalJSON(t *testing.T) {
	dur := big.NewRat(3, 2)
	p1 := pitch.New(3, 2, 0)
	p2 := pitch.New(4, 0, 0)

	testNoteMarshalUnmarshalJSON(t, "rest", note.New(dur))
	testNoteMarshalUnmarshalJSON(t, "mono", note.New(dur, p1))
	testNoteMarshalUnmarshalJSON(t, "poly", note.New(dur, p1, p2))
}

func testNoteMarshalUnmarshalJSON(t *testing.T, name string, n *note.Note) {
	t.Run(name, func(t *testing.T) {
		t.Log(n.Duration.String())

		require.NoError(t, n.Validate())

		d, err := json.Marshal(n)

		require.Nil(t, err)

		t.Log(string(d))

		assert.Greater(t, len(d), 0)

		var n2 note.Note
		err = json.Unmarshal(d, &n2)

		require.Nil(t, err)

		compareNotes(t, n, &n2)
	})
}

func TestNoteDot(t *testing.T) {
	n := note.Quarter(pitch.C4)

	assert.Equal(t, 0, n.Duration.Cmp(big.NewRat(1, 4)))

	n.Dot()

	assert.Equal(t, 0, n.Duration.Cmp(big.NewRat(3, 8)))

	n.Dot()

	assert.Equal(t, 0, n.Duration.Cmp(big.NewRat(9, 16)))
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
