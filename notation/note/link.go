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

type Links map[*pitch.Pitch]*Link

func (link *Link) Equal(other *Link) bool {
	return link.Type == other.Type && link.Target.Equal(other.Target)
}

func (links Links) Equal(other Links) bool {
	if len(links) != len(other) {
		return false
	}

	for p, link := range links {
		link2, found := other[p]
		if !found {
			return false
		}

		if !link.Equal(link2) {
			return false
		}
	}

	return true
}
