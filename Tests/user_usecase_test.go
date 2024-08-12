package Tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"

	"github.com/kika1s1/task_manager/domain"
	"github.com/kika1s1/task_manager/usecases"
    "github.com/kika1s1/task_manager/Tests/mocks"
)

func TestRegisterUser(t *testing.T) {
    mockUserRepo := new(mocks.UserRepositoryMock)
    mockUserRepo.On("RegisterUser", mock.Anything).Return(nil)

    userUsecase := Usecases.NewUserUsecase(mockUserRepo)
    user := Domain.User{
        Username: "testuser",
        Password: "hashedpassword",
    }

    err := userUsecase.RegisterUser(&user)

    assert.Nil(t, err)
    mockUserRepo.AssertExpectations(t)
}

func TestAuthenticateUser(t *testing.T) {
    mockUserRepo := new(mocks.UserRepositoryMock)
    mockUser := Domain.User{
        Username: "testuser",
        Password: "hashedpassword",
    }
    mockUserRepo.On("GetUserByUsername", "testuser").Return(&mockUser, nil)

    userUsecase := Usecases.NewUserUsecase(mockUserRepo)
    user, err := userUsecase.AuthenticateUser("testuser", "hashedpassword")

    assert.Nil(t, err)
    assert.Equal(t, "testuser", user.Username)
    mockUserRepo.AssertExpectations(t)
}
