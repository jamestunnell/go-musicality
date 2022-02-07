package note

import (
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func Whole(pitches ...*pitch.Pitch) *Note {
	return New(rat.New(1, 1), pitches...)
}

func Half(pitches ...*pitch.Pitch) *Note {
	return New(rat.New(1, 2), pitches...)
}

func Quarter(pitches ...*pitch.Pitch) *Note {
	return New(rat.New(1, 4), pitches...)
}

func Eighth(pitches ...*pitch.Pitch) *Note {
	return New(rat.New(1, 8), pitches...)
}

func Sixteenth(pitches ...*pitch.Pitch) *Note {
	return New(rat.New(1, 16), pitches...)
}
