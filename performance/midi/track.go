package midi

type Track struct {
	Name                string
	Channel, Instrument uint8
	Events              NoteEvents
}

const DefaultInstrument = uint8(0)
