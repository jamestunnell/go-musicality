package note_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestNoteValid(t *testing.T) {
	testNoteValid(t, "rest", rat.New(1, 2), func(n *note.Note) {})
	testNoteValid(t, "monophonic", rat.New(1, 2), func(n *note.Note) { n.Pitches.Add(pitch.A0) })
	testNoteValid(t, "polyphonic", rat.New(1, 2), func(n *note.Note) {
		n.Pitches.Add(pitch.A0)
		n.Pitches.Add(pitch.C0)
	})
	testNoteValid(t, "min attack", rat.New(1, 2), func(n *note.Note) {
		n.Attack = note.ControlMin
	})
	testNoteValid(t, "max attack", rat.New(1, 2), func(n *note.Note) {
		n.Attack = note.ControlMax
	})
	testNoteValid(t, "min separation", rat.New(1, 2), func(n *note.Note) {
		n.Separation = note.ControlMin
	})
	testNoteValid(t, "max separation", rat.New(1, 2), func(n *note.Note) {
		n.Separation = note.ControlMax
	})
}

func TestNoteInvalid(t *testing.T) {
	testNoteInvalid(t, "zero dur", rat.Zero(), func(n *note.Note) {})
	testNoteInvalid(t, "negative dur", rat.Zero(), func(n *note.Note) {})
	testNoteInvalid(t, "attack too high", rat.New(1, 4), func(n *note.Note) { n.Attack = note.ControlMax + 0.01 })
	testNoteInvalid(t, "attack too low", rat.New(1, 4), func(n *note.Note) { n.Attack = note.ControlMin - 0.01 })
	testNoteInvalid(t, "separation too high", rat.New(1, 4), func(n *note.Note) { n.Separation = note.ControlMax + 0.01 })
	testNoteInvalid(t, "separation too low", rat.New(1, 4), func(n *note.Note) { n.Separation = note.ControlMin - 0.01 })
}

type noteModFunc func(n *note.Note)

func TestNoteMarshalUnmarshalJSON(t *testing.T) {
	dur := rat.New(3, 2)
	p1 := pitch.New(3, 2)
	p2 := pitch.New(4, 0)

	cases := map[string]noteModFunc{
		"rest": func(n *note.Note) {},
		"monophonic": func(n *note.Note) {
			n.Pitches.Add(p1)
		},
		"polyphonic": func(n *note.Note) {
			n.Pitches.Add(p1)
			n.Pitches.Add(p2)
		},
		"valid, not normal attack": func(n *note.Note) {
			n.Attack = note.ControlNormal + 0.01
		},
		"valid, not normal separation": func(n *note.Note) {
			n.Separation = note.ControlNormal + 0.01
		},
		"invalid attack": func(n *note.Note) {
			n.Attack = note.ControlMax + 0.01
		},
		"invalid separation": func(n *note.Note) {
			n.Separation = note.ControlMax + 0.01
		},
		"with link": func(n *note.Note) {
			n.Links[p1] = &note.Link{Type: note.Tie, Target: p1}
		},
	}

	for name, mod := range cases {
		n := note.New(dur)

		mod(n)

		testNoteMarshalUnmarshalJSON(t, name, n)
	}
}

func TestUnmarshalFail(t *testing.T) {
	testUnmarshalFail(t, "wrong type", `"not-an-object"`)

	str := `{"duration": "1", "links": {}}` //"not-a-pitch":{"type":"tie","target":"C5"}
	testUnmarshalFail(t, "bad link pitch", str)

	// str = `{"duration": "1", "pitches": ["C5"], "links": {"C5":{"type":"tie","target":"not-a-pitch"}}}`
	// testUnmarshalFail(t, "bad link target pitch", str)
}

func testUnmarshalFail(t *testing.T, name, jsonStr string) {
	t.Run(name, func(t *testing.T) {
		var n note.Note

		d := []byte(jsonStr)
		err := json.Unmarshal(d, &n)

		assert.Error(t, err)
	})
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

func testNoteValid(t *testing.T, name string, dur rat.Rat, mod func(n *note.Note)) {
	t.Run(name, func(t *testing.T) {
		n := note.New(dur)

		mod(n)

		assert.Nil(t, n.Validate())
	})
}

func testNoteInvalid(t *testing.T, name string, dur rat.Rat, mod func(n *note.Note)) {
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
