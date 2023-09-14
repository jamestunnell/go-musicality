package midi

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
)

type NoteEvents []Event

func (events NoteEvents) Offsets() rat.Rats {
	offsets := rat.Rats{}

	for _, event := range events {
		found := false
		for _, offset := range offsets {
			if rat.IsEqual(offset, event.Offset()) {
				found = true

				break
			}
		}

		if !found {
			offsets = append(offsets, event.Offset())
		}
	}

	return offsets
}

func (events NoteEvents) WithOffset(offset *big.Rat) NoteEvents {
	withOffset := NoteEvents{}

	for _, event := range events {
		if rat.IsEqual(event.Offset(), offset) {
			withOffset = append(withOffset, event)
		}
	}

	return withOffset
}
