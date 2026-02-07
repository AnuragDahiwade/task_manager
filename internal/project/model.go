package project

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string `gorm:"size:150;not null"`
	OwnerID   string `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
