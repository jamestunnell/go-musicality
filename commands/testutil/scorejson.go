package testutil

import (
	"encoding/json"
	"testing"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
	"github.com/stretchr/testify/require"
)

const (
	TestPartName    = "testPart"
	TestSectionName = "testSection"
)

func InvalidScoreJSON(t *testing.T) []byte {
	s := ValidScore()

	// invalidate
	s.Sections[TestSectionName].Measures[0].PartNotes[TestPartName][0].Duration.Accum(rat.New(1, 4))

	return ScoreJSON(t, s)
}

func ValidScoreJSON(t *testing.T) []byte {
	return ScoreJSON(t, ValidScore())
}

func ValidScore() *score.Score {
	s := score.New()
	sec := section.New()
	m := measure.New(meter.New(4, 4))

	m.PartNotes[TestPartName] = []*note.Note{
		note.New(rat.New(1, 1), pitch.C4),
	}

	sec.Measures = append(sec.Measures, m)

	s.Sections[TestSectionName] = sec

	s.Program = append(s.Program, TestSectionName)

	return s
}

func ScoreJSON(t *testing.T, s *score.Score) []byte {
	d, err := json.Marshal(s)

	require.NoError(t, err)

	return d
}
