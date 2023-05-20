package user_repository

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
)

type UserRepository interface {
	RegisterUser(user *models.User) (*models.User, errs.MessageErr)
	LoginUser(user *models.User) errs.MessageErr
	UpdateUser(user *models.User) (*models.User, errs.MessageErr)
	DeleteUser(id uint) errs.MessageErr
}
