package notation_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/jamestunnell/go-musicality/pkg/notation"
	"github.com/stretchr/testify/assert"
)

func TestNoteNonPositiveDuration(t *testing.T) {
	durs := []*big.Rat{big.NewRat(0, 1), big.NewRat(-1, 16)}
	for _, dur := range durs {
		note, err := notation.NewNote(dur)

		assert.Nil(t, note)
		assert.NotNil(t, err)
	}
}

func TestNoteIsRestIsMonophonic(t *testing.T) {
	dur := big.NewRat(2, 1)

	rest, err := notation.NewNote(dur)

	assert.Nil(t, err)
	if !assert.NotNil(t, rest) {
		return
	}

	assert.True(t, rest.IsRest())
	assert.False(t, rest.IsMonophonic())

	mono, err := notation.NewNote(dur, notation.NewPitch(0, 0, 0))

	assert.Nil(t, err)
	if !assert.NotNil(t, mono) {
		return
	}

	assert.False(t, mono.IsRest())
	assert.True(t, mono.IsMonophonic())
}

func TestNoteMarshalUnmarshalJSON(t *testing.T) {
	dur := big.NewRat(3, 2)
	p1 := notation.NewPitch(3, 2, 0)
	p2 := notation.NewPitch(4, 0, 0)
	rest, err := notation.NewNote(dur)

	assert.NotNil(t, rest)
	assert.Nil(t, err)

	mono, err := notation.NewNote(dur, p1)

	assert.NotNil(t, mono)
	assert.Nil(t, err)

	poly, err := notation.NewNote(dur, p1, p2)

	assert.NotNil(t, poly)
	assert.Nil(t, err)

	for _, n := range []*notation.Note{rest, mono, poly} {
		d, err := json.Marshal(n)

		if !assert.Nil(t, err) {
			continue
		}

		t.Log(string(d))

		assert.Greater(t, len(d), 0)

		var n2 notation.Note
		err = json.Unmarshal(d, &n2)

		if !assert.Nil(t, err) {
			continue
		}

		compareNotes(t, n, &n2)
	}
}

func compareNotes(t *testing.T, n1, n2 *notation.Note) bool {
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
