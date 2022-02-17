package note_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
)

func TestNotesTotalDuration(t *testing.T) {
	notes := note.Notes{}

	assert.True(t, notes.TotalDuration().Zero())

	notes = append(notes, note.Eighth())

	assert.True(t, notes.TotalDuration().Equal(rat.New(1, 8)))
}
