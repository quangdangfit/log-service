// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/quangdangfit/log-service/internal/dto"
	mock "github.com/stretchr/testify/mock"

	paging "github.com/quangdangfit/log-service/pkg/paging"
)

// ILogService is an autogenerated mock type for the ILogService type
type ILogService struct {
	mock.Mock
}

// AddLog provides a mock function with given fields: ctx, req
func (_m *ILogService) AddLog(ctx context.Context, req *dto.AddLogReq) (*dto.Log, error) {
	ret := _m.Called(ctx, req)

	var r0 *dto.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.AddLogReq) (*dto.Log, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.AddLogReq) *dto.Log); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.AddLogReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: c, req
func (_m *ILogService) GetLogs(c context.Context, req *dto.GetLogsReq) (dto.Logs, *paging.Pagination, error) {
	ret := _m.Called(c, req)

	var r0 dto.Logs
	var r1 *paging.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetLogsReq) (dto.Logs, *paging.Pagination, error)); ok {
		return rf(c, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetLogsReq) dto.Logs); ok {
		r0 = rf(c, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dto.Logs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetLogsReq) *paging.Pagination); ok {
		r1 = rf(c, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*paging.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *dto.GetLogsReq) error); ok {
		r2 = rf(c, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewILogService creates a new instance of ILogService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewILogService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ILogService {
	mock := &ILogService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}