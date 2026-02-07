package user

import (
	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/internal/db"
)

func GetMe(c *gin.Context) {

	userID, _ := c.Get("user_id")

	var user User

	if err := db.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Hide password
	user.Password = ""

	c.JSON(200, user)
}
