package mononote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/centpitch"
	"github.com/jamestunnell/go-musicality/performance/mononote"
)

func TestNoteConverterEmpty(t *testing.T) {
	nc := mononote.NewConverter()
	notes := []*note.Note{}

	notes2, err := nc.Process(notes)

	assert.NoError(t, err)
	assert.Empty(t, notes2)
}

func TestNoteConverterOpts(t *testing.T) {
	nc := mononote.NewConverter()

	assert.Equal(t, mononote.DefaultCentsPerStep, nc.CentsPerStep())
	assert.False(t, nc.ReplaceSlursAndGlides())

	nc = mononote.NewConverter(
		mononote.OptionCentsPerStep(0),
		mononote.OptionReplaceSlursAndGlides())

	assert.Equal(t, 1, nc.CentsPerStep())
	assert.True(t, nc.ReplaceSlursAndGlides())

	nc = mononote.NewConverter(mononote.OptionCentsPerStep(-1))

	assert.Equal(t, 1, nc.CentsPerStep())

	nc = mononote.NewConverter(mononote.OptionCentsPerStep(101))

	assert.Equal(t, 100, nc.CentsPerStep())

	nc = mononote.NewConverter(mononote.OptionCentsPerStep(50))

	assert.Equal(t, 50, nc.CentsPerStep())
}

func TestNoteConverterMonophonicNote(t *testing.T) {
	nc := mononote.NewConverter()
	notes := []*note.Note{note.Sixteenth(pitch.D4)}

	notes2, err := nc.Process(notes)

	require.NoError(t, err)
	require.Len(t, notes2, 1)
	require.Len(t, notes2[0].PitchDurs, 1)

	assert.Equal(t, note.ControlNormal, notes2[0].Attack)
	assert.Equal(t, note.ControlNormal, notes2[0].Separation)
	assert.True(t, notes2[0].Start.Equal(rat.Zero()))
	verifyPD(t, notes2[0].PitchDurs[0], centpitch.New(pitch.D4, 0), rat.New(1, 16))
}

func TestNoteConverterSingleChord(t *testing.T) {
	nc := mononote.NewConverter()
	notes := []*note.Note{note.Half(pitch.G3, pitch.B3)}

	notes2, err := nc.Process(notes)

	require.NoError(t, err)
	require.Len(t, notes2, 2)

	assert.Len(t, notes2[0].PitchDurs, 1)
	assert.True(t, notes2[0].Duration().Equal(rat.New(1, 2)))
	assert.True(t, notes2[0].Start.Equal(rat.Zero()))

	assert.Len(t, notes2[1].PitchDurs, 1)
	assert.True(t, notes2[1].Duration().Equal(rat.New(1, 2)))
	assert.True(t, notes2[1].Start.Equal(rat.Zero()))
}

func TestNoteConverterNoteSimplifies(t *testing.T) {
	nc := mononote.NewConverter()
	notes := []*note.Note{
		note.Quarter(),
		note.Quarter(pitch.G3).Tie(pitch.G3),
		note.Quarter(pitch.G3).Slur(pitch.G3, pitch.G3),
		note.Quarter(pitch.G3),
		note.Quarter(),
	}

	notes2, err := nc.Process(notes)

	require.NoError(t, err)
	require.Len(t, notes2, 1)
	require.Len(t, notes2[0].PitchDurs, 1)

	assert.True(t, notes2[0].Start.Equal(rat.New(1, 4)))
	verifyPD(t, notes2[0].PitchDurs[0], centpitch.New(pitch.G3, 0), rat.New(3, 4))
}

func TestNoteConverterUnknownLinkType(t *testing.T) {
	nc := mononote.NewConverter()
	l := &note.Link{Type: "unknown", Source: pitch.G3, Target: pitch.G3}
	notes := []*note.Note{
		note.Quarter(pitch.G3).Link(l),
		note.Quarter(pitch.G3),
	}

	_, err := nc.Process(notes)

	assert.Error(t, err)
}

func TestNoteConverterIgnoresMissingLinks(t *testing.T) {
	nc := mononote.NewConverter()
	notes := []*note.Note{
		note.Quarter(pitch.G3).Tie(pitch.A3),
		note.Quarter(pitch.G3).Slur(pitch.F3, pitch.G3),
		note.Quarter(pitch.G3).Slur(pitch.G3, pitch.A3),
		note.Quarter(pitch.G3),
	}

	notes2, err := nc.Process(notes)

	require.NoError(t, err)
	require.Len(t, notes2, 4)
	require.Len(t, notes2[0].PitchDurs, 1)
	require.Len(t, notes2[1].PitchDurs, 1)
	require.Len(t, notes2[2].PitchDurs, 1)
	require.Len(t, notes2[3].PitchDurs, 1)

	assert.True(t, notes2[0].Start.Equal(rat.Zero()))
	assert.True(t, notes2[1].Start.Equal(rat.New(1, 4)))
	assert.True(t, notes2[2].Start.Equal(rat.New(1, 2)))
	assert.True(t, notes2[3].Start.Equal(rat.New(3, 4)))

	verifyPD(t, notes2[0].PitchDurs[0], centpitch.New(pitch.G3, 0), rat.New(1, 4))
	verifyPD(t, notes2[1].PitchDurs[0], centpitch.New(pitch.G3, 0), rat.New(1, 4))
	verifyPD(t, notes2[2].PitchDurs[0], centpitch.New(pitch.G3, 0), rat.New(1, 4))
	verifyPD(t, notes2[3].PitchDurs[0], centpitch.New(pitch.G3, 0), rat.New(1, 4))
}

