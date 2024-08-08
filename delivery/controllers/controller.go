package controllers

import (
	"github.com/gin-gonic/gin"
	 "github.com/kika1s1/task_manager/usecases"
)

type Controller struct {
	UserUsecase usecases.UserUsecase
	TaskUsecase usecases.TaskUsecase
}

func NewController(userUsecase usecases.UserUsecase, taskUsecase usecases.TaskUsecase) *Controller {
	return &Controller{
		UserUsecase: userUsecase,
		TaskUsecase: taskUsecase,
	}
}

func (ctrl *Controller) Register(c *gin.Context) {
	ctrl.UserUsecase.Register(c)
}

func (ctrl *Controller) Login(c *gin.Context) {
	ctrl.UserUsecase.Login(c)
}

func (ctrl *Controller) CreateTask(c *gin.Context) {
	ctrl.TaskUsecase.CreateTask(c)
}

func (ctrl *Controller) GetTasks(c *gin.Context) {
	ctrl.TaskUsecase.GetTasks(c)
}

func (ctrl *Controller) GetTaskByID(c *gin.Context) {
	ctrl.TaskUsecase.GetTaskByID(c)
}

func (ctrl *Controller) UpdateTask(c *gin.Context) {
	ctrl.TaskUsecase.UpdateTask(c)
}

func (ctrl *Controller) DeleteTask(c *gin.Context) {
	ctrl.TaskUsecase.DeleteTask(c)
}

func (ctrl *Controller) Promote(c *gin.Context) {
	ctrl.UserUsecase.Promote(c)
}
