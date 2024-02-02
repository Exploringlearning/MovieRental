// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	dto "movierental/internal/dto"

	mock "github.com/stretchr/testify/mock"
)

// Movie is an autogenerated mock type for the Movie type
type Movie struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *Movie) GetAll() ([]dto.Movie, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []dto.Movie
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]dto.Movie, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []dto.Movie); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Movie)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMovie creates a new instance of Movie. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMovie(t interface {
	mock.TestingT
	Cleanup(func())
}) *Movie {
	mock := &Movie{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
