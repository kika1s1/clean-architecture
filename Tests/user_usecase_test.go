
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
	"github.com/kika1s1/task_manager/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockUserRepo.On("Create", mock.Anything).Return(nil)

	userUsecase := usecases.NewUserUsecase(mockUserRepo)

	router := gin.Default()
	router.POST("/register", userUsecase.Register)

	t.Run("Successful registration", func(t *testing.T) {
		// Prepare request payload
		requestBody := map[string]interface{}{
			"username": "testuser",
			"password": "testpassword",
		}
		jsonBody, _ := json.Marshal(requestBody)

		// Create request
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		res := httptest.NewRecorder()

		// Perform request
		router.ServeHTTP(res, req)

		// Check response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// Check response body
		expectedResponse := map[string]interface{}{
			"message": "Registration successful",
		}
		var responseBody map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &responseBody)
		assert.Equal(t, expectedResponse, responseBody)
	})

	t.Run("Failed registration - invalid request body", func(t *testing.T) {
		// Prepare request payload
		requestBody := map[string]interface{}{
			"username": "testuser",
		}
		jsonBody, _ := json.Marshal(requestBody)

		// Create request
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		res := httptest.NewRecorder()

		// Perform request
		router.ServeHTTP(res, req)

		// Check response status code
		assert.Equal(t, http.StatusBadRequest, res.Code)

		// Check response body
		expectedResponse := map[string]interface{}{
			"error": "Invalid request body",
		}
		var responseBody map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &responseBody)
		assert.Equal(t, expectedResponse, responseBody)
	})
}

func TestLogin(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockUserRepo.On("GetByUsername", mock.Anything).Return(&domain.User{
		Username: "testuser",
		Password: "testpassword",
	}, nil)

	userUsecase := usecases.NewUserUsecase(mockUserRepo)

	router := gin.Default()
	router.POST("/login", userUsecase.Login)

	t.Run("Successful login", func(t *testing.T) {
		// Prepare request payload
		requestBody := map[string]interface{}{
			"username": "testuser",
			"password": "testpassword",
		}
		jsonBody, _ := json.Marshal(requestBody)

		// Create request
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		res := httptest.NewRecorder()

		// Perform request
		router.ServeHTTP(res, req)

		// Check response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// Check response body
		expectedResponse := map[string]interface{}{
			"message": "Login successful",
		}
		var responseBody map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &responseBody)
		assert.Equal(t, expectedResponse, responseBody)
	})

	t.Run("Failed login - invalid credentials", func(t *testing.T) {
		// Prepare request payload
		requestBody := map[string]interface{}{
			"username": "testuser",
			"password": "wrongpassword",
		}
		jsonBody, _ := json.Marshal(requestBody)

		// Create request
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		res := httptest.NewRecorder()

		// Perform request
		router.ServeHTTP(res, req)

		// Check response status code
		assert.Equal(t, http.StatusUnauthorized, res.Code)

		// Check response body
		expectedResponse := map[string]interface{}{
			"error": "Invalid credentials",
		}
		var responseBody map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &responseBody)
		assert.Equal(t, expectedResponse, responseBody)
	})
}

func TestPromote(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockUserRepo.On("GetByUsername", mock.Anything).Return(&domain.User{
		Username: "testuser",
		IsAdmin:     true,
	}, nil)
	mockUserRepo.On("Update", mock.Anything).Return(nil)

	userUsecase := usecases.NewUserUsecase(mockUserRepo)

	router := gin.Default()
	router.POST("/promote", userUsecase.Promote)

	t.Run("Successful promotion", func(t *testing.T) {
		// Prepare request payload
		requestBody := map[string]interface{}{
			"username": "testuser",
		}
		jsonBody, _ := json.Marshal(requestBody)

		// Create request
		req, _ := http.NewRequest("POST", "/promote", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		res := httptest.NewRecorder()

		// Perform request
		router.ServeHTTP(res, req)

		// Check response status code
		assert.Equal(t, http.StatusOK, res.Code)

		// Check response body
		expectedResponse := map[string]interface{}{
			"message": "Promotion successful",
		}
		var responseBody map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &responseBody)
		assert.Equal(t, expectedResponse, responseBody)
	})

	t.Run("Failed promotion - user not found", func(t *testing.T) {
		// Prepare request payload
		requestBody := map[string]interface{}{
			"username": "nonexistentuser",
		}
		jsonBody, _ := json.Marshal(requestBody)

		// Create request
		req, _ := http.NewRequest("POST", "/promote", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		res := httptest.NewRecorder()

		// Perform request
		router.ServeHTTP(res, req)

		// Check response status code
		assert.Equal(t, http.StatusNotFound, res.Code)

		// Check response body
		expectedResponse := map[string]interface{}{
			"error": "User not found",
		}
		var responseBody map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &responseBody)
		assert.Equal(t, expectedResponse, responseBody)
	})
}
