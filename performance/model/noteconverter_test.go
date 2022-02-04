package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/performance/model"
)

func TestNoteConverter(t *testing.T) {
	nc := model.NewNoteConverter()
	notes := []*note.Note{}

	notes2, err := nc.Process(notes)

	assert.NoError(t, err)
	assert.Empty(t, notes2)
}
