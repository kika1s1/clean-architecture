package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/domain"
	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/stretchr/testify/assert"
)

func generateToken(secret string, claims domain.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	// Set up environment variable for JWT_SECRET
	os.Setenv("JWT_SECRET", "testsecret")

	// Create a valid token
	claims := domain.Claims{
		Username: "testuser",
		IsAdmin:  true,
	}
	tokenString := generateToken("testsecret", claims)

	// Set up Gin and the AuthMiddleware
	r := gin.New()
	r.Use(infrastructure.AuthMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized"})
	})

	// Create a request with the valid token
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	// Test the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Authorized"}`, w.Body.String())
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	// Set up environment variable for JWT_SECRET
	os.Setenv("JWT_SECRET", "testsecret")

	// Set up Gin and the AuthMiddleware
	r := gin.New()
	r.Use(infrastructure.AuthMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized"})
	})

	// Create a request with an invalid token
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")

	// Test the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "Invalid token"}`, w.Body.String())
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	// Set up Gin and the AuthMiddleware
	r := gin.New()
	r.Use(infrastructure.AuthMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized"})
	})

	// Create a request with no token
	req, _ := http.NewRequest("GET", "/protected", nil)

	// Test the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "Missing or malformed token"}`, w.Body.String())
}
