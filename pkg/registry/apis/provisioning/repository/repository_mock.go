// Code generated by mockery v2.50.0. DO NOT EDIT.

package repository

import (
	context "context"
	http "net/http"

	field "k8s.io/apimachinery/pkg/util/validation/field"

	mock "github.com/stretchr/testify/mock"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/provisioning/v0alpha1"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// Config provides a mock function with no fields
func (_m *MockRepository) Config() *v0alpha1.Repository {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Config")
	}

	var r0 *v0alpha1.Repository
	if rf, ok := ret.Get(0).(func() *v0alpha1.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0alpha1.Repository)
		}
	}

	return r0
}

// Create provides a mock function with given fields: ctx, path, ref, data, message
func (_m *MockRepository) Create(ctx context.Context, path string, ref string, data []byte, message string) error {
	ret := _m.Called(ctx, path, ref, data, message)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte, string) error); ok {
		r0 = rf(ctx, path, ref, data, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, path, ref, message
func (_m *MockRepository) Delete(ctx context.Context, path string, ref string, message string) error {
	ret := _m.Called(ctx, path, ref, message)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, path, ref, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// History provides a mock function with given fields: ctx, path, ref
func (_m *MockRepository) History(ctx context.Context, path string, ref string) ([]v0alpha1.HistoryItem, error) {
	ret := _m.Called(ctx, path, ref)

	if len(ret) == 0 {
		panic("no return value specified for History")
	}

	var r0 []v0alpha1.HistoryItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]v0alpha1.HistoryItem, error)); ok {
		return rf(ctx, path, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []v0alpha1.HistoryItem); ok {
		r0 = rf(ctx, path, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v0alpha1.HistoryItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, path, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: ctx, path, ref
func (_m *MockRepository) Read(ctx context.Context, path string, ref string) (*FileInfo, error) {
	ret := _m.Called(ctx, path, ref)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 *FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*FileInfo, error)); ok {
		return rf(ctx, path, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *FileInfo); ok {
		r0 = rf(ctx, path, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, path, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadTree provides a mock function with given fields: ctx, ref
func (_m *MockRepository) ReadTree(ctx context.Context, ref string) ([]FileTreeEntry, error) {
	ret := _m.Called(ctx, ref)

	if len(ret) == 0 {
		panic("no return value specified for ReadTree")
	}

	var r0 []FileTreeEntry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]FileTreeEntry, error)); ok {
		return rf(ctx, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []FileTreeEntry); ok {
		r0 = rf(ctx, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]FileTreeEntry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Test provides a mock function with given fields: ctx
func (_m *MockRepository) Test(ctx context.Context) (*v0alpha1.TestResults, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Test")
	}

	var r0 *v0alpha1.TestResults
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*v0alpha1.TestResults, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *v0alpha1.TestResults); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0alpha1.TestResults)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, path, ref, data, message
func (_m *MockRepository) Update(ctx context.Context, path string, ref string, data []byte, message string) error {
	ret := _m.Called(ctx, path, ref, data, message)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte, string) error); ok {
		r0 = rf(ctx, path, ref, data, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validate provides a mock function with no fields
func (_m *MockRepository) Validate() field.ErrorList {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 field.ErrorList
	if rf, ok := ret.Get(0).(func() field.ErrorList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(field.ErrorList)
		}
	}

	return r0
}

// Webhook provides a mock function with given fields: ctx, req
func (_m *MockRepository) Webhook(ctx context.Context, req *http.Request) (*v0alpha1.WebhookResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Webhook")
	}

	var r0 *v0alpha1.WebhookResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) (*v0alpha1.WebhookResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *v0alpha1.WebhookResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0alpha1.WebhookResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: ctx, path, ref, data, message
func (_m *MockRepository) Write(ctx context.Context, path string, ref string, data []byte, message string) error {
	ret := _m.Called(ctx, path, ref, data, message)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte, string) error); ok {
		r0 = rf(ctx, path, ref, data, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
