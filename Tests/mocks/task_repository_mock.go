
package mocks

import (
	"github.com/kika1s1/task_manager/domain"
	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) CreateTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *TaskRepositoryMock) GetTaskByID(id string) (*domain.Task, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *TaskRepositoryMock) UpdateTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *TaskRepositoryMock) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TaskRepositoryMock) GetAllTasks() ([]*domain.Task, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}
