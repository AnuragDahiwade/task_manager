package routes

import (
	"github.com/AnuragDahiwade/task-manager/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")

	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})

	authGroup := api.Group("/auth")

	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)
}
