package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/internal/db"
	"github.com/AnuragDahiwade/task-manager/internal/user"

	"github.com/AnuragDahiwade/task-manager/pkg/utils"
)

func Register(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check existing user
	var existing user.User
	if err := db.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Password hashing failed"})
		return
	}

	newUser := user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
	}

	db.DB.Create(&newUser)

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user user.User

	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
