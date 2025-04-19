// Code generated by mockery v2.52.4. DO NOT EDIT.

package jobs

import (
	context "context"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/provisioning/v0alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MockJobProgressRecorder is an autogenerated mock type for the JobProgressRecorder type
type MockJobProgressRecorder struct {
	mock.Mock
}

type MockJobProgressRecorder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockJobProgressRecorder) EXPECT() *MockJobProgressRecorder_Expecter {
	return &MockJobProgressRecorder_Expecter{mock: &_m.Mock}
}

// Complete provides a mock function with given fields: ctx, err
func (_m *MockJobProgressRecorder) Complete(ctx context.Context, err error) v0alpha1.JobStatus {
	ret := _m.Called(ctx, err)

	if len(ret) == 0 {
		panic("no return value specified for Complete")
	}

	var r0 v0alpha1.JobStatus
	if rf, ok := ret.Get(0).(func(context.Context, error) v0alpha1.JobStatus); ok {
		r0 = rf(ctx, err)
	} else {
		r0 = ret.Get(0).(v0alpha1.JobStatus)
	}

	return r0
}

// MockJobProgressRecorder_Complete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Complete'
type MockJobProgressRecorder_Complete_Call struct {
	*mock.Call
}

// Complete is a helper method to define mock.On call
//   - ctx context.Context
//   - err error
func (_e *MockJobProgressRecorder_Expecter) Complete(ctx interface{}, err interface{}) *MockJobProgressRecorder_Complete_Call {
	return &MockJobProgressRecorder_Complete_Call{Call: _e.mock.On("Complete", ctx, err)}
}

func (_c *MockJobProgressRecorder_Complete_Call) Run(run func(ctx context.Context, err error)) *MockJobProgressRecorder_Complete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(error))
	})
	return _c
}

func (_c *MockJobProgressRecorder_Complete_Call) Return(_a0 v0alpha1.JobStatus) *MockJobProgressRecorder_Complete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJobProgressRecorder_Complete_Call) RunAndReturn(run func(context.Context, error) v0alpha1.JobStatus) *MockJobProgressRecorder_Complete_Call {
	_c.Call.Return(run)
	return _c
}

// Record provides a mock function with given fields: ctx, result
func (_m *MockJobProgressRecorder) Record(ctx context.Context, result JobResourceResult) {
	_m.Called(ctx, result)
}

// MockJobProgressRecorder_Record_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Record'
type MockJobProgressRecorder_Record_Call struct {
	*mock.Call
}

// Record is a helper method to define mock.On call
//   - ctx context.Context
//   - result JobResourceResult
func (_e *MockJobProgressRecorder_Expecter) Record(ctx interface{}, result interface{}) *MockJobProgressRecorder_Record_Call {
	return &MockJobProgressRecorder_Record_Call{Call: _e.mock.On("Record", ctx, result)}
}

func (_c *MockJobProgressRecorder_Record_Call) Run(run func(ctx context.Context, result JobResourceResult)) *MockJobProgressRecorder_Record_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(JobResourceResult))
	})
	return _c
}

func (_c *MockJobProgressRecorder_Record_Call) Return() *MockJobProgressRecorder_Record_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockJobProgressRecorder_Record_Call) RunAndReturn(run func(context.Context, JobResourceResult)) *MockJobProgressRecorder_Record_Call {
	_c.Run(run)
	return _c
}

// ResetResults provides a mock function with no fields
func (_m *MockJobProgressRecorder) ResetResults() {
	_m.Called()
}

// MockJobProgressRecorder_ResetResults_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetResults'
type MockJobProgressRecorder_ResetResults_Call struct {
	*mock.Call
}

// ResetResults is a helper method to define mock.On call
func (_e *MockJobProgressRecorder_Expecter) ResetResults() *MockJobProgressRecorder_ResetResults_Call {
	return &MockJobProgressRecorder_ResetResults_Call{Call: _e.mock.On("ResetResults")}
}

func (_c *MockJobProgressRecorder_ResetResults_Call) Run(run func()) *MockJobProgressRecorder_ResetResults_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJobProgressRecorder_ResetResults_Call) Return() *MockJobProgressRecorder_ResetResults_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockJobProgressRecorder_ResetResults_Call) RunAndReturn(run func()) *MockJobProgressRecorder_ResetResults_Call {
	_c.Run(run)
	return _c
}

// SetFinalMessage provides a mock function with given fields: ctx, msg
func (_m *MockJobProgressRecorder) SetFinalMessage(ctx context.Context, msg string) {
	_m.Called(ctx, msg)
}

