package notegen

import (
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/composition/pitchgen"
	"github.com/jamestunnell/go-musicality/composition/rhythmgen"
	"github.com/jamestunnell/go-musicality/notation/note"
)

type NoteGenerator struct {
	pitchGen  pitchgen.PitchGenerator
	rhythmGen rhythmgen.RhythmGenerator
}

func NewNoteGenerator(
	rg rhythmgen.RhythmGenerator,
	pg pitchgen.PitchGenerator) *NoteGenerator {
	return &NoteGenerator{
		rhythmGen: rg,
		pitchGen:  pg,
	}
}
func (ng *NoteGenerator) MakeNotes(dur rat.Rat) note.Notes {
	durs := ng.rhythmGen.MakeRhythm(dur)
	n := len(durs)
	pitches := ng.pitchGen.MakePitches(n)
	notes := make(note.Notes, n)

	for i := 0; i < n; i++ {
		notes[i] = note.New(durs[i], pitches[i])
	}

	return notes
}
