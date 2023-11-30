// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package httpc is a generated GoMock package.
package httpc

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpClientRepo is a mock of HttpClientRepo interface.
type MockHttpClientRepo struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientRepoMockRecorder
}

// MockHttpClientRepoMockRecorder is the mock recorder for MockHttpClientRepo.
type MockHttpClientRepoMockRecorder struct {
	mock *MockHttpClientRepo
}

// NewMockHttpClientRepo creates a new mock instance.
func NewMockHttpClientRepo(ctrl *gomock.Controller) *MockHttpClientRepo {
	mock := &MockHttpClientRepo{ctrl: ctrl}
	mock.recorder = &MockHttpClientRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpClientRepo) EXPECT() *MockHttpClientRepoMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockHttpClientRepo) Delete(addrs string, payload []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", addrs, payload)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockHttpClientRepoMockRecorder) Delete(addrs, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHttpClientRepo)(nil).Delete), addrs, payload)
}

// DeleteFormValue mocks base method.
func (m *MockHttpClientRepo) DeleteFormValue(method, key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteFormValue", method, key)
}

// DeleteFormValue indicates an expected call of DeleteFormValue.
func (mr *MockHttpClientRepoMockRecorder) DeleteFormValue(method, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFormValue", reflect.TypeOf((*MockHttpClientRepo)(nil).DeleteFormValue), method, key)
}

// DeleteHeader mocks base method.
func (m *MockHttpClientRepo) DeleteHeader(method, key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteHeader", method, key)
}

// DeleteHeader indicates an expected call of DeleteHeader.
func (mr *MockHttpClientRepoMockRecorder) DeleteHeader(method, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHeader", reflect.TypeOf((*MockHttpClientRepo)(nil).DeleteHeader), method, key)
}

// Get mocks base method.
func (m *MockHttpClientRepo) Get(addrs string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", addrs)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHttpClientRepoMockRecorder) Get(addrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClientRepo)(nil).Get), addrs)
}

// GetBasicAuth mocks base method.
func (m *MockHttpClientRepo) GetBasicAuth(method string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBasicAuth", method)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetBasicAuth indicates an expected call of GetBasicAuth.
func (mr *MockHttpClientRepoMockRecorder) GetBasicAuth(method interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBasicAuth", reflect.TypeOf((*MockHttpClientRepo)(nil).GetBasicAuth), method)
}

// GetFormValue mocks base method.
func (m *MockHttpClientRepo) GetFormValue(method string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFormValue", method)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetFormValue indicates an expected call of GetFormValue.
func (mr *MockHttpClientRepoMockRecorder) GetFormValue(method interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFormValue", reflect.TypeOf((*MockHttpClientRepo)(nil).GetFormValue), method)
}

// GetHeaders mocks base method.
func (m *MockHttpClientRepo) GetHeaders(method string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaders", method)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetHeaders indicates an expected call of GetHeaders.
func (mr *MockHttpClientRepoMockRecorder) GetHeaders(method interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaders", reflect.TypeOf((*MockHttpClientRepo)(nil).GetHeaders), method)
}

// Head mocks base method.
func (m *MockHttpClientRepo) Head(addrs string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head", addrs)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockHttpClientRepoMockRecorder) Head(addrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockHttpClientRepo)(nil).Head), addrs)
}

// Patch mocks base method.
func (m *MockHttpClientRepo) Patch(addrs string, payload []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", addrs, payload)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockHttpClientRepoMockRecorder) Patch(addrs, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockHttpClientRepo)(nil).Patch), addrs, payload)
}

// Post mocks base method.
func (m *MockHttpClientRepo) Post(addrs string, payload []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", addrs, payload)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockHttpClientRepoMockRecorder) Post(addrs, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHttpClientRepo)(nil).Post), addrs, payload)
}

// Put mocks base method.
func (m *MockHttpClientRepo) Put(addrs string, payload []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", addrs, payload)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockHttpClientRepoMockRecorder) Put(addrs, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockHttpClientRepo)(nil).Put), addrs, payload)
}

// SetBasicAuth mocks base method.
func (m *MockHttpClientRepo) SetBasicAuth(method, username, password string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBasicAuth", method, username, password)
}

// SetBasicAuth indicates an expected call of SetBasicAuth.
func (mr *MockHttpClientRepoMockRecorder) SetBasicAuth(method, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBasicAuth", reflect.TypeOf((*MockHttpClientRepo)(nil).SetBasicAuth), method, username, password)
}

// SetFormValue mocks base method.
func (m *MockHttpClientRepo) SetFormValue(method, key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFormValue", method, key, value)
}

// SetFormValue indicates an expected call of SetFormValue.
func (mr *MockHttpClientRepoMockRecorder) SetFormValue(method, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFormValue", reflect.TypeOf((*MockHttpClientRepo)(nil).SetFormValue), method, key, value)
}

// SetHeader mocks base method.
func (m *MockHttpClientRepo) SetHeader(method, key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHeader", method, key, value)
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockHttpClientRepoMockRecorder) SetHeader(method, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockHttpClientRepo)(nil).SetHeader), method, key, value)
}

// SetPatchHeader mocks base method.
func (m *MockHttpClientRepo) SetPatchHeader(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPatchHeader", key, value)
}

// SetPatchHeader indicates an expected call of SetPatchHeader.
func (mr *MockHttpClientRepoMockRecorder) SetPatchHeader(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPatchHeader", reflect.TypeOf((*MockHttpClientRepo)(nil).SetPatchHeader), key, value)
}