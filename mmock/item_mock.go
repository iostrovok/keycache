// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iostrovok/keycache (interfaces: IItem)

// Package mmock is a generated GoMock package.
package mmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIItem is a mock of IItem interface.
type MockIItem struct {
	ctrl     *gomock.Controller
	recorder *MockIItemMockRecorder
}

// MockIItemMockRecorder is the mock recorder for MockIItem.
type MockIItemMockRecorder struct {
	mock *MockIItem
}

// NewMockIItem creates a new mock instance.
func NewMockIItem(ctrl *gomock.Controller) *MockIItem {
	mock := &MockIItem{ctrl: ctrl}
	mock.recorder = &MockIItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIItem) EXPECT() *MockIItemMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockIItem) Decode(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockIItemMockRecorder) Decode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockIItem)(nil).Decode), arg0)
}

// Encode mocks base method.
func (m *MockIItem) Encode() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockIItemMockRecorder) Encode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockIItem)(nil).Encode))
}

// ID mocks base method.
func (m *MockIItem) ID() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(int)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockIItemMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockIItem)(nil).ID))
}

// Sign mocks base method.
func (m *MockIItem) Sign() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Sign indicates an expected call of Sign.
func (mr *MockIItemMockRecorder) Sign() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockIItem)(nil).Sign))
}
