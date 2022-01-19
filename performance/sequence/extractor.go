package sequence

import (
	"github.com/jamestunnell/go-musicality/notation/note"
)

func Extract(notes []*note.Note) []*Sequence {
	seqs := []*Sequence{}
	// offset := big.NewRat(0, 1)
	// slurring := false

	// for i, currentNote := range notes {
	// 	if note.Slurs {
	// 		slurring = true
	// 	}

	// 	dur := currentNote.Duration
	// 	attack := Attack(currentNote.Articulation)
	// 	separation := Separation(currentNote.Articulation, slurring)

	// 	var nextNote

	// 	if i == (len(notes)-1) {
	// 		// invent an imaginary next note
	// 		nextNote = note.Quarter()
	// 	} else {
	// 		nextNote = notes[i+1]
	// 	}

	// 	continuationMap := ContinuationMap(currentNote, nextNote, separation)

	// 	newContinuingSequences := map[]

	// 	// TODO - more stuff

	// 	// Update the offset
	// 	offset = new(big.Rat).Add(offset, dur)
	// }

	return seqs
}
