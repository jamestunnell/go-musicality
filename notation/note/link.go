package note

import "github.com/jamestunnell/go-musicality/notation/pitch"

const (
	Tie        = "tie"
	Portamento = "portamento"
	Glissando  = "glissando"
)

type Link struct {
	Target *pitch.Pitch
	Type   string
}
