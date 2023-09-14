package note_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
)

func TestConvenience(t *testing.T) {
	assert.True(t, rat.IsEqual(note.Whole().Duration, big.NewRat(1, 1)))
	assert.True(t, rat.IsEqual(note.Half().Duration, big.NewRat(1, 2)))
	assert.True(t, rat.IsEqual(note.Quarter().Duration, big.NewRat(1, 4)))
	assert.True(t, rat.IsEqual(note.Eighth().Duration, big.NewRat(1, 8)))
	assert.True(t, rat.IsEqual(note.Sixteenth().Duration, big.NewRat(1, 16)))
}
