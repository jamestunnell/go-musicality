package note_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/duration"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func TestNoteNonPositiveDuration(t *testing.T) {
	durs := []*duration.Duration{duration.Zero(), duration.New(-1, 16)}
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
	dur := duration.New(3, 2)
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

		n2 := &note.Note{Duration: duration.Zero()}
		err = json.Unmarshal(d, n2)

		require.Nil(t, err)

		compareNotes(t, n, n2)
	})
}

func TestNoteDot(t *testing.T) {
	n := note.Quarter(pitch.C4)
	n2 := n.Dot()
	n3 := n2.Dot()

	assert.True(t, n.Duration.Equal(duration.New(1, 4)))
	assert.True(t, n2.Duration.Equal(duration.New(3, 8)))
	assert.True(t, n3.Duration.Equal(duration.New(9, 16)))
}

func testNoteSetup(t *testing.T) (rest, mono, poly *note.Note) {
	dur := duration.New(3, 2)
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
