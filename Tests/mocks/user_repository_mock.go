

package mocks

import (
	"github.com/kika1s1/task_manager/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetUserByID(id string) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetUserByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) UpdateUser(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetAllUsers() ([]*domain.User, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}
