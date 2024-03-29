// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	ratings "mini-project-movie-api/businesses/ratings"

	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ratingDomain
func (_m *UseCase) Create(ratingDomain *ratings.Domain) ratings.Domain {
	ret := _m.Called(ratingDomain)

	var r0 ratings.Domain
	if rf, ok := ret.Get(0).(func(*ratings.Domain) ratings.Domain); ok {
		r0 = rf(ratingDomain)
	} else {
		r0 = ret.Get(0).(ratings.Domain)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *UseCase) Delete(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *UseCase) GetAll() []ratings.Domain {
	ret := _m.Called()

	var r0 []ratings.Domain
	if rf, ok := ret.Get(0).(func() []ratings.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ratings.Domain)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *UseCase) GetByID(id string) ratings.Domain {
	ret := _m.Called(id)

	var r0 ratings.Domain
	if rf, ok := ret.Get(0).(func(string) ratings.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(ratings.Domain)
	}

	return r0
}

// GetByMovieID provides a mock function with given fields: movieId
func (_m *UseCase) GetByMovieID(movieId string) []ratings.Domain {
	ret := _m.Called(movieId)

	var r0 []ratings.Domain
	if rf, ok := ret.Get(0).(func(string) []ratings.Domain); ok {
		r0 = rf(movieId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ratings.Domain)
		}
	}

	return r0
}

// GetByMovieIdAndUserID provides a mock function with given fields: movieId, userId
func (_m *UseCase) GetByMovieIdAndUserID(movieId string, userId string) ratings.Domain {
	ret := _m.Called(movieId, userId)

	var r0 ratings.Domain
	if rf, ok := ret.Get(0).(func(string, string) ratings.Domain); ok {
		r0 = rf(movieId, userId)
	} else {
		r0 = ret.Get(0).(ratings.Domain)
	}

	return r0
}

// GetByUserID provides a mock function with given fields: userId
func (_m *UseCase) GetByUserID(userId string) []ratings.Domain {
	ret := _m.Called(userId)

	var r0 []ratings.Domain
	if rf, ok := ret.Get(0).(func(string) []ratings.Domain); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ratings.Domain)
		}
	}

	return r0
}

// Update provides a mock function with given fields: id, ratingDomain
func (_m *UseCase) Update(id string, ratingDomain *ratings.Domain) ratings.Domain {
	ret := _m.Called(id, ratingDomain)

	var r0 ratings.Domain
	if rf, ok := ret.Get(0).(func(string, *ratings.Domain) ratings.Domain); ok {
		r0 = rf(id, ratingDomain)
	} else {
		r0 = ret.Get(0).(ratings.Domain)
	}

	return r0
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
