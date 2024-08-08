package repositories

import (
    "context"

    "github.com/kika1s1/task_manager/domain"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

// TaskRepository defines the interface for task repository methods.
type TaskRepository interface {
    CreateTask(task domain.Task) error
    GetTasks() ([]domain.Task, error)
    GetTaskByID(id primitive.ObjectID) (domain.Task, error)
    UpdateTask(id primitive.ObjectID, updatedTask domain.Task) error
    DeleteTask(id primitive.ObjectID) error
}

// taskRepository is the concrete implementation of TaskRepository.
type taskRepository struct {
    collection *mongo.Collection
}

// NewTaskRepository creates a new instance of TaskRepository.
func NewTaskRepository(client *mongo.Client) TaskRepository {
    return &taskRepository{
        collection: client.Database("task_manager").Collection("tasks"),
    }
}

func (r *taskRepository) CreateTask(task domain.Task) error {
    _, err := r.collection.InsertOne(context.Background(), task)
    return err
}

func (r *taskRepository) GetTasks() ([]domain.Task, error) {
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

func (r *taskRepository) GetTaskByID(id primitive.ObjectID) (domain.Task, error) {
    var task domain.Task
    err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
    return task, err
}

func (r *taskRepository) UpdateTask(id primitive.ObjectID, updatedTask domain.Task) error {
    _, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": updatedTask})
    return err
}

func (r *taskRepository) DeleteTask(id primitive.ObjectID) error {
    _, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
    return err
}
