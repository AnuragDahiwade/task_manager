package task

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/internal/db"
	"github.com/AnuragDahiwade/task-manager/internal/project"
)

var allowedStatus = map[string]bool{
	"todo":        true,
	"in-progress": true,
	"done":        true,
}

func CreateTask(c *gin.Context) {

	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	uid := c.GetString("user_id")

	if uid == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if req.ProjectID == "" {
		c.JSON(400, gin.H{"error": "project_id is required"})
		return
	}

	fmt.Println("UID =", uid)
	fmt.Println("ProjectID =", req.ProjectID)

	// Ownership check
	var proj project.Project

	if err := db.DB.
		Where("id = ? AND owner_id = ?", req.ProjectID, uid).
		First(&proj).Error; err != nil {

		c.JSON(403, gin.H{"error": "Invalid project"})
		return
	}

	task := Task{
		Title:       req.Title,
		Description: req.Description,
		ProjectID:   req.ProjectID,
		AssignedTo:  req.AssignedTo,
		DueDate:     req.DueDate,
		Priority:    req.Priority,
	}

	if err := db.DB.Create(&task).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, task)
}

func GetTasks(c *gin.Context) {

	uid := c.GetString("user_id")

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	status := c.Query("status")
	priority := c.Query("priority")
	projectID := c.Query("project_id")

	query := db.DB.
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("projects.owner_id = ?", uid)

	if status != "" {
		query = query.Where("tasks.status = ?", status)
	}

	if priority != "" {
		query = query.Where("tasks.priority = ?", priority)
	}

	if projectID != "" {
		query = query.Where("tasks.project_id = ?", projectID)
	}

	var tasks []Task

	query.
		Limit(limit).
		Offset(offset).
		Order("tasks.created_at desc").
		Find(&tasks)

	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"data":  tasks,
	})
}

func GetTaskByID(c *gin.Context) {

	id := c.Param("id")
	uid := c.GetString("user_id")

	var task Task

	err := db.DB.
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("tasks.id = ? AND projects.owner_id = ?", id, uid).
		First(&task).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Task not found or access denied",
		})
		return
	}

	c.JSON(200, task)
}

func UpdateTask(c *gin.Context) {

	id := c.Param("id")

	uid := c.GetString("user_id")

	var task Task

	err := db.DB.
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("tasks.id = ? AND projects.owner_id = ?", id, uid).
		First(&task).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	var req UpdateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if req.Status != "" && !allowedStatus[req.Status] {
		c.JSON(400, gin.H{"error": "Invalid status"})
		return
	}

	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.Status != "" {
		task.Status = req.Status
	}
	if req.Priority != "" {
		task.Priority = req.Priority
	}
	if req.AssignedTo != nil {
		task.AssignedTo = req.AssignedTo
	}
	if req.DueDate != nil {
		task.DueDate = req.DueDate
	}

	db.DB.Save(&task)

	c.JSON(200, task)
}

func DeleteTask(c *gin.Context) {

	id := c.Param("id")

	uid := c.GetString("user_id")

	err := db.DB.
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("tasks.id = ? AND projects.owner_id = ?", id, uid).
		Delete(&Task{}).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted"})
}
