package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/config"
	"github.com/AnuragDahiwade/task-manager/internal/db"
	"github.com/AnuragDahiwade/task-manager/internal/middleware"
	"github.com/AnuragDahiwade/task-manager/internal/project"
	"github.com/AnuragDahiwade/task-manager/internal/task"
	"github.com/AnuragDahiwade/task-manager/internal/user"
	"github.com/AnuragDahiwade/task-manager/routes"
)

func main() {

	// Load env
	config.LoadEnv()

	// Connect DB
	db.ConnectDB()

	// Run migrations
	db.DB.AutoMigrate(&user.User{}, &project.Project{}, &task.Task{})

	// Start Gin
	r := gin.New()

	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.RateLimit(),
	)

	// Register routes
	routes.RegisterRoutes(r)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
