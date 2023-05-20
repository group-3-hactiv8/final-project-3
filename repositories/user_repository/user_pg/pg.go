package user_pg

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
	"final-project-3/repositories/user_repository"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) RegisterUser(newUser *models.User) (*models.User, errs.MessageErr) {

}

func (u *userPG) LoginUser(user *models.User) errs.MessageErr {

}

func (u *userPG) UpdateUser(user *models.User) (*models.User, errs.MessageErr) {

}

func (u *userPG) DeleteUser(id uint) errs.MessageErr {

}
