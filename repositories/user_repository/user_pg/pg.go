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

func (u *userPG) SeedingAdmin() {
	newAdmin := &models.User{
		FullName: "admin",
		Email:    "admin@gmail.com",
		Password: "123456",
		Role:     "admin",
	}
	if err := u.db.Create(newAdmin).Error; err != nil {
		initialAdmin := &models.User{}

		err2 := u.db.Where("role = ?", "admin").Take(&initialAdmin).Error

		if err2 != nil { // error gabsia nge create pdhl blom ada admin
			log.Println("\n" + err.Error() + "\n")
		} else { // admin udah ada
			message := fmt.Sprintf("\nAdmin already exist with email %s\n", initialAdmin.Email)
			log.Println(message)
		}
	} else { // baru buat admin pertama kali
		log.Println("\nAdmin has been created successfully\n")
	}

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
func (u *userPG) GetUserByID(userID uint) (*models.User, errs.MessageErr) {
	user := &models.User{}
	err := u.db.Where("id = ?", userID).Take(user).Error

	if err != nil {
		message := fmt.Sprintf("User with ID %v not found", userID)
		err := errs.NewNotFound(message)
		return nil, err
	}

	return user, nil
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
