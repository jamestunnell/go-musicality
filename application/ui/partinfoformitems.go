package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-musicality/application"
)

type PartInfoFormComponents struct {
	NameEntry, MIDIChanEntry, MIDIInstrEntry *widget.Entry
	FormItems                                []*widget.FormItem
}

func NewPartInfoFormItems(pm *PartManager) *PartInfoFormComponents {
	nameEntry := widget.NewEntry()
	midiChanEntry := widget.NewEntry()
	midiInstrEntry := widget.NewEntry()

	nameEntry.Validator = func(s string) error {
		if len(s) == 0 {
			return fmt.Errorf("name is empty")
		}

		if pm.HasPart(s) {
			return fmt.Errorf("part '%s' already exists", s)
		}

		return nil
	}

	midiChanEntry.Validator = makeIntValidator(1, 16)
	midiInstrEntry.Validator = makeIntValidator(1, 128)

	nameItem := widget.NewFormItem("Name", nameEntry)
	midiChanItem := widget.NewFormItem("MIDI Channel", midiChanEntry)
	midiInstrItem := widget.NewFormItem("MIDI Instrument", midiInstrEntry)
	return &PartInfoFormComponents{
		NameEntry:      nameEntry,
		MIDIChanEntry:  midiChanEntry,
		MIDIInstrEntry: midiInstrEntry,
		FormItems:      []*widget.FormItem{nameItem, midiChanItem, midiInstrItem},
	}
}

func (pifc *PartInfoFormComponents) GetPartInfo() *application.PartInfo {
	c, err := strconv.Atoi(pifc.MIDIChanEntry.Text)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to convert MIDI channel text '%s' to int", pifc.MIDIChanEntry.Text)
	}

	i, err := strconv.Atoi(pifc.MIDIInstrEntry.Text)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to convert MIDI instrument text '%s' to int", pifc.MIDIInstrEntry.Text)
	}

	return &application.PartInfo{
		Name:           pifc.NameEntry.Text,
		MIDIChannel:    c,
		MIDIInstrument: i,
	}
}

func makeIntValidator(min, max int) fyne.StringValidator {
	return func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		if i < min || i > max {
			return fmt.Errorf("%d is not in range %d-%d", i, min, max)
		}

		return nil
	}
}
