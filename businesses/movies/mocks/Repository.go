// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	movies "mini-project-movie-api/businesses/movies"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: movieDomain
func (_m *Repository) Create(movieDomain *movies.Domain) movies.Domain {
	ret := _m.Called(movieDomain)

	var r0 movies.Domain
	if rf, ok := ret.Get(0).(func(*movies.Domain) movies.Domain); ok {
		r0 = rf(movieDomain)
	} else {
		r0 = ret.Get(0).(movies.Domain)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id string) bool {
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
func (_m *Repository) GetAll() []movies.Domain {
	ret := _m.Called()

	var r0 []movies.Domain
	if rf, ok := ret.Get(0).(func() []movies.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movies.Domain)
		}
	}

	return r0
}

// GetByGenreID provides a mock function with given fields: genreId
func (_m *Repository) GetByGenreID(genreId string) []movies.Domain {
	ret := _m.Called(genreId)

	var r0 []movies.Domain
	if rf, ok := ret.Get(0).(func(string) []movies.Domain); ok {
		r0 = rf(genreId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movies.Domain)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id string) movies.Domain {
	ret := _m.Called(id)

	var r0 movies.Domain
	if rf, ok := ret.Get(0).(func(string) movies.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(movies.Domain)
	}

	return r0
}

// GetByTitle provides a mock function with given fields: title
func (_m *Repository) GetByTitle(title string) movies.Domain {
	ret := _m.Called(title)

	var r0 movies.Domain
	if rf, ok := ret.Get(0).(func(string) movies.Domain); ok {
		r0 = rf(title)
	} else {
		r0 = ret.Get(0).(movies.Domain)
	}

	return r0
}

// GetLatest provides a mock function with given fields:
func (_m *Repository) GetLatest() []movies.Domain {
	ret := _m.Called()

	var r0 []movies.Domain
	if rf, ok := ret.Get(0).(func() []movies.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movies.Domain)
		}
	}

	return r0
}

// Update provides a mock function with given fields: id, movieDomain
func (_m *Repository) Update(id string, movieDomain *movies.Domain) movies.Domain {
	ret := _m.Called(id, movieDomain)

	var r0 movies.Domain
	if rf, ok := ret.Get(0).(func(string, *movies.Domain) movies.Domain); ok {
		r0 = rf(id, movieDomain)
	} else {
		r0 = ret.Get(0).(movies.Domain)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
