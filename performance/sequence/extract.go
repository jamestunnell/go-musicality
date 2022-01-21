package sequence

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func Extract(notes []*note.Note) []*Sequence {
	offset := big.NewRat(0, 1)
	completedSeqs := []*Sequence{}
	continuingSequences := map[*pitch.Pitch]*Sequence{}

	for i, currentNote := range notes {
		var nextNote *note.Note

		if i == (len(notes) - 1) {
			// invent an imaginary next note
			nextNote = note.Quarter()
		} else {
			nextNote = notes[i+1]
		}

		continuationMap := NewContinuationMap(currentNote, nextNote, currentNote.Separation)

		newContinuingSequences := map[*pitch.Pitch]*Sequence{}

		for _, p := range currentNote.Pitches.Pitches() {
			var seq *Sequence

			if continuingSeq, found := continuingSequences[p]; found {
				seq = continuingSeq

				elems := elements(currentNote.Duration, p, note.AttackMin, currentNote.Links[p])

				seq.Elements = append(seq.Elements, elems...)

				delete(continuingSequences, p)
			} else {
				elems := elements(currentNote.Duration, p, currentNote.Attack, currentNote.Links[p])

				seq = New(offset, elems...)
			}

			if tgt, found := continuationMap[p]; found {
				newContinuingSequences[tgt] = seq
			} else {
				completedSeqs = append(completedSeqs, seq)
			}
		}

		if len(continuingSequences) > 0 {
			panic("shouldn't be any continuing sequences")
		}

		continuingSequences = newContinuingSequences

		offset = new(big.Rat).Add(offset, currentNote.Duration)
	}

	for _, seq := range completedSeqs {
		seq.Simplify()
	}

	return completedSeqs
}

func elements(dur *big.Rat, p *pitch.Pitch, attack float64, link *note.Link) []*Element {
	var elems []*Element

	if link != nil {
		switch link.Type {
		case note.Portamento:
			panic("not supported")
		case note.Glissando:
			// reserve 25% of the original note duration for the starting pitch
			elems = []*Element{
				{
					Duration: new(big.Rat).Mul(dur, big.NewRat(1, 4)),
					Pitch:    p,
					Attack:   attack,
				},
			}

			diff := link.Target.Diff(p)
			semitones := Abs(diff)
			incr := diff / semitones
			subdur := new(big.Rat).Mul(dur, big.NewRat(3, 4*int64(semitones)))
			lastElem := elems[0]

			for i := 0; i < semitones; i++ {
				elem := &Element{
					Duration: subdur,
					Pitch:    lastElem.Pitch.Transpose(incr),
					Attack:   note.AttackMin,
				}

				elems = append(elems, elem)

				lastElem = elem
			}

		case note.Tie:
			elems = []*Element{
				{
					Duration: dur,
					Pitch:    p,
					Attack:   note.AttackMin,
				},
			}
		}
	} else {
		elems = []*Element{
			{
				Duration: dur,
				Pitch:    p,
				Attack:   attack,
			},
		}
	}

	return elems
}
