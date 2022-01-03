package models

import (
	"time"

	"gorm.io/gorm"
)

type Character struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	History 	string         `json:"description"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}