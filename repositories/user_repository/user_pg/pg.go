package user_pg

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
	"final-project-3/repositories/user_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) RegisterUser(newUser *models.User) (*models.User, errs.MessageErr) {
	if err := u.db.Create(newUser).Error; err != nil {
		log.Println(err.Error())
		message := fmt.Sprintf("Failed to register a new user with Email %s", newUser.Email)
		error := errs.NewInternalServerError(message)
		return nil, error
	}

	return newUser, nil
}

func (u *userPG) GetUserByEmail(user *models.User) errs.MessageErr {
	err := u.db.Where("email = ?", user.Email).Take(&user).Error
	// Karna di Take, objek user akan terupdate, termasuk passwordnya.
	// Makannya kita simpen dulu password dari request nya di service level.

	if err != nil {
		err2 := errs.NewBadRequest("Wrong email/password")
		return err2
	}

	return nil
}

func (u *userPG) UpdateUser(user *models.User) (*models.User, errs.MessageErr) {
	err := u.db.Model(user).Updates(user).Error

	if err != nil {
		err2 := errs.NewBadRequest(err.Error())
		return nil, err2
	}

	return user, nil
}

func (u *userPG) DeleteUser(id uint) errs.MessageErr {
	err := u.db.Delete(&models.User{}, id).Error

	if err != nil {
		err3 := errs.NewInternalServerError(err.Error())
		return err3
	}

	return nil
}
