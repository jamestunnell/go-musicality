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

func TestNoteValid(t *testing.T) {
	testNoteValid(t, "rest", big.NewRat(1, 2), func(n *note.Note) {})
	testNoteValid(t, "monophonic", big.NewRat(1, 2), func(n *note.Note) { n.Pitches.Add(pitch.A0) })
	testNoteValid(t, "polyphonic", big.NewRat(1, 2), func(n *note.Note) {
		n.Pitches.Add(pitch.A0)
		n.Pitches.Add(pitch.C0)
	})
	testNoteValid(t, "min attack", big.NewRat(1, 2), func(n *note.Note) {
		n.Attack = note.ControlMin
	})
	testNoteValid(t, "max attack", big.NewRat(1, 2), func(n *note.Note) {
		n.Attack = note.ControlMax
	})
	testNoteValid(t, "min separation", big.NewRat(1, 2), func(n *note.Note) {
		n.Separation = note.ControlMin
	})
	testNoteValid(t, "max separation", big.NewRat(1, 2), func(n *note.Note) {
		n.Separation = note.ControlMax
	})
}

func TestNoteInvalid(t *testing.T) {
	testNoteInvalid(t, "zero dur", big.NewRat(0, 1), func(n *note.Note) {})
	testNoteInvalid(t, "negative dur", big.NewRat(0, 1), func(n *note.Note) {})
	testNoteInvalid(t, "attack too high", big.NewRat(1, 4), func(n *note.Note) { n.Attack = note.ControlMax + 0.01 })
	testNoteInvalid(t, "attack too low", big.NewRat(1, 4), func(n *note.Note) { n.Attack = note.ControlMin - 0.01 })
	testNoteInvalid(t, "separation too high", big.NewRat(1, 4), func(n *note.Note) { n.Separation = note.ControlMax + 0.01 })
	testNoteInvalid(t, "separation too low", big.NewRat(1, 4), func(n *note.Note) { n.Separation = note.ControlMin - 0.01 })
}

func TestNoteMarshalUnmarshalJSON(t *testing.T) {
	dur := big.NewRat(3, 2)
	p1 := pitch.New(3, 2)
	p2 := pitch.New(4, 0)

	testNoteMarshalUnmarshalJSON(t, "rest", note.New(dur))
	testNoteMarshalUnmarshalJSON(t, "mono", note.New(dur, p1))
	testNoteMarshalUnmarshalJSON(t, "poly", note.New(dur, p1, p2))

	n := note.New(dur, p1)

	n.Attack = note.ControlMax + 0.1

	testNoteMarshalUnmarshalJSON(t, "invalid attack", n)

	n.Separation = note.ControlMax + 0.1

	testNoteMarshalUnmarshalJSON(t, "invalid separation", n)

	n.Attack = note.ControlNormal

	testNoteMarshalUnmarshalJSON(t, "normal attack", n)

	n.Separation = note.ControlNormal

	testNoteMarshalUnmarshalJSON(t, "normal separation", n)
}

func testNoteMarshalUnmarshalJSON(t *testing.T, name string, n *note.Note) {
	t.Run(name, func(t *testing.T) {
		t.Log(n.Duration.String())

		d, err := json.Marshal(n)

		if !assert.Nil(t, err) {
			t.Fatal(err.Error())
		}

		t.Log(string(d))

		assert.Greater(t, len(d), 0)

		var n2 note.Note
		err = json.Unmarshal(d, &n2)

		require.Nil(t, err)

		compareNotes(t, n, &n2)
	})
}

func testNoteValid(t *testing.T, name string, dur *big.Rat, mod func(n *note.Note)) {
	t.Run(name, func(t *testing.T) {
		n := note.New(dur)

		mod(n)

		assert.Nil(t, n.Validate())
	})
}

func testNoteInvalid(t *testing.T, name string, dur *big.Rat, mod func(n *note.Note)) {
	t.Run(name, func(t *testing.T) {
		n := note.New(dur)

		mod(n)

		assert.NotNil(t, n.Validate())
	})
}

func compareNotes(t *testing.T, n1, n2 *note.Note) {
	intersect := n1.Pitches.Intersect(n2.Pitches)
	assert.Equal(t, intersect.Len(), n1.Pitches.Len())
	assert.Equal(t, n1.Duration, n2.Duration)
	assert.Equal(t, n1.Attack, n2.Attack)
	assert.Equal(t, n1.Separation, n2.Separation)
}
