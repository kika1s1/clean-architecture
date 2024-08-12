package Tests

import (
	"testing"

	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
    jwtService := infrastructure.NewJWTService()
    token, err := jwtService.GenerateToken("1", "testuser")

    assert.Nil(t, err)
    assert.NotEmpty(t, token)
}

func TestValidateJWT(t *testing.T) {
    jwtService := Infrastructure.NewJWTService()
    token, _ := jwtService.GenerateToken("1", "testuser")

    claims, err := jwtService.ValidateToken(token)

    assert.Nil(t, err)
    assert.Equal(t, "1", claims.UserID)
    assert.Equal(t, "testuser", claims.Username)
}
