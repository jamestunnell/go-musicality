package note

import (
	"encoding/json"

	"github.com/jamestunnell/go-musicality/notation/duration"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/pkg/util"
)

type Note struct {
	Pitches    []*pitch.Pitch       `json:"pitches,omitempty"`
	Duration   *duration.Duration   `json:"duration"`
	Attributes map[string]Attribute `json:"attributes,omitempty"`
}

func New(dur *duration.Duration, pitches ...*pitch.Pitch) *Note {
	pitches2 := make([]*pitch.Pitch, len(pitches))

	for i, p := range pitches {
		pitches2[i] = p.Clone()
	}

	return &Note{Pitches: pitches2, Duration: dur.Clone()}
}

func (n *Note) Validate() error {
	if !n.Duration.Positive() {
		return util.NewNonPositiveDurationError(n.Duration)
	}

	return nil
}

func (n *Note) Clone() *Note {
	return New(n.Duration, n.Pitches...)
}

func (n *Note) JSON() string {
	d, _ := json.Marshal(n)

	return string(d)
}

func (n *Note) Dot() *Note {
	n2 := n.Clone()

	n2.Duration = n.Duration.Mul(duration.New(3, 2))

	return n2
}

func (n *Note) IsRest() bool {
	return len(n.Pitches) == 0
}

func (n *Note) IsMonophonic() bool {
	return len(n.Pitches) == 1
}
