package Tests

import (
    "testing"

    "github.com/kika1s1/task_manager/usecases"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"

    "github.com/kika1s1/task_manager/domain"
    "github.com/kika1s1/task_manager/Tests/mocks"
)

func TestCreateTask(t *testing.T) {
    mockTaskRepo := new(mocks.TaskRepositoryMock)
    mockTaskRepo.On("CreateTask", mock.Anything).Return(nil)

    taskUsecase := usecases.NewTaskUsecase(mockTaskRepo)
    task := Domain.Task{
        Title:       "Test Task",
        Description: "This is a test task.",
    }

    err := taskUsecase.CreateTask(&task)

    assert.Nil(t, err)
    mockTaskRepo.AssertExpectations(t)
}

func TestGetTaskByID(t *testing.T) {
    mockTaskRepo := new(mocks.TaskRepositoryMock)
    mockTask := Domain.Task{
        ID:    "1",
        Title: "Test Task",
    }
    mockTaskRepo.On("GetTaskByID", "1").Return(&mockTask, nil)

    taskUsecase := Usecases.NewTaskUsecase(mockTaskRepo)
    task, err := taskUsecase.GetTaskByID("1")

    assert.Nil(t, err)
    assert.Equal(t, "Test Task", task.Title)
    mockTaskRepo.AssertExpectations(t)
}
