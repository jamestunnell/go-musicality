package note

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/validation"
)

type Note struct {
	Pitches            *pitch.Set
	Duration           *big.Rat
	Attack, Separation float64
	Links              Links
}

type noteJSON struct {
	Pitches  []string    `json:"pitches,omitempty"`
	Duration *big.Rat    `json:"duration"`
	LinkMap  linkLiteMap `json:"links,omitempty"`
	// Attack     float64          `json:"attack"`
	// Separation float64          `json:"separation"`
}

const (
	ControlMin    = 0.0
	ControlNormal = 0.5
	ControlMax    = 1.0
)

func New(dur *big.Rat, pitches ...*pitch.Pitch) *Note {
	return &Note{
		Pitches:    pitch.NewSet(pitches...),
		Duration:   dur,
		Attack:     ControlNormal,
		Separation: ControlNormal,
		Links:      Links{},
	}
}

func (n *Note) Equal(other *Note) bool {
	if !rat.IsEqual(n.Duration, other.Duration) {
		return false
	}

	if !n.Pitches.Equal(other.Pitches) {
		return false
	}

	if n.Attack != other.Attack {
		return false
	}

	if n.Separation != other.Separation {
		return false
	}

	if !n.Links.Equal(other.Links) {
		return false
	}

	return true
}

func (n *Note) Validate() *validation.Result {
	errs := []error{}

	if err := validation.VerifyPositiveRat("duration", n.Duration); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyInRangeFloat("attack", n.Attack, ControlMin, ControlMax); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyInRangeFloat("separation", n.Separation, ControlMin, ControlMax); err != nil {
		errs = append(errs, err)
	}

	if len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "note",
		Errors:     errs,
		SubResults: []*validation.Result{},
	}
}

func (n *Note) MarshalJSON() ([]byte, error) {
	j := noteJSON{
		Pitches:  n.Pitches.Pitches().Strings(),
		Duration: n.Duration,
		LinkMap:  n.Links.ToLinkLiteMap(),
	}

	d, err := json.Marshal(j)
	if err != nil {
		return []byte{}, err
	}

	if n.Attack != ControlNormal {
		d, err = sjson.SetBytes(d, "attack", n.Attack)
		if err != nil {
			err = fmt.Errorf("failed to add non-normal attack to JSON: %w", err)

			return []byte{}, err
		}
	}

	if n.Separation != ControlNormal {
		d, err = sjson.SetBytes(d, "separation", n.Separation)
		if err != nil {
			err = fmt.Errorf("failed to add non-normal separation to JSON: %w", err)

			return []byte{}, err
		}
	}

	return d, nil
}

func (n *Note) UnmarshalJSON(d []byte) error {
	var j noteJSON

	err := json.Unmarshal(d, &j)
	if err != nil {
		return err
	}

	links, err := j.LinkMap.ToLinks()
	if err != nil {
		return fmt.Errorf("failed to make links from map: %w", err)
	}

	pitches := pitch.NewSet()
	for _, pStr := range j.Pitches {
		p, err := pitch.Parse(pStr)
		if err != nil {
			return fmt.Errorf("failed to parse pitch '%s': %w", pStr, err)
		}

		pitches.Add(p)
	}

	attack := ControlNormal

	result := gjson.GetBytes(d, "attack")
	if result.Exists() {
		attack = result.Float()
	}

	separation := ControlNormal

	result = gjson.GetBytes(d, "separation")
	if result.Exists() {
		separation = result.Float()
	}

	n.Pitches = pitches
	n.Duration = j.Duration
	n.Attack = attack
	n.Separation = separation
	n.Links = links

	return nil
}
