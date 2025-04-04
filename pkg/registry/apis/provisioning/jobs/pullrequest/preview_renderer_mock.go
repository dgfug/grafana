// Code generated by mockery v2.52.4. DO NOT EDIT.

package pullrequest

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockPreviewRenderer is an autogenerated mock type for the PreviewRenderer type
type MockPreviewRenderer struct {
	mock.Mock
}

type MockPreviewRenderer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPreviewRenderer) EXPECT() *MockPreviewRenderer_Expecter {
	return &MockPreviewRenderer_Expecter{mock: &_m.Mock}
}

// IsAvailable provides a mock function with given fields: ctx
func (_m *MockPreviewRenderer) IsAvailable(ctx context.Context) bool {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsAvailable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPreviewRenderer_IsAvailable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAvailable'
type MockPreviewRenderer_IsAvailable_Call struct {
	*mock.Call
}

// IsAvailable is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPreviewRenderer_Expecter) IsAvailable(ctx interface{}) *MockPreviewRenderer_IsAvailable_Call {
	return &MockPreviewRenderer_IsAvailable_Call{Call: _e.mock.On("IsAvailable", ctx)}
}

func (_c *MockPreviewRenderer_IsAvailable_Call) Run(run func(ctx context.Context)) *MockPreviewRenderer_IsAvailable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPreviewRenderer_IsAvailable_Call) Return(_a0 bool) *MockPreviewRenderer_IsAvailable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPreviewRenderer_IsAvailable_Call) RunAndReturn(run func(context.Context) bool) *MockPreviewRenderer_IsAvailable_Call {
	_c.Call.Return(run)
	return _c
}

// RenderDashboardPreview provides a mock function with given fields: ctx, namespace, repoName, path, ref
func (_m *MockPreviewRenderer) RenderDashboardPreview(ctx context.Context, namespace string, repoName string, path string, ref string) (string, error) {
	ret := _m.Called(ctx, namespace, repoName, path, ref)

	if len(ret) == 0 {
		panic("no return value specified for RenderDashboardPreview")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) (string, error)); ok {
		return rf(ctx, namespace, repoName, path, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) string); ok {
		r0 = rf(ctx, namespace, repoName, path, ref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, namespace, repoName, path, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPreviewRenderer_RenderDashboardPreview_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RenderDashboardPreview'
type MockPreviewRenderer_RenderDashboardPreview_Call struct {
	*mock.Call
}

// RenderDashboardPreview is a helper method to define mock.On call
//   - ctx context.Context
//   - namespace string
//   - repoName string
//   - path string
//   - ref string
func (_e *MockPreviewRenderer_Expecter) RenderDashboardPreview(ctx interface{}, namespace interface{}, repoName interface{}, path interface{}, ref interface{}) *MockPreviewRenderer_RenderDashboardPreview_Call {
	return &MockPreviewRenderer_RenderDashboardPreview_Call{Call: _e.mock.On("RenderDashboardPreview", ctx, namespace, repoName, path, ref)}
}

func (_c *MockPreviewRenderer_RenderDashboardPreview_Call) Run(run func(ctx context.Context, namespace string, repoName string, path string, ref string)) *MockPreviewRenderer_RenderDashboardPreview_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockPreviewRenderer_RenderDashboardPreview_Call) Return(_a0 string, _a1 error) *MockPreviewRenderer_RenderDashboardPreview_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPreviewRenderer_RenderDashboardPreview_Call) RunAndReturn(run func(context.Context, string, string, string, string) (string, error)) *MockPreviewRenderer_RenderDashboardPreview_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPreviewRenderer creates a new instance of MockPreviewRenderer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPreviewRenderer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPreviewRenderer {
	mock := &MockPreviewRenderer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
