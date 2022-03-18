package block

import (
	"reflect"
)

type Port struct {
	Type                     PortType
	StartValue, CurrentValue interface{}
	ValueType                reflect.Type
}

type PortType int

const (
	ControlPort PortType = iota
	InputPort
	OutputPort
)

func NewControl(startVal interface{}) *Port {
	return newPort(ControlPort, startVal)
}

func NewInput(startVal interface{}) *Port {
	return newPort(InputPort, startVal)
}

func NewOutput(startVal interface{}) *Port {
	return newPort(OutputPort, startVal)
}

func newPort(typ PortType, startVal interface{}) *Port {
	return &Port{
		Type:         typ,
		CurrentValue: startVal,
		StartValue:   startVal,
		ValueType:    reflect.TypeOf(startVal),
	}
}
