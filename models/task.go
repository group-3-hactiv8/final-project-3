package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	UserId      uint   `json:"user_id"`
	User        User
	CategoryID  uint `json:"category_id"`
	Category    Category
}

// func (sc *SocialMedia) BeforeCreate(tx *gorm.DB) error {
// 	_, err := govalidator.ValidateStruct(sc)

// 	if err != nil {
// 		return errs.NewUnprocessableEntity(err.Error())
// 	}
// 	return nil
// }
