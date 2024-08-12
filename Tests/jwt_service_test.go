package tests

import (
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kika1s1/task_manager/domain"
	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	// Set up environment variable for JWT_SECRET
	os.Setenv("JWT_SECRET", "testsecret")

	// Generate a JWT token
	tokenString, err := infrastructure.GenerateJWT("testuser", true)

	// Ensure no error occurred during token generation
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Parse the token to verify its claims
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("testsecret"), nil
	})

	// Ensure no error occurred during token parsing
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	// Extract and validate the claims
	claims, ok := token.Claims.(*domain.Claims)
	assert.True(t, ok)
	assert.Equal(t, "testuser", claims.Username)
	assert.True(t, claims.IsAdmin)

	// Check if the token expiration time is approximately 24 hours from now
	assert.WithinDuration(t, time.Now().Add(24*time.Hour), time.Unix(claims.ExpiresAt, 0), time.Minute)
}
