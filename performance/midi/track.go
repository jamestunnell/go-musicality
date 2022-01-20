package midi

type Track struct {
	Name                string
	Channel, Instrument uint8
	Events              []*Event
}

const DefaultInstrument = uint8(0)
