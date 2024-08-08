package usecases

import (
    "github.com/gin-gonic/gin"
    "github.com/kika1s1/task_manager/domain"
    "github.com/kika1s1/task_manager/repositories"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "net/http"
)

// TaskUsecase defines the interface for task use cases.
type TaskUsecase interface {
    CreateTask(c *gin.Context)
    GetTasks(c *gin.Context)
    GetTaskByID(c *gin.Context)
    UpdateTask(c *gin.Context)
    DeleteTask(c *gin.Context)
}

// taskUsecase is the concrete implementation of TaskUsecase.
type taskUsecase struct {
    taskRepo repositories.TaskRepository
}

// NewTaskUsecase creates a new instance of TaskUsecase.
func NewTaskUsecase(taskRepo repositories.TaskRepository) TaskUsecase {
    return &taskUsecase{taskRepo: taskRepo}
}

func (u *taskUsecase) CreateTask(c *gin.Context) {
    var task domain.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := u.taskRepo.CreateTask(task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (u *taskUsecase) GetTasks(c *gin.Context) {
    tasks, err := u.taskRepo.GetTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func (u *taskUsecase) GetTaskByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    task, err := u.taskRepo.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func (u *taskUsecase) UpdateTask(c *gin.Context) {
    idParam := c.Param("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var updatedTask domain.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := u.taskRepo.UpdateTask(id, updatedTask); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (u *taskUsecase) DeleteTask(c *gin.Context) {
    idParam := c.Param("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := u.taskRepo.DeleteTask(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
