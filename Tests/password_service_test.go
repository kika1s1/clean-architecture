package Tests

import (
	"testing"

	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
    passwordService := infrastructure.NewPasswordService()
    hashedPassword, err := passwordService.HashPassword("password123")

    assert.Nil(t, err)
    assert.NotEmpty(t, hashedPassword)
}

func TestCheckPasswordHash(t *testing.T) {
    passwordService := Infrastructure.NewPasswordService()
    hashedPassword, _ := passwordService.HashPassword("password123")

    match := passwordService.CheckPasswordHash("password123", hashedPassword)

    assert.True(t, match)
}
