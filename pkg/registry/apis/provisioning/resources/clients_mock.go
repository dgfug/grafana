// Code generated by mockery v2.52.4. DO NOT EDIT.

package resources

import (
	mock "github.com/stretchr/testify/mock"
	dynamic "k8s.io/client-go/dynamic"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// MockResourceClients is an autogenerated mock type for the ResourceClients type
type MockResourceClients struct {
	mock.Mock
}

type MockResourceClients_Expecter struct {
	mock *mock.Mock
}

func (_m *MockResourceClients) EXPECT() *MockResourceClients_Expecter {
	return &MockResourceClients_Expecter{mock: &_m.Mock}
}

// Folder provides a mock function with no fields
func (_m *MockResourceClients) Folder() (dynamic.ResourceInterface, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Folder")
	}

	var r0 dynamic.ResourceInterface
	var r1 error
	if rf, ok := ret.Get(0).(func() (dynamic.ResourceInterface, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() dynamic.ResourceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.ResourceInterface)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResourceClients_Folder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Folder'
type MockResourceClients_Folder_Call struct {
	*mock.Call
}

// Folder is a helper method to define mock.On call
func (_e *MockResourceClients_Expecter) Folder() *MockResourceClients_Folder_Call {
	return &MockResourceClients_Folder_Call{Call: _e.mock.On("Folder")}
}

func (_c *MockResourceClients_Folder_Call) Run(run func()) *MockResourceClients_Folder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResourceClients_Folder_Call) Return(_a0 dynamic.ResourceInterface, _a1 error) *MockResourceClients_Folder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResourceClients_Folder_Call) RunAndReturn(run func() (dynamic.ResourceInterface, error)) *MockResourceClients_Folder_Call {
	_c.Call.Return(run)
	return _c
}

// ForKind provides a mock function with given fields: gvk
func (_m *MockResourceClients) ForKind(gvk schema.GroupVersionKind) (dynamic.ResourceInterface, schema.GroupVersionResource, error) {
	ret := _m.Called(gvk)

	if len(ret) == 0 {
		panic("no return value specified for ForKind")
	}

	var r0 dynamic.ResourceInterface
	var r1 schema.GroupVersionResource
	var r2 error
	if rf, ok := ret.Get(0).(func(schema.GroupVersionKind) (dynamic.ResourceInterface, schema.GroupVersionResource, error)); ok {
		return rf(gvk)
	}
	if rf, ok := ret.Get(0).(func(schema.GroupVersionKind) dynamic.ResourceInterface); ok {
		r0 = rf(gvk)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.ResourceInterface)
		}
	}

	if rf, ok := ret.Get(1).(func(schema.GroupVersionKind) schema.GroupVersionResource); ok {
		r1 = rf(gvk)
	} else {
		r1 = ret.Get(1).(schema.GroupVersionResource)
	}

	if rf, ok := ret.Get(2).(func(schema.GroupVersionKind) error); ok {
		r2 = rf(gvk)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockResourceClients_ForKind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForKind'
type MockResourceClients_ForKind_Call struct {
	*mock.Call
}

// ForKind is a helper method to define mock.On call
//   - gvk schema.GroupVersionKind
func (_e *MockResourceClients_Expecter) ForKind(gvk interface{}) *MockResourceClients_ForKind_Call {
	return &MockResourceClients_ForKind_Call{Call: _e.mock.On("ForKind", gvk)}
}

func (_c *MockResourceClients_ForKind_Call) Run(run func(gvk schema.GroupVersionKind)) *MockResourceClients_ForKind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(schema.GroupVersionKind))
	})
	return _c
}

func (_c *MockResourceClients_ForKind_Call) Return(_a0 dynamic.ResourceInterface, _a1 schema.GroupVersionResource, _a2 error) *MockResourceClients_ForKind_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockResourceClients_ForKind_Call) RunAndReturn(run func(schema.GroupVersionKind) (dynamic.ResourceInterface, schema.GroupVersionResource, error)) *MockResourceClients_ForKind_Call {
	_c.Call.Return(run)
	return _c
}

// ForResource provides a mock function with given fields: gvr
func (_m *MockResourceClients) ForResource(gvr schema.GroupVersionResource) (dynamic.ResourceInterface, schema.GroupVersionKind, error) {
	ret := _m.Called(gvr)

	if len(ret) == 0 {
		panic("no return value specified for ForResource")
	}

	var r0 dynamic.ResourceInterface
	var r1 schema.GroupVersionKind
	var r2 error
	if rf, ok := ret.Get(0).(func(schema.GroupVersionResource) (dynamic.ResourceInterface, schema.GroupVersionKind, error)); ok {
		return rf(gvr)
	}
	if rf, ok := ret.Get(0).(func(schema.GroupVersionResource) dynamic.ResourceInterface); ok {
		r0 = rf(gvr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.ResourceInterface)
		}
	}

	if rf, ok := ret.Get(1).(func(schema.GroupVersionResource) schema.GroupVersionKind); ok {
		r1 = rf(gvr)
	} else {
		r1 = ret.Get(1).(schema.GroupVersionKind)
	}

	if rf, ok := ret.Get(2).(func(schema.GroupVersionResource) error); ok {
		r2 = rf(gvr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockResourceClients_ForResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForResource'
type MockResourceClients_ForResource_Call struct {
	*mock.Call
}

// ForResource is a helper method to define mock.On call
//   - gvr schema.GroupVersionResource
func (_e *MockResourceClients_Expecter) ForResource(gvr interface{}) *MockResourceClients_ForResource_Call {
	return &MockResourceClients_ForResource_Call{Call: _e.mock.On("ForResource", gvr)}
}

func (_c *MockResourceClients_ForResource_Call) Run(run func(gvr schema.GroupVersionResource)) *MockResourceClients_ForResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(schema.GroupVersionResource))
	})
	return _c
}

func (_c *MockResourceClients_ForResource_Call) Return(_a0 dynamic.ResourceInterface, _a1 schema.GroupVersionKind, _a2 error) *MockResourceClients_ForResource_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockResourceClients_ForResource_Call) RunAndReturn(run func(schema.GroupVersionResource) (dynamic.ResourceInterface, schema.GroupVersionKind, error)) *MockResourceClients_ForResource_Call {
	_c.Call.Return(run)
	return _c
}

// User provides a mock function with no fields
func (_m *MockResourceClients) User() (dynamic.ResourceInterface, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for User")
	}

	var r0 dynamic.ResourceInterface
	var r1 error
	if rf, ok := ret.Get(0).(func() (dynamic.ResourceInterface, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() dynamic.ResourceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.ResourceInterface)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResourceClients_User_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'User'
type MockResourceClients_User_Call struct {
	*mock.Call
}

// User is a helper method to define mock.On call
func (_e *MockResourceClients_Expecter) User() *MockResourceClients_User_Call {
	return &MockResourceClients_User_Call{Call: _e.mock.On("User")}
}

func (_c *MockResourceClients_User_Call) Run(run func()) *MockResourceClients_User_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResourceClients_User_Call) Return(_a0 dynamic.ResourceInterface, _a1 error) *MockResourceClients_User_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResourceClients_User_Call) RunAndReturn(run func() (dynamic.ResourceInterface, error)) *MockResourceClients_User_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockResourceClients creates a new instance of MockResourceClients. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockResourceClients(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockResourceClients {
	mock := &MockResourceClients{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
