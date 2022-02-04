package model

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

type NoteConverter struct {
	// allowGlide   bool
	centsPerStep int
	offset       *big.Rat
	completed    []*Note
	continuing   map[*pitch.Pitch]*Note
}

type NoteConverterOptionFunc func(nc *NoteConverter)

func NewNoteConverter(opts ...NoteConverterOptionFunc) *NoteConverter {
	nc := &NoteConverter{}

	for _, opt := range opts {
		opt(nc)
	}

	return nc
}

// func OptionAllowPortamento() OptionFunc { return setAllowPortamento }

func OptionCentsPerStep(centsPerStep int) NoteConverterOptionFunc {
	return func(nc *NoteConverter) {
		nc.centsPerStep = centsPerStep
	}
}

// func setAllowPortamento(nc *NoteConverter) {
// 	nc.allowGlide = true
// }

func (nc *NoteConverter) Process(notes []*note.Note) ([]*Note, error) {
	nc.offset = big.NewRat(0, 1)
	nc.completed = []*Note{}
	nc.continuing = map[*pitch.Pitch]*Note{}

	for i, current := range notes {
		var next *note.Note

		if i == (len(notes) - 1) {
			// pretend the next note is a rest
			next = note.Quarter()
		} else {
			next = notes[i+1]
		}

		if err := nc.processNote(current, next); err != nil {
			err = fmt.Errorf("failed to process note %d: %w", i, err)

			return []*Note{}, err
		}

		nc.offset = new(big.Rat).Add(nc.offset, current.Duration)
	}

	if len(nc.continuing) > 0 {
		err := fmt.Errorf("continuing notes left over: %v", nc.continuing)

		return []*Note{}, err
	}

	for _, n := range nc.completed {
		n.Simplify()
	}

	return nc.completed, nil
}

func (nc *NoteConverter) processPitchDurs(current, next *pitch.Pitch, attack, separation float64, pitchDurs ...*PitchDur) {
	// pd := NewPitchDur(NewPitch(p, 0), dur)

	// is this a continuation?
	n := nc.continuing[current]

	if n != nil {
		n.PitchDurs = append(n.PitchDurs, pitchDurs...)
		n.Separation = separation

		delete(nc.continuing, current)
	} else {
		n = NewNote(nc.offset, pitchDurs...)

		n.Attack = attack
		n.Separation = separation
	}

	if next != nil {
		nc.continuing[next] = n
	} else {
		nc.completed = append(nc.completed, n)
	}
}

func (nc *NoteConverter) processNote(current, next *note.Note) error {
	a := current.Attack
	s := current.Separation
	dur := current.Duration

	for _, p := range current.Pitches.Pitches() {
		link := current.Links[p]

		if link == nil {
			nc.processPitchDurs(p, nil, a, s, NewPitchDur(NewPitch(p, 0), dur))

			continue
		}

		var target *pitch.Pitch

		if next.Pitches.Contains(link.Target) {
			target = link.Target
		}

		if link.Type == note.Tie || link.Type == note.Slur || link.Target == p {
			nc.processPitchDurs(p, nil, a, s, NewPitchDur(NewPitch(p, 0), dur))
		} else {
			switch link.Type {
			case note.Glide:
				pds := MakeSteps(dur, p, link.Target, nc.centsPerStep)

				nc.processPitchDurs(p, target, a, s, pds...)
			case note.StepSlurred, note.Step:
				pds := MakeSteps(dur, p, link.Target, CentsPerSemitoneInt)

				if link.Type == note.Step {
					for _, pd := range pds {
						nc.processPitchDurs(p, nil, a, s, pd)
					}
				} else {
					nc.processPitchDurs(p, target, a, s, pds...)
				}
			default:
				return fmt.Errorf("unknown link type '%s' from pitch '%s'", link.Type, p.String())
			}
		}
	}

	return nil
}
