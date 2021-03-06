// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_selector is a generated GoMock package.
package mock_selector

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
)

// MockResourceSelector is a mock of ResourceSelector interface.
type MockResourceSelector struct {
	ctrl     *gomock.Controller
	recorder *MockResourceSelectorMockRecorder
}

// MockResourceSelectorMockRecorder is the mock recorder for MockResourceSelector.
type MockResourceSelectorMockRecorder struct {
	mock *MockResourceSelector
}

// NewMockResourceSelector creates a new mock instance.
func NewMockResourceSelector(ctrl *gomock.Controller) *MockResourceSelector {
	mock := &MockResourceSelector{ctrl: ctrl}
	mock.recorder = &MockResourceSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceSelector) EXPECT() *MockResourceSelectorMockRecorder {
	return m.recorder
}

// GetMeshServicesByServiceSelector mocks base method.
func (m *MockResourceSelector) GetMeshServicesByServiceSelector(ctx context.Context, selector *types.ServiceSelector) ([]*v1alpha1.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshServicesByServiceSelector", ctx, selector)
	ret0, _ := ret[0].([]*v1alpha1.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshServicesByServiceSelector indicates an expected call of GetMeshServicesByServiceSelector.
func (mr *MockResourceSelectorMockRecorder) GetMeshServicesByServiceSelector(ctx, selector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshServicesByServiceSelector", reflect.TypeOf((*MockResourceSelector)(nil).GetMeshServicesByServiceSelector), ctx, selector)
}

// GetMeshWorkloadsByIdentitySelector mocks base method.
func (m *MockResourceSelector) GetMeshWorkloadsByIdentitySelector(ctx context.Context, identitySelector *types.IdentitySelector) ([]*v1alpha1.MeshWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshWorkloadsByIdentitySelector", ctx, identitySelector)
	ret0, _ := ret[0].([]*v1alpha1.MeshWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshWorkloadsByIdentitySelector indicates an expected call of GetMeshWorkloadsByIdentitySelector.
func (mr *MockResourceSelectorMockRecorder) GetMeshWorkloadsByIdentitySelector(ctx, identitySelector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshWorkloadsByIdentitySelector", reflect.TypeOf((*MockResourceSelector)(nil).GetMeshWorkloadsByIdentitySelector), ctx, identitySelector)
}

// GetMeshWorkloadsByWorkloadSelector mocks base method.
func (m *MockResourceSelector) GetMeshWorkloadsByWorkloadSelector(ctx context.Context, workloadSelector *types.WorkloadSelector) ([]*v1alpha1.MeshWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshWorkloadsByWorkloadSelector", ctx, workloadSelector)
	ret0, _ := ret[0].([]*v1alpha1.MeshWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshWorkloadsByWorkloadSelector indicates an expected call of GetMeshWorkloadsByWorkloadSelector.
func (mr *MockResourceSelectorMockRecorder) GetMeshWorkloadsByWorkloadSelector(ctx, workloadSelector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshWorkloadsByWorkloadSelector", reflect.TypeOf((*MockResourceSelector)(nil).GetMeshWorkloadsByWorkloadSelector), ctx, workloadSelector)
}

// GetMeshServiceByRefSelector mocks base method.
func (m *MockResourceSelector) GetMeshServiceByRefSelector(ctx context.Context, kubeServiceName, kubeServiceNamespace, kubeServiceCluster string) (*v1alpha1.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshServiceByRefSelector", ctx, kubeServiceName, kubeServiceNamespace, kubeServiceCluster)
	ret0, _ := ret[0].(*v1alpha1.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshServiceByRefSelector indicates an expected call of GetMeshServiceByRefSelector.
func (mr *MockResourceSelectorMockRecorder) GetMeshServiceByRefSelector(ctx, kubeServiceName, kubeServiceNamespace, kubeServiceCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshServiceByRefSelector", reflect.TypeOf((*MockResourceSelector)(nil).GetMeshServiceByRefSelector), ctx, kubeServiceName, kubeServiceNamespace, kubeServiceCluster)
}

// GetMeshWorkloadByRefSelector mocks base method.
func (m *MockResourceSelector) GetMeshWorkloadByRefSelector(ctx context.Context, podEventWatcherName, podEventWatcherNamespace, podEventWatcherCluster string) (*v1alpha1.MeshWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshWorkloadByRefSelector", ctx, podEventWatcherName, podEventWatcherNamespace, podEventWatcherCluster)
	ret0, _ := ret[0].(*v1alpha1.MeshWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshWorkloadByRefSelector indicates an expected call of GetMeshWorkloadByRefSelector.
func (mr *MockResourceSelectorMockRecorder) GetMeshWorkloadByRefSelector(ctx, podEventWatcherName, podEventWatcherNamespace, podEventWatcherCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshWorkloadByRefSelector", reflect.TypeOf((*MockResourceSelector)(nil).GetMeshWorkloadByRefSelector), ctx, podEventWatcherName, podEventWatcherNamespace, podEventWatcherCluster)
}
