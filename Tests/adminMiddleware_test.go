package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/domain"
	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestAdminMiddleware_Authorized(t *testing.T) {
	// Set up Gin and the AdminMiddleware
	r := gin.New()
	r.Use(func(c *gin.Context) {
		// Simulate setting the "isAdmin" key in the context
		c.Set("isAdmin", &domain.Claims{IsAdmin: true})
	})
	r.Use(infrastructure.AdminMiddleware())
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Admin Authorized"})
	})

	// Create a request
	req, _ := http.NewRequest("GET", "/admin", nil)

	// Test the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Admin Authorized"}`, w.Body.String())
}

func TestAdminMiddleware_Unauthorized(t *testing.T) {
	// Set up Gin and the AdminMiddleware
	r := gin.New()
	r.Use(func(c *gin.Context) {
		// Simulate no "isAdmin" key in the context
	})
	r.Use(infrastructure.AdminMiddleware())
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Admin Authorized"})
	})

	// Create a request
	req, _ := http.NewRequest("GET", "/admin", nil)

	// Test the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "Unauthorized"}`, w.Body.String())
}

func TestAdminMiddleware_Forbidden(t *testing.T) {
	// Set up Gin and the AdminMiddleware
	r := gin.New()
	r.Use(func(c *gin.Context) {
		// Simulate setting the "isAdmin" key in the context with IsAdmin set to false
		c.Set("isAdmin", &domain.Claims{IsAdmin: false})
	})
	r.Use(infrastructure.AdminMiddleware())
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Admin Authorized"})
	})

	// Create a request
	req, _ := http.NewRequest("GET", "/admin", nil)

	// Test the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.JSONEq(t, `{"error": "Forbidden"}`, w.Body.String())
}
