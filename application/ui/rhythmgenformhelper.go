package ui

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/meter"
)

var (
	meters = []*meter.Meter{
		meter.FourFour(),
		meter.ThreeFour(),
		meter.TwoFour(),
		meter.SixEight(),
		meter.TwoTwo(),
	}
)

type RhythmGenFormHelper struct {
	selectedMeter               *meter.Meter
	nameEntry, smallestDurEntry *widget.Entry
	meterSelect                 *widget.Select
	formItems                   []*widget.FormItem
}

func NewRhythmGenFormHelper(m *ItemManager) ItemFormHelper {
	meterStrings := make([]string, len(meters))
	for i := 0; i < len(meters); i++ {
		meterStrings[i] = meters[i].String()
	}

	h := &RhythmGenFormHelper{}

	h.nameEntry = widget.NewEntry()
	h.meterSelect = widget.NewSelect(meterStrings, func(s string) {
		for i := 0; i < len(meterStrings); i++ {
			if s == meterStrings[i] {
				h.selectedMeter = meters[i]
			}
		}
	})
	h.smallestDurEntry = widget.NewEntry()

	h.nameEntry.Validator = MakeNameValidator(m)
	h.smallestDurEntry.Validator = func(s string) error {
		r, ok := rat.FromString(s)
		if !ok {
			return fmt.Errorf("'%s' is not a valid rational number (a/b)", s)
		}

		if !r.Positive() {
			return fmt.Errorf("'%s' is not positive", s)
		}

		if h.selectedMeter == nil {
			return fmt.Errorf("no meter selected", s)
		}

		if r.Greater(h.selectedMeter.BeatDuration) {
			err := fmt.Errorf("'%s' is not less or equal to meter beat dur %s",
				s, h.selectedMeter.BeatDuration)

			return err
		}

		return nil
	}

	nameItem := widget.NewFormItem("Name", h.nameEntry)
	meterItem := widget.NewFormItem("Meter", h.meterSelect)
	smallestDurItem := widget.NewFormItem("Smallest Duration", h.smallestDurEntry)

	h.formItems = []*widget.FormItem{nameItem, meterItem, smallestDurItem}

	return h
}

func (h *RhythmGenFormHelper) FormItems() []*widget.FormItem {
	return h.formItems
}

func (h *RhythmGenFormHelper) MakeItem() Item {
	if h.selectedMeter == nil {
		log.Fatal().Msgf("no meter selected")
	}

	r, ok := rat.FromString(h.smallestDurEntry.Text)
	if !ok {
		log.Fatal().Msgf("failed to convert smallest dur text '%s' to rat.Rat", h.smallestDurEntry.Text)
	}

	return &RhythmGen{
		name:        h.nameEntry.Text,
		Meter:       h.selectedMeter,
		SmallestDur: r,
	}
}
