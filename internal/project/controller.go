package project

import (
	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/internal/db"
)

func CreateProject(c *gin.Context) {

	var req CreateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	project := Project{
		Name:    req.Name,
		OwnerID: userID.(string),
	}

	db.DB.Create(&project)

	c.JSON(201, project)
}

func GetMyProjects(c *gin.Context) {

	userID, _ := c.Get("user_id")

	var projects []Project

	db.DB.Where("owner_id = ?", userID).
		Order("created_at desc").
		Find(&projects)

	c.JSON(200, projects)
}

func GetProjectByID(c *gin.Context) {

	id := c.Param("id")

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	uid := userID.(string)

	var project Project

	err := db.DB.
		Where("id = ? AND owner_id = ?", id, uid).
		First(&project).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Project not found or access denied",
		})
		return
	}

	c.JSON(200, project)
}

func UpdateProject(c *gin.Context) {

	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var project Project

	if err := db.DB.First(&project, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Project not found"})
		return
	}

	// Ownership check
	if project.OwnerID != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	var req UpdateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	project.Name = req.Name

	db.DB.Save(&project)

	c.JSON(200, project)
}

func DeleteProject(c *gin.Context) {

	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var project Project

	if err := db.DB.First(&project, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Project not found"})
		return
	}

	if project.OwnerID != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	db.DB.Delete(&project)

	c.JSON(200, gin.H{"message": "Deleted"})
}
