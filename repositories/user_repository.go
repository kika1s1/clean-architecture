package repositories

import (
	"context"
	"errors"

	"github.com/kika1s1/task_manager/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		collection: client.Database("task_manager").Collection("users"),
	}
}

func (r *UserRepository) Register(user domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) Promote(username string) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": bson.M{"isAdmin": true}})
	return err
}

// count number of user exist 
func (r *UserRepository) CountUsers() (int64, error) {
	count, err := r.collection.CountDocuments(context.Background(), bson.D{})
	return count, err
}