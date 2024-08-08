package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/domain"
	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/kika1s1/task_manager/repositories"
	"net/http"
)

type UserUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = hashedPassword

	if err := u.userRepo.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (u *UserUsecase) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedUser, err := u.userRepo.FindByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if infrastructure.ComparePassword(storedUser.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := infrastructure.GenerateJWT(storedUser.Username, storedUser.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserUsecase) Promote(c *gin.Context) {
	username := c.Param("username")

	err := u.userRepo.Promote(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error promoting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}
