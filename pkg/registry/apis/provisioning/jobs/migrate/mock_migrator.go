// Code generated by mockery v2.52.4. DO NOT EDIT.

package migrate

import (
	context "context"

	jobs "github.com/grafana/grafana/pkg/registry/apis/provisioning/jobs"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/grafana/grafana/pkg/registry/apis/provisioning/repository"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/provisioning/v0alpha1"
)

// MockMigrator is an autogenerated mock type for the Migrator type
type MockMigrator struct {
	mock.Mock
}

type MockMigrator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMigrator) EXPECT() *MockMigrator_Expecter {
	return &MockMigrator_Expecter{mock: &_m.Mock}
}

// Migrate provides a mock function with given fields: ctx, rw, opts, progress
func (_m *MockMigrator) Migrate(ctx context.Context, rw repository.ReaderWriter, opts v0alpha1.MigrateJobOptions, progress jobs.JobProgressRecorder) error {
	ret := _m.Called(ctx, rw, opts, progress)

	if len(ret) == 0 {
		panic("no return value specified for Migrate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.ReaderWriter, v0alpha1.MigrateJobOptions, jobs.JobProgressRecorder) error); ok {
		r0 = rf(ctx, rw, opts, progress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMigrator_Migrate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Migrate'
type MockMigrator_Migrate_Call struct {
	*mock.Call
}

// Migrate is a helper method to define mock.On call
//   - ctx context.Context
//   - rw repository.ReaderWriter
//   - opts v0alpha1.MigrateJobOptions
//   - progress jobs.JobProgressRecorder
func (_e *MockMigrator_Expecter) Migrate(ctx interface{}, rw interface{}, opts interface{}, progress interface{}) *MockMigrator_Migrate_Call {
	return &MockMigrator_Migrate_Call{Call: _e.mock.On("Migrate", ctx, rw, opts, progress)}
}

func (_c *MockMigrator_Migrate_Call) Run(run func(ctx context.Context, rw repository.ReaderWriter, opts v0alpha1.MigrateJobOptions, progress jobs.JobProgressRecorder)) *MockMigrator_Migrate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repository.ReaderWriter), args[2].(v0alpha1.MigrateJobOptions), args[3].(jobs.JobProgressRecorder))
	})
	return _c
}

func (_c *MockMigrator_Migrate_Call) Return(_a0 error) *MockMigrator_Migrate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMigrator_Migrate_Call) RunAndReturn(run func(context.Context, repository.ReaderWriter, v0alpha1.MigrateJobOptions, jobs.JobProgressRecorder) error) *MockMigrator_Migrate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMigrator creates a new instance of MockMigrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMigrator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMigrator {
	mock := &MockMigrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