// MockJobProgressRecorder_SetFinalMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFinalMessage'
type MockJobProgressRecorder_SetFinalMessage_Call struct {
	*mock.Call
}

// SetFinalMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - msg string
func (_e *MockJobProgressRecorder_Expecter) SetFinalMessage(ctx interface{}, msg interface{}) *MockJobProgressRecorder_SetFinalMessage_Call {
	return &MockJobProgressRecorder_SetFinalMessage_Call{Call: _e.mock.On("SetFinalMessage", ctx, msg)}
}

func (_c *MockJobProgressRecorder_SetFinalMessage_Call) Run(run func(ctx context.Context, msg string)) *MockJobProgressRecorder_SetFinalMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockJobProgressRecorder_SetFinalMessage_Call) Return() *MockJobProgressRecorder_SetFinalMessage_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockJobProgressRecorder_SetFinalMessage_Call) RunAndReturn(run func(context.Context, string)) *MockJobProgressRecorder_SetFinalMessage_Call {
	_c.Run(run)
	return _c
}

// SetMessage provides a mock function with given fields: ctx, msg
func (_m *MockJobProgressRecorder) SetMessage(ctx context.Context, msg string) {
	_m.Called(ctx, msg)
}

// MockJobProgressRecorder_SetMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetMessage'
type MockJobProgressRecorder_SetMessage_Call struct {
	*mock.Call
}

// SetMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - msg string
func (_e *MockJobProgressRecorder_Expecter) SetMessage(ctx interface{}, msg interface{}) *MockJobProgressRecorder_SetMessage_Call {
	return &MockJobProgressRecorder_SetMessage_Call{Call: _e.mock.On("SetMessage", ctx, msg)}
}

func (_c *MockJobProgressRecorder_SetMessage_Call) Run(run func(ctx context.Context, msg string)) *MockJobProgressRecorder_SetMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockJobProgressRecorder_SetMessage_Call) Return() *MockJobProgressRecorder_SetMessage_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockJobProgressRecorder_SetMessage_Call) RunAndReturn(run func(context.Context, string)) *MockJobProgressRecorder_SetMessage_Call {
	_c.Run(run)
	return _c
}

// SetTotal provides a mock function with given fields: ctx, total
func (_m *MockJobProgressRecorder) SetTotal(ctx context.Context, total int) {
	_m.Called(ctx, total)
}

// MockJobProgressRecorder_SetTotal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTotal'
type MockJobProgressRecorder_SetTotal_Call struct {
	*mock.Call
}

// SetTotal is a helper method to define mock.On call
//   - ctx context.Context
//   - total int
func (_e *MockJobProgressRecorder_Expecter) SetTotal(ctx interface{}, total interface{}) *MockJobProgressRecorder_SetTotal_Call {
	return &MockJobProgressRecorder_SetTotal_Call{Call: _e.mock.On("SetTotal", ctx, total)}
}

func (_c *MockJobProgressRecorder_SetTotal_Call) Run(run func(ctx context.Context, total int)) *MockJobProgressRecorder_SetTotal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockJobProgressRecorder_SetTotal_Call) Return() *MockJobProgressRecorder_SetTotal_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockJobProgressRecorder_SetTotal_Call) RunAndReturn(run func(context.Context, int)) *MockJobProgressRecorder_SetTotal_Call {
	_c.Run(run)
	return _c
}

// TooManyErrors provides a mock function with no fields
func (_m *MockJobProgressRecorder) TooManyErrors() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TooManyErrors")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockJobProgressRecorder_TooManyErrors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TooManyErrors'
type MockJobProgressRecorder_TooManyErrors_Call struct {
	*mock.Call
}

// TooManyErrors is a helper method to define mock.On call
func (_e *MockJobProgressRecorder_Expecter) TooManyErrors() *MockJobProgressRecorder_TooManyErrors_Call {
	return &MockJobProgressRecorder_TooManyErrors_Call{Call: _e.mock.On("TooManyErrors")}
}

func (_c *MockJobProgressRecorder_TooManyErrors_Call) Run(run func()) *MockJobProgressRecorder_TooManyErrors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockJobProgressRecorder_TooManyErrors_Call) Return(_a0 error) *MockJobProgressRecorder_TooManyErrors_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockJobProgressRecorder_TooManyErrors_Call) RunAndReturn(run func() error) *MockJobProgressRecorder_TooManyErrors_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockJobProgressRecorder creates a new instance of MockJobProgressRecorder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockJobProgressRecorder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockJobProgressRecorder {
	mock := &MockJobProgressRecorder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
