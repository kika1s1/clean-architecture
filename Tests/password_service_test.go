package tests

import (
	"errors"
	"testing"

	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestCheckPasswordHardness(t *testing.T) {
	tests := []struct {
		password string
		expected error
	}{
		{"Password1!", nil}, // Valid password
		{"short", errors.New("password must be at least 8 characters long and include at least one uppercase letter, one lowercase letter, one number, and one special character")}, // Too short
		{"NoSpecial1", errors.New("password must be at least 8 characters long and include at least one uppercase letter, one lowercase letter, one number, and one special character")}, // Missing special character
		{"NoNumber!", errors.New("password must be at least 8 characters long and include at least one uppercase letter, one lowercase letter, one number, and one special character")}, // Missing number
		{"NoUppercase1!", errors.New("password must be at least 8 characters long and include at least one uppercase letter, one lowercase letter, one number, and one special character")}, // Missing uppercase letter
		{"NOLOWERCASE1!", errors.New("password must be at least 8 characters long and include at least one uppercase letter, one lowercase letter, one number, and one special character")}, // Missing lowercase letter
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			err := infrastructure.CheckPasswordHardness(tt.password)
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestHashPassword(t *testing.T) {
	password := "SecureP@ssw0rd"
	hashedPassword, err := infrastructure.HashPassword(password)

	// Ensure no error occurred during hashing
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Check that the hashed password can be validated
	isMatch := infrastructure.ComparePassword(hashedPassword, password)
	assert.True(t, isMatch)
}

func TestComparePassword(t *testing.T) {
	password := "Password123!"
	hashedPassword, err := infrastructure.HashPassword(password)
	assert.NoError(t, err)

	tests := []struct {
		hash     string
		pass     string
		expected bool
	}{
		{hashedPassword, password, true},             // Valid case
		{hashedPassword, "WrongPassword", false},     // Invalid password
		{"", "Password123!", false},                   // Empty hash
		{"InvalidHash", password, false},              // Invalid hash
	}

	for _, tt := range tests {
		t.Run(tt.pass, func(t *testing.T) {
			result := infrastructure.ComparePassword(tt.hash, tt.pass)
			assert.Equal(t, tt.expected, result)
		})
	}
}
