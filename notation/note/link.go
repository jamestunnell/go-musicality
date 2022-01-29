package note

import "github.com/jamestunnell/go-musicality/notation/pitch"

const (
	Tie   = "tie"
	Glide = "glide"
	Step  = "step"
)

type Link struct {
	Target *pitch.Pitch
	Type   string
}
