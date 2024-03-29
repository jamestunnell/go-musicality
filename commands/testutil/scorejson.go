package testutil

import (
	"bytes"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/jamestunnell/go-musicality/notation/key"
	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
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
	d := ScoreJSON(t, s)

	return bytes.Replace(d, []byte("sections"), []byte("not-sections"), 1)
}

func ValidScoreJSON(t *testing.T) []byte {
	return ScoreJSON(t, ValidScore())
}

func ValidScore() *score.Score {
	s := score.New()
	sec := section.New()
	m := measure.New()

	m.PartNotes[TestPartName] = []*note.Note{
		note.New(big.NewRat(1, 1), pitch.C4),
	}

	m.MeterChange = meter.FourFour()
	m.KeyChange = key.NewMajor("C")

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
