// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	nationality "github.com/edukmx/nuitee/internal/domain/nationality"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindByIso provides a mock function with given fields: iso
func (_m *Repository) FindByIso(iso string) (*nationality.Nationality, error) {
	ret := _m.Called(iso)

	var r0 *nationality.Nationality
	if rf, ok := ret.Get(0).(func(string) *nationality.Nationality); ok {
		r0 = rf(iso)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nationality.Nationality)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(iso)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
