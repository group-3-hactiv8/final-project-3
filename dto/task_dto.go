package dto

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type NewTaskRequest struct {
	Title       string `json:"title" valid:"required~Your Title is required"`
	Description string `json:"description" valid:"required~Your Description is required"`
	CategoryId  uint   `json:"category_id" valid:"required~Your Category ID is required"`
}

func (t *NewTaskRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (t *NewTaskRequest) TaskRequestToModel() *models.Task {
	return &models.Task{
		Title:       t.Title,
		Description: t.Description,
		CategoryId:  t.CategoryId,
		Status:      false, // hardcoded. Every new task has a false status.
	}
}

type NewTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      uint      `json:"user_id"`
	CategoryId  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateStatusOfATaskRequest struct {
	Status bool `json:"status" valid:"required~Your Status is required"`
}

func (t *UpdateStatusOfATaskRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (t *UpdateStatusOfATaskRequest) TaskRequestToModel() *models.Task {
	return &models.Task{
		Status: t.Status,
	}
}

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (t *UpdateTaskRequest)  TaskRequestToModel() *models.Task {
	return &models.Task{
		Title:       t.Title,
		Description: t.Description,
	}
}


type GetAllTasksResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	User        UserData  `json:"user"`
}
type UpdateTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      uint      `json:"user_id"`
	CategoryId  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateCategoryIdOfATasIdkRequest struct {
	CategoryId  uint   `json:"category_id" valid:"required~Your Category ID is required"`
}

func (t *UpdateCategoryIdOfATasIdkRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (t *UpdateCategoryIdOfATasIdkRequest) TaskRequestCategoryToModel() *models.Task {
	return &models.Task{
		CategoryId: t.CategoryId,
	}
}

type UpdateCategoryIdOfTaskIdResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      uint      `json:"user_id"`
	CategoryId  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type DeleteTaskResponse struct {
	Message string `json:"message"`
}
