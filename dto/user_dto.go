package dto

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type NewUserRequest struct {
	FullName string `json:"full_name" valid:"required~Your Full Name is required"`
	Email    string `json:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `json:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

func (u *NewUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (u *NewUserRequest) UserRequestToModel() *models.User {
	return &models.User{
		FullName: u.FullName,
		Email:    u.Email,
		Password: u.Password,
	}
}

type NewUserResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	Email    string `json:"email" valid:"required~Your Email is required"`
	Password string `json:"password" valid:"required~Your password is required"`
}

func (u *LoginUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (u *LoginUserRequest) LoginUserRequestToModel() *models.User {
	return &models.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" valid:"required~Your Full Name is required"`
	Email    string `json:"email" valid:"required~Your email is required, email~Invalid email format"`
}

func (u *UpdateUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (u *UpdateUserRequest) UpdateUserRequestToModel() *models.User {
	return &models.User{
		FullName: u.FullName,
		Email:    u.Email,
	}
}

type UpdateUserResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
