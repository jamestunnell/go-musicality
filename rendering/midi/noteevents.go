package midi

import "github.com/jamestunnell/go-musicality/notation/rat"

type NoteEvents []Event

func (events NoteEvents) Offsets() rat.Rats {
	offsets := []rat.Rat{}

	for _, event := range events {
		found := false
		for _, offset := range offsets {
			if offset.Equal(event.Offset()) {
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

func (events NoteEvents) WithOffset(offset rat.Rat) NoteEvents {
	withOffset := NoteEvents{}

	for _, event := range events {
		if event.Offset().Equal(offset) {
			withOffset = append(withOffset, event)
		}
	}

	return withOffset
}
