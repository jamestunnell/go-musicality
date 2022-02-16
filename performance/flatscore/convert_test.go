package flatscore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
	"github.com/jamestunnell/go-musicality/performance/flatscore"
)

func TestConverterEmpty(t *testing.T) {
	fs, err := flatscore.Convert(score.New())

	assert.NoError(t, err)
	assert.Empty(t, fs.Parts)
	assert.True(t, fs.Duration().Equal(rat.Zero()))
}

func TestConvertInvalidScore(t *testing.T) {
	s := score.New()

	s.Program = append(s.Program, "missing")

	_, err := flatscore.Convert(s)

	assert.Error(t, err)
}

func TestConvert(t *testing.T) {
	m1 := measure.New()
	m2 := measure.New()
	m3 := measure.New()
	m4 := measure.New()
	sec1 := section.New()
	sec2 := section.New()
	s := score.New()

	m1.PartNotes["piano"] = note.Notes{
		note.Whole(pitch.C3, pitch.Bb3),
	}
	m2.PartNotes["piano"] = note.Notes{
		note.Whole(pitch.C3, pitch.Bb3),
	}
	m3.PartNotes["bass"] = note.Notes{
		note.Whole(pitch.G2),
	}
	m4.PartNotes["bass"] = note.Notes{
		note.Whole(pitch.G2),
	}

	sec1.Measures = append(sec1.Measures, m1)
	sec1.Measures = append(sec1.Measures, m2)

	sec2.Measures = append(sec2.Measures, m3)
	sec2.Measures = append(sec2.Measures, m4)

	s.Program = append(s.Program, "section1", "section2")

	s.Sections["section1"] = sec1
	s.Sections["section2"] = sec2

	fs, err := flatscore.Convert(s)

	assert.NoError(t, err)
	assert.Len(t, fs.Parts, 2)
	assert.True(t, fs.Duration().Equal(rat.New(4, 1)))
}
