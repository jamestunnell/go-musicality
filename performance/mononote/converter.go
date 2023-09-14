package mononote

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/centpitch"
)

type Converter struct {
	replaceSlursAndGlides bool
	centsPerStep          int
	offset                *big.Rat
	completed             []*MonoNote
	continuing            map[*pitch.Pitch]*MonoNote
}

type ConverterOptionFunc func(nc *Converter)

const DefaultCentsPerStep = 10

func NewConverter(opts ...ConverterOptionFunc) *Converter {
	nc := &Converter{
		replaceSlursAndGlides: false,
		centsPerStep:          DefaultCentsPerStep,
	}

	for _, opt := range opts {
		opt(nc)
	}

	return nc
}

func OptionReplaceSlursAndGlides() ConverterOptionFunc { return setReplaceSlursAndGlides }

func OptionCentsPerStep(centsPerStep int) ConverterOptionFunc {
	return func(nc *Converter) {
		if centsPerStep <= 0 {
			centsPerStep = 1
		}

		if centsPerStep > centpitch.CentsPerSemitoneFlt {
			centsPerStep = centpitch.CentsPerSemitoneFlt
		}

		nc.centsPerStep = centsPerStep
	}
}

func setReplaceSlursAndGlides(nc *Converter) {
	nc.replaceSlursAndGlides = true
}

func (nc *Converter) CentsPerStep() int {
	return nc.centsPerStep
}

func (nc *Converter) ReplaceSlursAndGlides() bool {
	return nc.replaceSlursAndGlides
}

func (nc *Converter) Process(notes []*note.Note) ([]*MonoNote, error) {
	nc.offset = rat.Zero()
	nc.completed = []*MonoNote{}
	nc.continuing = map[*pitch.Pitch]*MonoNote{}

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

			return []*MonoNote{}, err
		}

		nc.offset = rat.Add(nc.offset, current.Duration)
	}

	if len(nc.continuing) > 0 {
		err := fmt.Errorf("continuing notes left over: %v", nc.continuing)

		return []*MonoNote{}, err
	}

	for _, n := range nc.completed {
		n.Simplify()
	}

	return nc.completed, nil
}

func (nc *Converter) processNote(current, next *note.Note) error {
	a := current.Attack
	s := current.Separation
	dur := current.Duration

	for _, l := range current.Links {
		switch l.Type {
		case note.LinkTie, note.LinkGlide, note.LinkSlur, note.LinkStep, note.LinkStepSlurred:
		default:
			return fmt.Errorf("note link has unknown type '%s'", l.Type)
		}
	}

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
			nc.processPitchDurs(p, target, a, s, NewPitchDur(centpitch.New(p, 0), dur))

			continue
		}

		switch link.Type {
		case note.LinkTie, note.LinkSlur:
			nc.processPitchDurs(p, target, a, s, NewPitchDur(centpitch.New(p, 0), dur))
		case note.LinkGlide:
			pds := MakeSteps(dur, p, link.Target, nc.centsPerStep)

			nc.processPitchDurs(p, target, a, s, pds...)
		case note.LinkStepSlurred, note.LinkStep:
			pds := MakeSteps(dur, p, link.Target, centpitch.CentsPerSemitoneInt)

			if link.Type == note.LinkStep {
				for _, pd := range pds {
					nc.processPitchDurs(p, nil, a, s, pd)
				}
			} else {
				nc.processPitchDurs(p, target, a, s, pds...)
			}
		}
	}

	return nil
}

func (nc *Converter) processPitchDurs(
	current, next *pitch.Pitch,
	attack, separation float64,
	pitchDurs ...*PitchDur) {
	// is this a continuation?
	n := nc.continuing[current]

	if n != nil {
		n.PitchDurs = append(n.PitchDurs, pitchDurs...)
		n.Separation = separation

		delete(nc.continuing, current)
	} else {
		n = New(nc.offset, pitchDurs...)

		n.Attack = attack
		n.Separation = separation
	}

	if next != nil {
		nc.continuing[next] = n
	} else {
		nc.completed = append(nc.completed, n)
	}
}
