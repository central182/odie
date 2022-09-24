// Code generated by MockGen. DO NOT EDIT.
// Source: application.go

// Package application_mock is a generated GoMock package.
package application_mock

import (
	reflect "reflect"

	application "github.com/central182/odie/internal/domain/application"
	entry "github.com/central182/odie/internal/domain/dictionary/entry"
	headword "github.com/central182/odie/internal/domain/dictionary/headword"
	gomock "github.com/golang/mock/gomock"
)

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// GetEntriesByHeadword mocks base method.
func (m *MockApplication) GetEntriesByHeadword(arg0 headword.Headword) ([]entry.Entry, application.GetEntriesByHeadwordError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntriesByHeadword", arg0)
	ret0, _ := ret[0].([]entry.Entry)
	ret1, _ := ret[1].(application.GetEntriesByHeadwordError)
	return ret0, ret1
}

// GetEntriesByHeadword indicates an expected call of GetEntriesByHeadword.
func (mr *MockApplicationMockRecorder) GetEntriesByHeadword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntriesByHeadword", reflect.TypeOf((*MockApplication)(nil).GetEntriesByHeadword), arg0)
}
