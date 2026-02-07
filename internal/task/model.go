package task

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string `gorm:"size:200;not null"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"size:20;default:todo"`
	Priority    string `gorm:"size:20;default:medium"`

	ProjectID  string  `gorm:"type:uuid;not null;index"`
	AssignedTo *string `gorm:"type:uuid;index"`

	DueDate *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
