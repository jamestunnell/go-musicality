package note

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/validation"
)

type Note struct {
	Pitches      *pitch.Set
	Duration     *big.Rat
	Articulation string
	Slurs        bool
	Links        map[*pitch.Pitch]*Link
}

type noteJSON struct {
	Pitches      []string         `json:"pitches,omitempty"`
	Duration     *big.Rat         `json:"duration"`
	Articulation string           `json:"articulation,omitempty"`
	Slurs        bool             `json:"slurs,omitempty"`
	Links        map[string]*Link `json: "links,omitempty"`
}

const (
	Normal        = ""
	Legato        = "legato"
	Tenuto        = "tenuto"
	Accent        = "accent"
	Marcato       = "marcato"
	Portato       = "portato"
	Staccato      = "staccato"
	Staccatissimo = "staccatissimo"
)

func New(dur *big.Rat, pitches ...*pitch.Pitch) *Note {
	return &Note{
		Pitches:      pitch.NewSet(pitches...),
		Duration:     dur,
		Articulation: Normal,
		Slurs:        false,
		Links:        make(map[*pitch.Pitch]*Link),
	}
}

func (n *Note) Validate() *validation.Result {
	errs := []error{}

	if n.Duration.Cmp(big.NewRat(0, 1)) != 1 {
		errs = append(errs, validation.NewErrNonPositiveRat("duration", n.Duration))
	}

	switch n.Articulation {
	case "", Legato, Tenuto, Accent, Marcato, Portato, Staccato, Staccatissimo:
		// do nothing
	default:
		errs = append(errs, fmt.Errorf("unknown articulation %s", n.Articulation))
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

func (n *Note) Dot() {
	n.Duration = new(big.Rat).Mul(n.Duration, big.NewRat(3, 2))
}

func (n *Note) IsRest() bool {
	return n.Pitches.Len() == 0
}

func (n *Note) IsMonophonic() bool {
	return n.Pitches.Len() == 1
}

func (n *Note) MarshalJSON() ([]byte, error) {
	links := map[string]*Link{}

	for p, link := range n.Links {
		links[p.String()] = link
	}

	j := noteJSON{
		Pitches:      n.Pitches.Pitches().Strings(),
		Duration:     n.Duration,
		Articulation: n.Articulation,
		Slurs:        n.Slurs,
		Links:        links,
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

	n.Links = links
	n.Pitches = pitches
	n.Duration = j.Duration
	n.Articulation = j.Articulation
	n.Slurs = j.Slurs
	n.Links = links

	return nil
}
