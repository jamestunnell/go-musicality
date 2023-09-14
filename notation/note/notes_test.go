package note_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
)

func TestNotesTotalDuration(t *testing.T) {
	notes := note.Notes{}

	assert.True(t, rat.IsZero(notes.TotalDuration()))

	notes = append(notes, note.Eighth())

	assert.True(t, rat.IsEqual(notes.TotalDuration(), big.NewRat(1, 8)))
}
