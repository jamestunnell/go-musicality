package note

import (
	"github.com/jamestunnell/go-musicality/notation/duration"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func Whole(pitches ...*pitch.Pitch) *Note {
	return New(duration.New(1, 1), pitches...)
}

func Half(pitches ...*pitch.Pitch) *Note {
	return New(duration.New(1, 2), pitches...)
}

func Quarter(pitches ...*pitch.Pitch) *Note {
	return New(duration.New(1, 4), pitches...)
}

func Eighth(pitches ...*pitch.Pitch) *Note {
	return New(duration.New(1, 8), pitches...)
}

func Sixteenth(pitches ...*pitch.Pitch) *Note {
	return New(duration.New(1, 16), pitches...)
}
