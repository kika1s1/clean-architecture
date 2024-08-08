package repositories

import (
	"context"

	"github.com/kika1s1/task_manager/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(client *mongo.Client) *TaskRepository {
	return &TaskRepository{
		collection: client.Database("task_manager").Collection("tasks"),
	}
}

func (r *TaskRepository) CreateTask(task domain.Task) error {
	_, err := r.collection.InsertOne(context.Background(), task)
	return err
}

func (r *TaskRepository) GetTasks() ([]domain.Task, error) {
	cur, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var tasks []domain.Task
	for cur.Next(context.Background()) {
		var task domain.Task
		if err := cur.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(id primitive.ObjectID) (domain.Task, error) {
	var task domain.Task
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	return task, err
}

func (r *TaskRepository) UpdateTask(id primitive.ObjectID, updatedTask domain.Task) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": updatedTask})
	return err
}

func (r *TaskRepository) DeleteTask(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
