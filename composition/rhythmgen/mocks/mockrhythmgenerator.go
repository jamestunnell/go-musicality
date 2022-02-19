// Code generated by MockGen. DO NOT EDIT.
// Source: rhythmgenerator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	rat "github.com/jamestunnell/go-musicality/common/rat"
)

// MockRhythmGenerator is a mock of RhythmGenerator interface.
type MockRhythmGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockRhythmGeneratorMockRecorder
}

// MockRhythmGeneratorMockRecorder is the mock recorder for MockRhythmGenerator.
type MockRhythmGeneratorMockRecorder struct {
	mock *MockRhythmGenerator
}

// NewMockRhythmGenerator creates a new mock instance.
func NewMockRhythmGenerator(ctrl *gomock.Controller) *MockRhythmGenerator {
	mock := &MockRhythmGenerator{ctrl: ctrl}
	mock.recorder = &MockRhythmGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRhythmGenerator) EXPECT() *MockRhythmGeneratorMockRecorder {
	return m.recorder
}

// NextDur mocks base method.
func (m *MockRhythmGenerator) NextDur() rat.Rat {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextDur")
	ret0, _ := ret[0].(rat.Rat)
	return ret0
}

// NextDur indicates an expected call of NextDur.
func (mr *MockRhythmGeneratorMockRecorder) NextDur() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextDur", reflect.TypeOf((*MockRhythmGenerator)(nil).NextDur))
}

// Reset mocks base method.
func (m *MockRhythmGenerator) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockRhythmGeneratorMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockRhythmGenerator)(nil).Reset))
}
