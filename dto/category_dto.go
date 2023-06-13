package dto

import (
	"time"
	"final-project-3/models"
	"final-project-3/pkg/errs"
	"github.com/asaskevich/govalidator"
)

type NewCategoryRequest struct {
	Type string `json:"type" valid:"required~Your Type is required"`
}

type NewCategoryResponse struct {
	ID   int    		`json:"id"`
	Type string 		`json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

func (newCty *NewCategoryRequest) NewCategoryRequestToModel() *models.Category {
	return &models.Category{
		Type: newCty.Type,
	}
}

func (newCty *NewCategoryRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(newCty)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

type UpdateCategoryRequest struct {
	Type string `json:"type" valid:"required~Your Type is required"`
}

func (newCty *UpdateCategoryRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(newCty)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

type UpdateCategoryResponse struct {
	ID   uint    		`json:"id"`
	Type string 		`json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllCategoryResponse struct {
	ID        uint       `json:"id"`
	Type      string     `json:"type"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	Tasks     []TaskForGetAllCategoryResponse `json:"Tasks"`
}

type TaskForGetAllCategoryResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      uint      `json:"user_id"`
	CategoryId  uint      `json:"category_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DeleteCategoryResponse struct {
	Message string `json:"message"`
}