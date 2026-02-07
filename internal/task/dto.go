package task

import "time"

type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required,min=3"`
	Description string     `json:"description"`
	ProjectID   string     `json:"project_id" binding:"required"`
	AssignedTo  *string    `json:"assigned_to"`
	DueDate     *time.Time `json:"due_date"`
	Priority    string     `json:"priority"`
}

type UpdateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	AssignedTo  *string    `json:"assigned_to"`
	DueDate     *time.Time `json:"due_date"`
}
