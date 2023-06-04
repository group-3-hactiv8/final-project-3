package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Status      bool   `json:"status"`
	UserId      uint   `json:"user_id"`
	User        User
	CategoryId  uint `json:"category_id"`
	Category    Category
}
