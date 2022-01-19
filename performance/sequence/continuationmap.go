package sequence

import (
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

type ContinuationMap map[*pitch.Pitch]*pitch.Pitch

func NewContinuationMap(currentNote, nextNote *note.Note, sep SeparationLevel) ContinuationMap {
	m := ContinuationMap{}
	linked := pitch.NewSet()
	targeted := pitch.NewSet()

	for p, link := range currentNote.Links {
		if nextNote.Pitches.Contains(link.Target) {
			m[p] = link.Target

			linked.Add(p)
			targeted.Add(link.Target)
		}
	}

	if sep == Separation0 {
		unlinked := currentNote.Pitches.Diff(linked)
		untargeted := nextNote.Pitches.Diff(targeted)

		//   Optimization.linking(unlinked, untargeted).each do |pitch,tgt_pitch|
		//     map[pitch] = tgt_pitch
		//   end
	}

	return m
}
