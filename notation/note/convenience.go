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
