package rhythms_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/generation/rhythms"
	"github.com/jamestunnell/go-musicality/generation/temperley"
	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
	"github.com/jamestunnell/go-musicality/rendering/midi"
)

func TestRandomMeasure(t *testing.T) {
	rand.Seed(time.Now().Unix())

	met := meter.SixEight()
	smallestDur := rat.New(1, 8)

	s := score.New()

	s.Program = []string{"random"}

	sec := section.New(section.OptStartMeter(met))

	s.Sections["random"] = sec

	for i := 0; i < 24; i++ {
		m := measure.New()

		m.PartNotes["piano"] = randomMeasure(met, smallestDur)

		sec.Measures = append(sec.Measures, m)
	}

	assert.NoError(t, midi.WriteSMF(s, "./test.mid"))
}

func randomMeasure(met *meter.Meter, smallestDur rat.Rat) note.Notes {
	elems := rhythms.RandomMeasure(met, smallestDur)

	// t.Log(strings.Join(elems.Strings(), " "))

	n := len(elems)
	pitchModel, _ := temperley.NewMajorPitchModel(0, uint64(time.Now().Unix()))

	// require.NoError(t, err)

	pitches := pitchModel.MakePitches(uint(n))

	// t.Log(strings.Join(pitches.Strings(), " "))

	notes := make(note.Notes, n)

	for i := 0; i < len(elems); i++ {
		notes[i] = note.New(elems[i].Duration, pitches[i])
	}

	return notes
}
