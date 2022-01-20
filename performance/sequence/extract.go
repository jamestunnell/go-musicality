package sequence

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func Extract(notes []*note.Note) []*Sequence {
	offset := big.NewRat(0, 1)
	slurring := false

	completedSeqs := []*Sequence{}
	continuingSequences := map[*pitch.Pitch]*Sequence{}

	for i, currentNote := range notes {
		if currentNote.Slurs {
			slurring = true
		}

		dur := currentNote.Duration
		attack := Attack(currentNote.Articulation)
		separation := Separation(currentNote.Articulation, slurring)

		var nextNote *note.Note

		if i == (len(notes) - 1) {
			// invent an imaginary next note
			nextNote = note.Quarter()
		} else {
			nextNote = notes[i+1]
		}

		continuationMap := NewContinuationMap(currentNote, nextNote, separation)

		newContinuingSequences := map[*pitch.Pitch]*Sequence{}

		for _, p := range currentNote.Pitches.Pitches() {
			var seq *Sequence

			if continuingSeq, found := continuingSequences[p]; found {
				seq = continuingSeq

				seq.Elements = append(seq.Elements, elements(currentNote, p, AttackNone)...)

				delete(continuingSequences, p)
			} else {
				seq = New(offset, elements(currentNote, p, attack)...)
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

		// Update the offset
		offset = new(big.Rat).Add(offset, dur)
	}

	for _, seq := range completedSeqs {
		seq.Simplify()
	}

	return completedSeqs
}

func elements(n *note.Note, p *pitch.Pitch, attack float64) []*Element {
	d := n.Duration
	l := n.Links[p]

	var elems []*Element

	if l != nil {
		switch l.Type {
		case note.Portamento:
			panic("not supported")
		case note.Glissando:
			// reserve 25% of the original note duration for the starting pitch
			elems = []*Element{
				{
					Duration: new(big.Rat).Mul(d, big.NewRat(1,4)),
					Pitch:    p,
					Attack:   attack,
				},
			}

			diff := l.Target.Diff(p)
			semitones := Abs(diff)
			incr := diff / semitones
			subdur := new(big.Rat).Mul(d, big.NewRat(3,4*int64(semitones)))
			lastElem := elems[0]

			for i := 0; i < semitones; i++ {
				elem := &Element{
					Duration: subdur,
					Pitch:    lastElem.Pitch.Transpose(incr),
					Attack:   AttackNone,
				}

				elems = append(elems, elem)

				lastElem = elem
			}

		case note.Tie:
			elems = []*Element{
				{
					Duration: d,
					Pitch:    p,
					Attack:   AttackNone,
				},
			}
		}
	} else {
		elems = []*Element{
			{
				Duration: d,
				Pitch:    p,
				Attack:   attack,
			},
		}
	}

	return elems
}
