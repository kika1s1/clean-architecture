package Tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupEnv() {
    os.Setenv("JWT_SECRET", "mysecret")
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
    setupEnv()

    router := gin.Default()
    middleware := infrastructure.AuthMiddleware
    router.GET("/test", middleware, func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "success"})
    })

    token := generateToken(t, "1", "testuser")

    req, _ := http.NewRequest("GET", "/test", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    assert.JSONEq(t, `{"message": "success"}`, resp.Body.String())
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
    setupEnv()

    router := gin.Default()
    middleware := infrastructure.AuthMiddleware
    router.GET("/test", middleware, func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "success"})
    })

    req, _ := http.NewRequest("GET", "/test", nil)
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusUnauthorized, resp.Code)
    assert.JSONEq(t, `{"error": "Missing or malformed token"}`, resp.Body.String())
}

func TestAuthMiddleware_MalformedToken(t *testing.T) {
    setupEnv()

    router := gin.Default()
    middleware := infrastructure.AuthMiddleware
    router.GET("/test", middleware, func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "success"})
    })

    req, _ := http.NewRequest("GET", "/test", nil)
    req.Header.Set("Authorization", "Bearer invalidtoken")
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusUnauthorized, resp.Code)
    assert.JSONEq(t, `{"error": "Invalid token"}`, resp.Body.String())
}

func TestAuthMiddleware_InvalidTokenClaims(t *testing.T) {
    setupEnv()

    router := gin.Default()
    middleware := infrastructure.AuthMiddleware
    router.GET("/test", middleware, func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "success"})
    })

    // Generate a token with invalid claims
    invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyIn0.invalidsignature"

    req, _ := http.NewRequest("GET", "/test", nil)
    req.Header.Set("Authorization", "Bearer "+invalidToken)
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusUnauthorized, resp.Code)
    assert.JSONEq(t, `{"error": "Invalid token claims"}`, resp.Body.String())
}

// Helper function to generate JWT token
func generateToken(t *testing.T, userID, username string) string {
    jwtSecret := "mysecret"
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Domain.Claims{
        UserID:   userID,
        Username: username,
    })

    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        t.Fatalf("Failed to generate token: %v", err)
    }

    return tokenString
}
