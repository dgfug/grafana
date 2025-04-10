// Code generated by mockery v2.52.4. DO NOT EDIT.

package sync

import (
	context "context"

	jobs "github.com/grafana/grafana/pkg/registry/apis/provisioning/jobs"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/grafana/grafana/pkg/registry/apis/provisioning/repository"

	resources "github.com/grafana/grafana/pkg/registry/apis/provisioning/resources"
)

// MockIncrementalSyncFn is an autogenerated mock type for the IncrementalSyncFn type
type MockIncrementalSyncFn struct {
	mock.Mock
}

type MockIncrementalSyncFn_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIncrementalSyncFn) EXPECT() *MockIncrementalSyncFn_Expecter {
	return &MockIncrementalSyncFn_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, repo, previousRef, currentRef, repositoryResources, progress
func (_m *MockIncrementalSyncFn) Execute(ctx context.Context, repo repository.Versioned, previousRef string, currentRef string, repositoryResources resources.RepositoryResources, progress jobs.JobProgressRecorder) error {
	ret := _m.Called(ctx, repo, previousRef, currentRef, repositoryResources, progress)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Versioned, string, string, resources.RepositoryResources, jobs.JobProgressRecorder) error); ok {
		r0 = rf(ctx, repo, previousRef, currentRef, repositoryResources, progress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIncrementalSyncFn_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockIncrementalSyncFn_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - repo repository.Versioned
//   - previousRef string
//   - currentRef string
//   - repositoryResources resources.RepositoryResources
//   - progress jobs.JobProgressRecorder
func (_e *MockIncrementalSyncFn_Expecter) Execute(ctx interface{}, repo interface{}, previousRef interface{}, currentRef interface{}, repositoryResources interface{}, progress interface{}) *MockIncrementalSyncFn_Execute_Call {
	return &MockIncrementalSyncFn_Execute_Call{Call: _e.mock.On("Execute", ctx, repo, previousRef, currentRef, repositoryResources, progress)}
}

func (_c *MockIncrementalSyncFn_Execute_Call) Run(run func(ctx context.Context, repo repository.Versioned, previousRef string, currentRef string, repositoryResources resources.RepositoryResources, progress jobs.JobProgressRecorder)) *MockIncrementalSyncFn_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repository.Versioned), args[2].(string), args[3].(string), args[4].(resources.RepositoryResources), args[5].(jobs.JobProgressRecorder))
	})
	return _c
}

func (_c *MockIncrementalSyncFn_Execute_Call) Return(_a0 error) *MockIncrementalSyncFn_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIncrementalSyncFn_Execute_Call) RunAndReturn(run func(context.Context, repository.Versioned, string, string, resources.RepositoryResources, jobs.JobProgressRecorder) error) *MockIncrementalSyncFn_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIncrementalSyncFn creates a new instance of MockIncrementalSyncFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIncrementalSyncFn(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIncrementalSyncFn {
	mock := &MockIncrementalSyncFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
