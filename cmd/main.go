package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/config"
	"github.com/AnuragDahiwade/task-manager/internal/db"
	"github.com/AnuragDahiwade/task-manager/routes"
)

func main() {

	// Load env
	config.LoadEnv()

	// Connect DB
	db.ConnectDB()

	// Start Gin
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
