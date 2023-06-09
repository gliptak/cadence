// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package isolationgroupapi is a generated GoMock package.
package isolationgroupapi

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/uber/cadence/common/types"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// GetDomainState mocks base method.
func (m *MockHandler) GetDomainState(ctx context.Context, request types.GetDomainIsolationGroupsRequest) (*types.GetDomainIsolationGroupsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainState", ctx, request)
	ret0, _ := ret[0].(*types.GetDomainIsolationGroupsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainState indicates an expected call of GetDomainState.
func (mr *MockHandlerMockRecorder) GetDomainState(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainState", reflect.TypeOf((*MockHandler)(nil).GetDomainState), ctx, request)
}

// GetGlobalState mocks base method.
func (m *MockHandler) GetGlobalState(ctx context.Context) (*types.GetGlobalIsolationGroupsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGlobalState", ctx)
	ret0, _ := ret[0].(*types.GetGlobalIsolationGroupsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGlobalState indicates an expected call of GetGlobalState.
func (mr *MockHandlerMockRecorder) GetGlobalState(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlobalState", reflect.TypeOf((*MockHandler)(nil).GetGlobalState), ctx)
}

// UpdateDomainState mocks base method.
func (m *MockHandler) UpdateDomainState(ctx context.Context, state types.UpdateDomainIsolationGroupsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDomainState", ctx, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDomainState indicates an expected call of UpdateDomainState.
func (mr *MockHandlerMockRecorder) UpdateDomainState(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDomainState", reflect.TypeOf((*MockHandler)(nil).UpdateDomainState), ctx, state)
}

// UpdateGlobalState mocks base method.
func (m *MockHandler) UpdateGlobalState(ctx context.Context, state types.UpdateGlobalIsolationGroupsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGlobalState", ctx, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGlobalState indicates an expected call of UpdateGlobalState.
func (mr *MockHandlerMockRecorder) UpdateGlobalState(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGlobalState", reflect.TypeOf((*MockHandler)(nil).UpdateGlobalState), ctx, state)
}