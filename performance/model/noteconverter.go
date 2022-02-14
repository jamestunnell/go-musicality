package model

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

type NoteConverter struct {
	replaceSlursAndGlides bool
	centsPerStep          int
	offset                rat.Rat
	completed             []*Note
	continuing            map[*pitch.Pitch]*Note
}

type NoteConverterOptionFunc func(nc *NoteConverter)

const DefaultCentsPerStep = 10

func NewNoteConverter(opts ...NoteConverterOptionFunc) *NoteConverter {
	nc := &NoteConverter{
		replaceSlursAndGlides: false,
		centsPerStep:          DefaultCentsPerStep,
	}

	for _, opt := range opts {
		opt(nc)
	}

	return nc
}

func OptionReplaceSlursAndGlides() NoteConverterOptionFunc { return setReplaceSlursAndGlides }

func OptionCentsPerStep(centsPerStep int) NoteConverterOptionFunc {
	return func(nc *NoteConverter) {
		if centsPerStep < 1 {
			centsPerStep = 1
		}

		nc.centsPerStep = centsPerStep
	}
}

func setReplaceSlursAndGlides(nc *NoteConverter) {
	nc.replaceSlursAndGlides = true
}

func (nc *NoteConverter) Process(notes []*note.Note) ([]*Note, error) {
	nc.offset = rat.Zero()
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

		nc.offset.Accum(current.Duration)
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

func (nc *NoteConverter) processNote(current, next *note.Note) error {
	a := current.Attack
	s := current.Separation
	dur := current.Duration

	for _, p := range current.Pitches.Pitches() {
		link, found := current.Links.FindBySource(p)

		if found && nc.replaceSlursAndGlides {
			switch link.Type {
			case note.LinkSlur:
				link = nil
				s = note.ControlMin
			case note.LinkStepSlurred, note.LinkGlide:
				link = &note.Link{
					Type:   note.LinkStep,
					Target: link.Target,
				}
			}
		}

		var target *pitch.Pitch

		if link != nil && next.Pitches.Contains(link.Target) {
			target = link.Target
		}

		// no link or a link where pitch doesn't change can be handled simply
		if link == nil || link.Target == p {
			nc.processPitchDurs(p, target, a, s, NewPitchDur(NewPitch(p, 0), dur))

			continue
		}

		switch link.Type {
		case note.LinkTie, note.LinkSlur:
			nc.processPitchDurs(p, target, a, s, NewPitchDur(NewPitch(p, 0), dur))
		case note.LinkGlide:
			pds := MakeSteps(dur, p, link.Target, nc.centsPerStep)

			nc.processPitchDurs(p, target, a, s, pds...)
		case note.LinkStepSlurred, note.LinkStep:
			pds := MakeSteps(dur, p, link.Target, CentsPerSemitoneInt)

			if link.Type == note.LinkStep {
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

	return nil
}

func (nc *NoteConverter) processPitchDurs(current, next *pitch.Pitch, attack, separation float64, pitchDurs ...*PitchDur) {
	// is this a continuation?
	n := nc.continuing[current]

	if n != nil {
		n.PitchDurs = append(n.PitchDurs, pitchDurs...)
		n.Separation = separation

		delete(nc.continuing, current)
	} else {
		n = NewNote(nc.offset.Clone(), pitchDurs...)

		n.Attack = attack
		n.Separation = separation
	}

	if next != nil {
		nc.continuing[next] = n
	} else {
		nc.completed = append(nc.completed, n)
	}
}
