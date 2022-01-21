package sequence

import (
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

type PitchMap map[*pitch.Pitch]*pitch.Pitch

func NewContinuationMap(currentNote, nextNote *note.Note, sep float64) PitchMap {
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

	if sep == note.SeparationMin {
		unlinked := currentNote.Pitches.Diff(linked)
		untargeted := nextNote.Pitches.Diff(targeted)
		pm := OptimizeLinks(unlinked, untargeted)

		for src, tgt := range pm {
			m[src] = tgt
		}
	}

	return m
}
