package infrastructure

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/kika1s1/task_manager/domain"

	"log"
	"os"
)

func GenerateJWT(username string, isAdmin bool) (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file +: %s", err)
	}
	JWT_SECRET := os.Getenv("JWT_SECRET")
	claims := domain.Claims{
		Username:  username,
		IsAdmin:   isAdmin,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET))
}