func TestNoteConverterSlursAllowedByDefault(t *testing.T) {
	nc := mononote.NewConverter()
	notes := []*note.Note{
		note.Eighth(pitch.G3).Slur(pitch.G3, pitch.A3),
		note.Eighth(pitch.A3),
	}

	notes2, err := nc.Process(notes)

	require.NoError(t, err)
	require.Len(t, notes2, 1)
	require.Len(t, notes2[0].PitchDurs, 2)

	verifyPD(t, notes2[0].PitchDurs[0], centpitch.New(pitch.G3, 0), rat.New(1, 8))
	verifyPD(t, notes2[0].PitchDurs[1], centpitch.New(pitch.A3, 0), rat.New(1, 8))
}

func TestNoteConverterGlidesAllowedByDefault(t *testing.T) {
	nc := mononote.NewConverter(mononote.OptionCentsPerStep(25))
	notes := []*note.Note{
		note.Whole(pitch.C5).Glide(pitch.C5, pitch.Db5),
		note.Half(pitch.Db5),
	}

	notes2, err := nc.Process(notes)

	require.NoError(t, err)
	require.Len(t, notes2, 1)
	require.Len(t, notes2[0].PitchDurs, 5)

	verifyPD(t, notes2[0].PitchDurs[0], centpitch.New(pitch.C5, 0), rat.New(1, 4))
	verifyPD(t, notes2[0].PitchDurs[1], centpitch.New(pitch.C5, 25), rat.New(1, 4))
	verifyPD(t, notes2[0].PitchDurs[2], centpitch.New(pitch.C5, 50), rat.New(1, 4))
	verifyPD(t, notes2[0].PitchDurs[3], centpitch.New(pitch.C5, 75), rat.New(1, 4))
	verifyPD(t, notes2[0].PitchDurs[4], centpitch.New(pitch.C5, 100), rat.New(1, 2))
}

func TestNoteConverterReplaceGlides(t *testing.T) {
	nc := mononote.NewConverter(mononote.OptionReplaceSlursAndGlides())
	notes := []*note.Note{
		note.Whole(pitch.C5).Glide(pitch.C5, pitch.D5),
		note.Half(pitch.D5),
	}

	notes2, err := nc.Process(notes)

	require.NoError(t, err)
	require.Len(t, notes2, 3)
	require.Len(t, notes2[0].PitchDurs, 1)
	require.Len(t, notes2[1].PitchDurs, 1)
	require.Len(t, notes2[2].PitchDurs, 1)

	verifyPD(t, notes2[0].PitchDurs[0], centpitch.New(pitch.C5, 0), rat.New(1, 2))
	verifyPD(t, notes2[1].PitchDurs[0], centpitch.New(pitch.Db5, 0), rat.New(1, 2))
	verifyPD(t, notes2[2].PitchDurs[0], centpitch.New(pitch.D5, 0), rat.New(1, 2))
}

func TestNoteConverterReplaceSlur(t *testing.T) {
	nc := mononote.NewConverter(mononote.OptionReplaceSlursAndGlides())
	notes := []*note.Note{
		note.Quarter(pitch.C5).Slur(pitch.C5, pitch.G5),
		note.Eighth(pitch.G5),
	}

	notes2, err := nc.Process(notes)

	require.NoError(t, err)
	require.Len(t, notes2, 2)
	require.Len(t, notes2[0].PitchDurs, 1)
	require.Len(t, notes2[1].PitchDurs, 1)

	verifyPD(t, notes2[0].PitchDurs[0], centpitch.New(pitch.C5, 0), rat.New(1, 4))
	verifyPD(t, notes2[1].PitchDurs[0], centpitch.New(pitch.G5, 0), rat.New(1, 8))
}

func verifyPD(t *testing.T, pd *mononote.PitchDur, expectedPitch *centpitch.CentPitch, expectedDur rat.Rat) {
	assert.True(t, pd.Duration.Equal(expectedDur))
	assert.True(t, pd.Pitch.Equal(expectedPitch))
}
