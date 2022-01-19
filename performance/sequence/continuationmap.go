package sequence

import (
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

type PitchMap map[*pitch.Pitch]*pitch.Pitch

func NewContinuationMap(currentNote, nextNote *note.Note, sep SeparationLevel) PitchMap {
	m := PitchMap{}
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

		for src, tgt := range OptimizeLinks(unlinked, untargeted) {
			m[src] = tgt
		}
	}

	return m
}
