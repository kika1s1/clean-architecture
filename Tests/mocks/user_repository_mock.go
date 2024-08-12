package mocks

import (
	"github.com/kika1s1/task_manager/domain"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

// Register provides a mock function with given fields: user
func (_m *MockUserRepository) Register(user domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByUsername provides a mock function with given fields: username
func (_m *MockUserRepository) FindByUsername(username string) (domain.User, error) {
	ret := _m.Called(username)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Promote provides a mock function with given fields: username
func (_m *MockUserRepository) Promote(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountUsers provides a mock function with given fields:
func (_m *MockUserRepository) CountUsers() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
