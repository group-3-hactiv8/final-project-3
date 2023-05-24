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
	CreatedAt time.Time `json:"date"`
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