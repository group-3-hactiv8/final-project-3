package dto

// import (
// 	"final-project-3/models"
// 	"final-project-3/pkg/errs"
// 	"time"

// 	"github.com/asaskevich/govalidator"
// )

type NewUserRequest struct {
}

// func (u *NewUserRequest) ValidateStruct() errs.MessageErr {
// 	_, err := govalidator.ValidateStruct(u)

// 	if err != nil {
// 		return errs.NewBadRequest(err.Error())
// 	}

// 	isAbove8 := govalidator.InRangeInt(u.Age, 9, MaxInt) // kalau 9 tetep true (lower bound)

// 	if !isAbove8 {
// 		return errs.NewUnprocessableEntity("Age must have value above 8")
// 	}

// 	return nil

// }

//	func (u *NewUserRequest) UserRequestToModel() *models.User {
//		return &models.User{
//			Username: u.Username,
//			Email:    u.Email,
//			Password: u.Password,
//			Age:      u.Age,
//		}
//	}
type NewUserResponse struct {
}

type LoginUserRequest struct {
}

type LoginUserResponse struct {
}

type UpdateUserRequest struct {
}

type UpdateUserResponse struct {
}

type DeleteUserResponse struct {
}
