package note

import "github.com/jamestunnell/go-musicality/notation/pitch"

const (
	Tie         = "tie"
	Slur        = "slur"
	Glide       = "glide"
	Step        = "step"
	StepSlurred = "stepSlurred"
)

type Link struct {
	Target *pitch.Pitch `json:"target"`
	Type   string       `json:"type"`
}
