package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/usecases"
)

// Controller defines the interface for the task and user controllers.
type Controller interface {
    Register(c *gin.Context)
    Login(c *gin.Context)
    CreateTask(c *gin.Context)
    GetTasks(c *gin.Context)
    GetTaskByID(c *gin.Context)
    UpdateTask(c *gin.Context)
    DeleteTask(c *gin.Context)
    Promote(c *gin.Context)
}

// controller is the concrete implementation of Controller.
type controller struct {
    UserUsecase usecases.UserUsecase
    TaskUsecase usecases.TaskUsecase
}

// NewController creates a new instance of Controller.
func NewController(userUsecase usecases.UserUsecase, taskUsecase usecases.TaskUsecase) Controller {
    return &controller{
        UserUsecase: userUsecase,
        TaskUsecase: taskUsecase,
    }
}

func (ctrl *controller) Register(c *gin.Context) {
    ctrl.UserUsecase.Register(c)
}

func (ctrl *controller) Login(c *gin.Context) {
    ctrl.UserUsecase.Login(c)
}

func (ctrl *controller) CreateTask(c *gin.Context) {
    ctrl.TaskUsecase.CreateTask(c)
}

func (ctrl *controller) GetTasks(c *gin.Context) {
    ctrl.TaskUsecase.GetTasks(c)
}

func (ctrl *controller) GetTaskByID(c *gin.Context) {
    ctrl.TaskUsecase.GetTaskByID(c)
}

func (ctrl *controller) UpdateTask(c *gin.Context) {
    ctrl.TaskUsecase.UpdateTask(c)
}

func (ctrl *controller) DeleteTask(c *gin.Context) {
    ctrl.TaskUsecase.DeleteTask(c)
}

func (ctrl *controller) Promote(c *gin.Context) {
    ctrl.UserUsecase.Promote(c)
}
