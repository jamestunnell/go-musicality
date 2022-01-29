package note

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/xeipuuv/gojsonschema"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/validation"
)

type Note struct {
	Pitches            *pitch.Set
	Duration           *big.Rat
	Attack, Separation float64
	Links              map[*pitch.Pitch]*Link
}

type noteJSON struct {
	Pitches    []string         `json:"pitches,omitempty"`
	Duration   *big.Rat         `json:"duration"`
	Attack     float64          `json:"attack,omitempty"`
	Separation float64          `json:"separation,omitempty"`
	Links      map[string]*Link `json: "links,omitempty"`
}

const (
	AttackMin    = -1.0
	AttackNormal = 0.0
	AttackMax    = 1.0

	SeparationMin    = -1.0
	SeparationNormal = 0.0
	SeparationMax    = 1.0
)

func New(dur *big.Rat, pitches ...*pitch.Pitch) *Note {
	return &Note{
		Pitches:    pitch.NewSet(pitches...),
		Duration:   dur,
		Attack:     AttackNormal,
		Separation: SeparationNormal,
		Links:      make(map[*pitch.Pitch]*Link),
	}
}

func (n *Note) Validate() *validation.Result {
	errs := []error{}

	if err := validation.VerifyPositiveRat("duration", n.Duration); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyInRangeFloat("attack", n.Attack, AttackMin, AttackMax); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyInRangeFloat("separation", n.Separation, SeparationMin, SeparationMax); err != nil {
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

func ValidateJSON(documentLoader gojsonschema.JSONLoader) error {
	schemaLoader := gojsonschema.NewStringLoader(schema)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("failed to validate JSON: %w", err)
	}

	if !result.Valid() {
		errs := strings.Builder{}
		for _, desc := range result.Errors() {
			errs.WriteRune('\n')
			errs.WriteString(desc.String())
		}

		return fmt.Errorf("invalid note JSON: %s", errs.String())
	}

	return nil
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

	return json.Marshal(j)
}

func (n *Note) UnmarshalJSON(d []byte) error {
	var j noteJSON

	err := json.Unmarshal(d, &j)
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
