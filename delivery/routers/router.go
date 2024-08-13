package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/delivery/controllers"
	"github.com/kika1s1/task_manager/infrastructure"
	"github.com/kika1s1/task_manager/repositories"
	"github.com/kika1s1/task_manager/usecases" // Update the import statement to match the correct case
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(client)
	taskRepo := repositories.NewTaskRepository(client)

	// Initialize use cases
	userUsecase := usecases.NewUserUsecase(*userRepo)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	// Initialize controller
	ctrl := controllers.NewController(userUsecase, taskUsecase)

	// Public routes
	r.POST("/auth/register", ctrl.Register)
	r.POST("/auth/login", ctrl.Login)

	// Auth middleware
	authMiddleware := infrastructure.AuthMiddleware()

	// Task routes
	r.GET("/tasks", authMiddleware, ctrl.GetTasks)
	r.GET("/tasks/:id", authMiddleware, ctrl.GetTaskByID)

	// Admin routes
	adminMiddleware := infrastructure.AdminMiddleware()
	r.PUT("/promote/:username", authMiddleware, adminMiddleware, ctrl.Promote)
	r.DELETE("/tasks/:id", authMiddleware, adminMiddleware, ctrl.DeleteTask)
	r.POST("/tasks", authMiddleware, adminMiddleware, ctrl.CreateTask)
	r.PUT("/tasks/:id", authMiddleware, adminMiddleware, ctrl.UpdateTask)

	return r
}
