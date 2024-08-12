package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/domain"
	"github.com/kika1s1/task_manager/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockTaskRepository is a mock implementation of TaskRepository.
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(id primitive.ObjectID) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(id primitive.ObjectID, updatedTask domain.Task) error {
	args := m.Called(id, updatedTask)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}

// Remove the existing declaration of setupRouter function

func TestCreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	usecase := usecases.NewTaskUsecase(mockRepo)

	task := domain.Task{Title: "Test Task"}
	mockRepo.On("CreateTask", task).Return(nil)

	payload, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.CreateTask(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message": "Task created successfully"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestGetTasks(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	usecase := usecases.NewTaskUsecase(mockRepo)

	tasks := []domain.Task{{Title: "Test Task"}}
	mockRepo.On("GetTasks").Return(tasks, nil)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.GetTasks(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[{"Title":"Test Task"}]`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	usecase := usecases.NewTaskUsecase(mockRepo)

	id, _ := primitive.ObjectIDFromHex("605c72efb6a7f63a9d3e58b2")
	task := domain.Task{Title: "Test Task"}
	mockRepo.On("GetTaskByID", id).Return(task, nil)

	req, _ := http.NewRequest("GET", "/tasks/"+id.Hex(), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.GetTaskByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"Title":"Test Task"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	usecase := usecases.NewTaskUsecase(mockRepo)

	id, _ := primitive.ObjectIDFromHex("605c72efb6a7f63a9d3e58b2")
	updatedTask := domain.Task{Title: "Updated Task"}
	mockRepo.On("UpdateTask", id, updatedTask).Return(nil)

	payload, _ := json.Marshal(updatedTask)
	req, _ := http.NewRequest("PUT", "/tasks/"+id.Hex(), bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.UpdateTask(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Task updated successfully"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	usecase := usecases.NewTaskUsecase(mockRepo)

	id, _ := primitive.ObjectIDFromHex("605c72efb6a7f63a9d3e58b2")
	mockRepo.On("DeleteTask", id).Return(nil)

	req, _ := http.NewRequest("DELETE", "/tasks/"+id.Hex(), nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.DeleteTask(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Task deleted successfully"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}
