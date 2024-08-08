package domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	IsAdmin  bool   `bson:"isAdmin" json:"isAdmin"`
}

type Claims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
	ExpiresAt int64 `json:"expiresAt"`
}

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
}
