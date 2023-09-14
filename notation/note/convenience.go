package note

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func Whole(pitches ...*pitch.Pitch) *Note {
	return New(big.NewRat(1, 1), pitches...)
}

func Half(pitches ...*pitch.Pitch) *Note {
	return New(big.NewRat(1, 2), pitches...)
}

func Quarter(pitches ...*pitch.Pitch) *Note {
	return New(big.NewRat(1, 4), pitches...)
}

func Eighth(pitches ...*pitch.Pitch) *Note {
	return New(big.NewRat(1, 8), pitches...)
}

func Sixteenth(pitches ...*pitch.Pitch) *Note {
	return New(big.NewRat(1, 16), pitches...)
}

func (n *Note) Tie(p *pitch.Pitch) *Note {
	l := &Link{Type: LinkTie, Source: p, Target: p}

	return n.Link(l)
}

func (n *Note) Slur(src, tgt *pitch.Pitch) *Note {
	l := &Link{Type: LinkSlur, Source: src, Target: tgt}

	return n.Link(l)
}

func (n *Note) Glide(src, tgt *pitch.Pitch) *Note {
	l := &Link{Type: LinkGlide, Source: src, Target: tgt}

	return n.Link(l)
}

func (n *Note) Step(src, tgt *pitch.Pitch) *Note {
	l := &Link{Type: LinkStep, Source: src, Target: tgt}

	return n.Link(l)
}

func (n *Note) StepSlurred(src, tgt *pitch.Pitch) *Note {
	l := &Link{Type: LinkStepSlurred, Source: src, Target: tgt}

	return n.Link(l)
}

func (n *Note) Link(l *Link) *Note {
	n.Links = append(n.Links, l)

	return n
}
