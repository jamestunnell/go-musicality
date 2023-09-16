package key

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/validation"
)

type Key struct {
	Tonic string `json:"tonic"`
	Mode  string `json:"mode"`
}

const (
	Major = "major"
	Minor = "minor"
)

func new(tonic, mode string) *Key {
	return &Key{Tonic: tonic, Mode: mode}
}

func NewMajor(tonic string) *Key {
	return new(tonic, Major)
}

func NewMinor(tonic string) *Key {
	return new(tonic, Minor)
}

func (k *Key) Validate() *validation.Result {
	errs := []error{}

	if _, err := pitch.ParseSemitone(k.Tonic); err != nil {
		err = fmt.Errorf("invalid tonic '%s': %w", k.Tonic, err)

		errs = append(errs, err)
	}

	switch k.Mode {
	case Major, Minor:
		// do nothing
	default:
		err := fmt.Errorf("unknown mode '%s'", k.Mode)

		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return &validation.Result{
			Context:    "key",
			Errors:     errs,
			SubResults: []*validation.Result{},
		}
	}

	return nil
}
