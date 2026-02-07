package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/internal/db"
)

func GetAllUsers(c *gin.Context) {

	var users []User

	db.DB.Find(&users)

	// Hide passwords
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUserRole(c *gin.Context) {

	id := c.Param("id")

	var req struct {
		Role string `json:"role" binding:"required,oneof=admin user"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user User

	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	user.Role = req.Role

	db.DB.Save(&user)

	c.JSON(200, user)
}
