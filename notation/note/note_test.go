package note_test

import (
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func TestEqual(t *testing.T) {
	n1 := note.Whole(pitch.A0)
	n2 := note.Half(pitch.A0)
	n3 := note.Whole(pitch.A0, pitch.A1)

	assert.True(t, n1.Equal(n1))
	assert.False(t, n1.Equal(n2))
	assert.False(t, n1.Equal(n3))

	n4 := note.Whole(pitch.A0)
	n5 := note.Whole(pitch.A0)
	n6 := note.Whole(pitch.A0)
	l := &note.Link{Source: pitch.A4, Target: pitch.A4, Type: note.LinkTie}

	n4.Attack = 0.33
	n5.Separation = 0.67
	n6.Links = append(n6.Links, l)

	assert.False(t, n1.Equal(n4))
	assert.False(t, n1.Equal(n5))
	assert.False(t, n1.Equal(n6))
}

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
	testNoteInvalid(t, "zero dur", rat.Zero(), func(n *note.Note) {})
	testNoteInvalid(t, "negative dur", rat.Zero(), func(n *note.Note) {})
	testNoteInvalid(t, "attack too high", big.NewRat(1, 4), func(n *note.Note) { n.Attack = note.ControlMax + 0.01 })
	testNoteInvalid(t, "attack too low", big.NewRat(1, 4), func(n *note.Note) { n.Attack = note.ControlMin - 0.01 })
	testNoteInvalid(t, "separation too high", big.NewRat(1, 4), func(n *note.Note) { n.Separation = note.ControlMax + 0.01 })
	testNoteInvalid(t, "separation too low", big.NewRat(1, 4), func(n *note.Note) { n.Separation = note.ControlMin - 0.01 })
}

type noteModFunc func(n *note.Note)

func TestNoteMarshalUnmarshalJSON(t *testing.T) {
	dur := big.NewRat(3, 2)
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
			l := &note.Link{Source: pitch.C0, Target: pitch.C0, Type: note.LinkTie}
			n.Links = append(n.Links, l)
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

	str := `{
		"duration": "1",
		"pitches": ["not-a-pitch"]
	}`

	testUnmarshalFail(t, "bad pitch", str)

	str = `{
		"duration": "1",
		"pitches": ["C5"],
		"links": {
			"not-a-pitch":{
				"type":"tie",
				"target": "C5"
			}
		}
	}`
	testUnmarshalFail(t, "bad link source pitch", str)

	str = strings.Replace(str, "not-a-pitch", "C5", 1)
	str = strings.Replace(str, `"target": "C5"`, `"target": "not-a-pitch"`, 1)

	testUnmarshalFail(t, "bad link target pitch", str)
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
		d, err := json.Marshal(n)

		if !assert.Nil(t, err) {
			t.Fatal(err.Error())
		}

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
