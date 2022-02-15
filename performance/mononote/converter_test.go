package mononote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/performance/mononote"
)

func TestNoteConverter(t *testing.T) {
	nc := mononote.NewConverter()
	notes := []*note.Note{}

	notes2, err := nc.Process(notes)

	assert.NoError(t, err)
	assert.Empty(t, notes2)
}
