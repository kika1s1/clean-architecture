package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/domain"
	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/kika1s1/task_manager/repositories"
	"github.com/kika1s1/task_manager/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByUsername(username string) (domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) Promote(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserRepository) CountUsers() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestRegister(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecases.NewUserUsecase(mockRepo)

	user := domain.User{
		Username: "testuser",
		Password: "P@ssw0rd",
	}

	mockRepo.On("FindByUsername", user.Username).Return(domain.User{}, nil)
	mockRepo.On("CountUsers").Return(int64(0), nil)
	mockRepo.On("Register", mock.Anything).Return(nil)

	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.Register(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message": "User registered successfully"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestRegister_UserExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecases.NewUserUsecase(mockRepo)

	user := domain.User{
		Username: "testuser",
		Password: "P@ssw0rd",
	}

	mockRepo.On("FindByUsername", user.Username).Return(user, nil)

	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.Register(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.JSONEq(t, `{"error": "Username already exists"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecases.NewUserUsecase(mockRepo)

	user := domain.User{
		Username: "testuser",
		Password: "P@ssw0rd",
	}

	mockRepo.On("FindByUsername", user.Username).Return(user, nil)
	infrastructure.ComparePassword = func(hashedPassword, password string) bool {
		return true
	}
	infrastructure.GenerateJWT = func(username string, isAdmin bool) (string, error) {
		return "mockToken", nil
	}

	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"token": "mockToken"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecases.NewUserUsecase(mockRepo)

	user := domain.User{
		Username: "testuser",
		Password: "P@ssw0rd",
	}

	mockRepo.On("FindByUsername", user.Username).Return(domain.User{}, errors.New("user not found"))

	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "Invalid credentials"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestPromote(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecases.NewUserUsecase(mockRepo)

	username := "testuser"
	mockRepo.On("Promote", username).Return(nil)

	req, _ := http.NewRequest("POST", "/promote/"+username, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.Promote(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "User promoted successfully"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestPromote_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecases.NewUserUsecase(mockRepo)

	username := "testuser"
	mockRepo.On("Promote", username).Return(errors.New("promotion error"))

	req, _ := http.NewRequest("POST", "/promote/"+username, nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	usecase.Promote(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "Error promoting user"}`, w.Body.String())
	mockRepo.AssertExpectations(t)
}
