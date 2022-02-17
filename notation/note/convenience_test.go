package note_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
)

func TestConvenience(t *testing.T) {
	assert.True(t, note.Whole().Duration.Equal(rat.New(1, 1)))
	assert.True(t, note.Half().Duration.Equal(rat.New(1, 2)))
	assert.True(t, note.Quarter().Duration.Equal(rat.New(1, 4)))
	assert.True(t, note.Eighth().Duration.Equal(rat.New(1, 8)))
	assert.True(t, note.Sixteenth().Duration.Equal(rat.New(1, 16)))
}
