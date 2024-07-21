// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	models "github.com/djfemz/savannahTechTask/app/models"
	mock "github.com/stretchr/testify/mock"
)

// GithubAuxiliaryRepository is an autogenerated mock type for the GithubAuxiliaryRepository type
type GithubAuxiliaryRepository struct {
	mock.Mock
}

// FindById provides a mock function with given fields: id
func (_m *GithubAuxiliaryRepository) FindById(id uint) (*models.GithubAuxiliaryRepository, error) {
	ret := _m.Called(id)

	var r0 *models.GithubAuxiliaryRepository
	if rf, ok := ret.Get(0).(func(uint) *models.GithubAuxiliaryRepository); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GithubAuxiliaryRepository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: name
func (_m *GithubAuxiliaryRepository) FindByName(name string) (*models.GithubAuxiliaryRepository, error) {
	ret := _m.Called(name)

	var r0 *models.GithubAuxiliaryRepository
	if rf, ok := ret.Get(0).(func(string) *models.GithubAuxiliaryRepository); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GithubAuxiliaryRepository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: repository
func (_m *GithubAuxiliaryRepository) Save(repository *models.GithubAuxiliaryRepository) (*models.GithubAuxiliaryRepository, error) {
	ret := _m.Called(repository)

	var r0 *models.GithubAuxiliaryRepository
	if rf, ok := ret.Get(0).(func(*models.GithubAuxiliaryRepository) *models.GithubAuxiliaryRepository); ok {
		r0 = rf(repository)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GithubAuxiliaryRepository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.GithubAuxiliaryRepository) error); ok {
		r1 = rf(repository)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGithubAuxiliaryRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewGithubAuxiliaryRepository creates a new instance of GithubAuxiliaryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGithubAuxiliaryRepository(t mockConstructorTestingTNewGithubAuxiliaryRepository) *GithubAuxiliaryRepository {
	mock := &GithubAuxiliaryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
