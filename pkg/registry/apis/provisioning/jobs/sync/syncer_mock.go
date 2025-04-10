// Code generated by mockery v2.52.4. DO NOT EDIT.

package sync

import (
	context "context"

	jobs "github.com/grafana/grafana/pkg/registry/apis/provisioning/jobs"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/grafana/grafana/pkg/registry/apis/provisioning/repository"

	resources "github.com/grafana/grafana/pkg/registry/apis/provisioning/resources"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/provisioning/v0alpha1"
)

// MockSyncer is an autogenerated mock type for the Syncer type
type MockSyncer struct {
	mock.Mock
}

type MockSyncer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSyncer) EXPECT() *MockSyncer_Expecter {
	return &MockSyncer_Expecter{mock: &_m.Mock}
}

// Sync provides a mock function with given fields: ctx, repo, options, repositoryResources, clients, progress
func (_m *MockSyncer) Sync(ctx context.Context, repo repository.ReaderWriter, options v0alpha1.SyncJobOptions, repositoryResources resources.RepositoryResources, clients resources.ResourceClients, progress jobs.JobProgressRecorder) (string, error) {
	ret := _m.Called(ctx, repo, options, repositoryResources, clients, progress)

	if len(ret) == 0 {
		panic("no return value specified for Sync")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.ReaderWriter, v0alpha1.SyncJobOptions, resources.RepositoryResources, resources.ResourceClients, jobs.JobProgressRecorder) (string, error)); ok {
		return rf(ctx, repo, options, repositoryResources, clients, progress)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.ReaderWriter, v0alpha1.SyncJobOptions, resources.RepositoryResources, resources.ResourceClients, jobs.JobProgressRecorder) string); ok {
		r0 = rf(ctx, repo, options, repositoryResources, clients, progress)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.ReaderWriter, v0alpha1.SyncJobOptions, resources.RepositoryResources, resources.ResourceClients, jobs.JobProgressRecorder) error); ok {
		r1 = rf(ctx, repo, options, repositoryResources, clients, progress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSyncer_Sync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sync'
type MockSyncer_Sync_Call struct {
	*mock.Call
}

// Sync is a helper method to define mock.On call
//   - ctx context.Context
//   - repo repository.ReaderWriter
//   - options v0alpha1.SyncJobOptions
//   - repositoryResources resources.RepositoryResources
//   - clients resources.ResourceClients
//   - progress jobs.JobProgressRecorder
func (_e *MockSyncer_Expecter) Sync(ctx interface{}, repo interface{}, options interface{}, repositoryResources interface{}, clients interface{}, progress interface{}) *MockSyncer_Sync_Call {
	return &MockSyncer_Sync_Call{Call: _e.mock.On("Sync", ctx, repo, options, repositoryResources, clients, progress)}
}

func (_c *MockSyncer_Sync_Call) Run(run func(ctx context.Context, repo repository.ReaderWriter, options v0alpha1.SyncJobOptions, repositoryResources resources.RepositoryResources, clients resources.ResourceClients, progress jobs.JobProgressRecorder)) *MockSyncer_Sync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repository.ReaderWriter), args[2].(v0alpha1.SyncJobOptions), args[3].(resources.RepositoryResources), args[4].(resources.ResourceClients), args[5].(jobs.JobProgressRecorder))
	})
	return _c
}

func (_c *MockSyncer_Sync_Call) Return(_a0 string, _a1 error) *MockSyncer_Sync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSyncer_Sync_Call) RunAndReturn(run func(context.Context, repository.ReaderWriter, v0alpha1.SyncJobOptions, resources.RepositoryResources, resources.ResourceClients, jobs.JobProgressRecorder) (string, error)) *MockSyncer_Sync_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSyncer creates a new instance of MockSyncer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSyncer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSyncer {
	mock := &MockSyncer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
