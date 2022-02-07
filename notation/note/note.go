package note

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/validation"
)

type Note struct {
	Pitches            *pitch.Set
	Duration           rat.Rat
	Attack, Separation float64
	Links              map[*pitch.Pitch]*Link
}

type noteJSON struct {
	Pitches    []string         `json:"pitches,omitempty"`
	Duration   rat.Rat          `json:"duration"`
	Attack     float64          `json:"attack"`
	Separation float64          `json:"separation"`
	Links      map[string]*Link `json: "links,omitempty"`
}

const (
	ControlMin    = 0.0
	ControlNormal = 0.5
	ControlMax    = 1.0
)

func New(dur rat.Rat, pitches ...*pitch.Pitch) *Note {
	return &Note{
		Pitches:    pitch.NewSet(pitches...),
		Duration:   dur,
		Attack:     ControlNormal,
		Separation: ControlNormal,
		Links:      make(map[*pitch.Pitch]*Link),
	}
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
	links := map[string]*Link{}

	for p, link := range n.Links {
		links[p.String()] = link
	}

	j := noteJSON{
		Pitches:    n.Pitches.Pitches().Strings(),
		Duration:   n.Duration,
		Attack:     n.Attack,
		Separation: n.Separation,
		Links:      links,
	}

	d, err := json.Marshal(j)
	if err != nil {
		return []byte{}, err
	}

	if n.Attack == ControlNormal {
		d, err = sjson.DeleteBytes(d, "attack")
		if err != nil {
			err = fmt.Errorf("failed to remove normal attack from JSON: %w", err)

			return []byte{}, err
		}
	}

	if n.Separation == ControlNormal {
		d, err = sjson.DeleteBytes(d, "separation")
		if err != nil {
			err = fmt.Errorf("failed to remove normal separation from JSON: %w", err)

			return []byte{}, err
		}
	}

	return d, nil
}

func (n *Note) UnmarshalJSON(d []byte) error {
	var err error

	result := gjson.GetBytes(d, "attack")
	if !result.Exists() {
		d, err = sjson.SetBytes(d, "attack", ControlNormal)
		if err != nil {
			return fmt.Errorf("failed to add normal attack to JSON: %w", err)
		}
	}

	result = gjson.GetBytes(d, "separation")
	if !result.Exists() {
		d, err = sjson.SetBytes(d, "separation", ControlNormal)
		if err != nil {
			return fmt.Errorf("failed to add normal separation to JSON: %w", err)
		}
	}

	var j noteJSON

	err = json.Unmarshal(d, &j)
	if err != nil {
		return err
	}

	links := map[*pitch.Pitch]*Link{}
	for pStr, link := range j.Links {
		p, err := pitch.Parse(pStr)
		if err != nil {
			return fmt.Errorf("failed to parse pitch '%s': %w", pStr, err)
		}

		links[p] = link
	}

	pitches := pitch.NewSet()
	for _, pStr := range j.Pitches {
		p, err := pitch.Parse(pStr)
		if err != nil {
			return fmt.Errorf("failed to parse pitch '%s': %w", pStr, err)
		}

		pitches.Add(p)
	}

	n.Pitches = pitches
	n.Duration = j.Duration
	n.Attack = j.Attack
	n.Separation = j.Separation
	n.Links = links

	return nil
}
