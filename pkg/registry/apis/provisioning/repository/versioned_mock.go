// Code generated by mockery v2.52.4. DO NOT EDIT.

package repository

import (
	context "context"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/provisioning/v0alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MockVersioned is an autogenerated mock type for the Versioned type
type MockVersioned struct {
	mock.Mock
}

type MockVersioned_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVersioned) EXPECT() *MockVersioned_Expecter {
	return &MockVersioned_Expecter{mock: &_m.Mock}
}

// CompareFiles provides a mock function with given fields: ctx, base, ref
func (_m *MockVersioned) CompareFiles(ctx context.Context, base string, ref string) ([]VersionedFileChange, error) {
	ret := _m.Called(ctx, base, ref)

	if len(ret) == 0 {
		panic("no return value specified for CompareFiles")
	}

	var r0 []VersionedFileChange
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]VersionedFileChange, error)); ok {
		return rf(ctx, base, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []VersionedFileChange); ok {
		r0 = rf(ctx, base, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]VersionedFileChange)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, base, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVersioned_CompareFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompareFiles'
type MockVersioned_CompareFiles_Call struct {
	*mock.Call
}

// CompareFiles is a helper method to define mock.On call
//   - ctx context.Context
//   - base string
//   - ref string
func (_e *MockVersioned_Expecter) CompareFiles(ctx interface{}, base interface{}, ref interface{}) *MockVersioned_CompareFiles_Call {
	return &MockVersioned_CompareFiles_Call{Call: _e.mock.On("CompareFiles", ctx, base, ref)}
}

func (_c *MockVersioned_CompareFiles_Call) Run(run func(ctx context.Context, base string, ref string)) *MockVersioned_CompareFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockVersioned_CompareFiles_Call) Return(_a0 []VersionedFileChange, _a1 error) *MockVersioned_CompareFiles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVersioned_CompareFiles_Call) RunAndReturn(run func(context.Context, string, string) ([]VersionedFileChange, error)) *MockVersioned_CompareFiles_Call {
	_c.Call.Return(run)
	return _c
}

// History provides a mock function with given fields: ctx, path, ref
func (_m *MockVersioned) History(ctx context.Context, path string, ref string) ([]v0alpha1.HistoryItem, error) {
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

// MockVersioned_History_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'History'
type MockVersioned_History_Call struct {
	*mock.Call
}

// History is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
//   - ref string
func (_e *MockVersioned_Expecter) History(ctx interface{}, path interface{}, ref interface{}) *MockVersioned_History_Call {
	return &MockVersioned_History_Call{Call: _e.mock.On("History", ctx, path, ref)}
}

func (_c *MockVersioned_History_Call) Run(run func(ctx context.Context, path string, ref string)) *MockVersioned_History_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockVersioned_History_Call) Return(_a0 []v0alpha1.HistoryItem, _a1 error) *MockVersioned_History_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVersioned_History_Call) RunAndReturn(run func(context.Context, string, string) ([]v0alpha1.HistoryItem, error)) *MockVersioned_History_Call {
	_c.Call.Return(run)
	return _c
}

// LatestRef provides a mock function with given fields: ctx
func (_m *MockVersioned) LatestRef(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LatestRef")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVersioned_LatestRef_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LatestRef'
type MockVersioned_LatestRef_Call struct {
	*mock.Call
}

// LatestRef is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockVersioned_Expecter) LatestRef(ctx interface{}) *MockVersioned_LatestRef_Call {
	return &MockVersioned_LatestRef_Call{Call: _e.mock.On("LatestRef", ctx)}
}

func (_c *MockVersioned_LatestRef_Call) Run(run func(ctx context.Context)) *MockVersioned_LatestRef_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockVersioned_LatestRef_Call) Return(_a0 string, _a1 error) *MockVersioned_LatestRef_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVersioned_LatestRef_Call) RunAndReturn(run func(context.Context) (string, error)) *MockVersioned_LatestRef_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVersioned creates a new instance of MockVersioned. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVersioned(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVersioned {
	mock := &MockVersioned{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
