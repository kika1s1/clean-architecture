package tests
import (
	"testing"
	"time"


	"github.com/kika1s1/task_manager/domain"
	"github.com/stretchr/testify/assert"
)

func TestTaskModel(t *testing.T) {
	// Test Task creation
	task := &domain.Task{
		ID:        "1",
		Title:     "Test Task",
		Details:   "This is a test task",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	assert.Equal(t, "1", task.ID)
	assert.Equal(t, "Test Task", task.Title)
	assert.Equal(t, "This is a test task", task.Details)
	assert.NotNil(t, task.CreatedAt)
	assert.NotNil(t, task.UpdatedAt)

	// Test Task update
	newTitle := "Updated Task"
	task.Title = newTitle
	task.UpdatedAt = time.Now()

	assert.Equal(t, newTitle, task.Title)
	assert.NotEqual(t, task.CreatedAt, task.UpdatedAt)
}

func TestUserModel(t *testing.T) {
	// Test User creation
	user := &domain.User{
		ID:        "1",
		Username:  "testuser",
		Password:  "hashedpassword",
		Role:      "regular",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "hashedpassword", user.Password)
	assert.Equal(t, "regular", user.Role)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)

	// Test User update
	newUsername := "updateduser"
	user.Username = newUsername
	user.UpdatedAt = time.Now()

	assert.Equal(t, newUsername, user.Username)
	assert.NotEqual(t, user.CreatedAt, user.UpdatedAt)
}

func TestUserRoleChange(t *testing.T) {
	user := &domain.User{
		ID:        "1",
		Username:  "testuser",
		Password:  "hashedpassword",
		Role:      "regular",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Promote user to admin
	user.Role = "admin"
	user.UpdatedAt = time.Now()

	assert.Equal(t, "admin", user.Role)
}
