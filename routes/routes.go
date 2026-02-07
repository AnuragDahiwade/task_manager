package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/internal/auth"
	"github.com/AnuragDahiwade/task-manager/internal/middleware"
	"github.com/AnuragDahiwade/task-manager/internal/project"
	"github.com/AnuragDahiwade/task-manager/internal/task"
	"github.com/AnuragDahiwade/task-manager/internal/user"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Auth
	authGroup := api.Group("/auth")
	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// Normal
	protected.GET("/me", user.GetMe)

	// Admin
	admin := protected.Group("/admin")
	admin.Use(middleware.AdminOnly())

	admin.GET("/users", user.GetAllUsers)
	admin.PUT("/users/:id/role", user.UpdateUserRole)

	// project routes
	projectGroup := protected.Group("/projects")

	projectGroup.POST("", project.CreateProject)
	projectGroup.GET("", project.GetMyProjects)
	projectGroup.GET("/:id", project.GetProjectByID)
	projectGroup.PUT("/:id", project.UpdateProject)
	projectGroup.DELETE("/:id", project.DeleteProject)

	// task routes
	taskGroup := protected.Group("/tasks")

	taskGroup.POST("", task.CreateTask)
	taskGroup.GET("", task.GetTasks)
	taskGroup.GET("/:id", task.GetTaskByID)
	taskGroup.PUT("/:id", task.UpdateTask)
	taskGroup.DELETE("/:id", task.DeleteTask)

}
