package sequence

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/duration"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

type Element struct {
	Duration   *duration.Duration
	Pitch      *pitch.Pitch
	Attack     float32
	Separation float32
}

func (e *Element) Validate() error {
	if !e.Duration.Positive() {
		return fmt.Errorf("duration %v is non-positive", e.Duration)
	}

	return nil
}
