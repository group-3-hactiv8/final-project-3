package models

import (
	"final-project-3/pkg/errs"
	"fmt"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null" json:"full_name"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewUnprocessableEntity(err.Error())
	}

	if u.Role == "admin" || u.Role == "member" {
		return nil
	} else {
		message := fmt.Sprintf("Invalid Role value: %s", u.Role)
		return errs.NewUnprocessableEntity(message)
	}
}
