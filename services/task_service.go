package services

import (
	"final-project-3/dto"
	"final-project-3/pkg/errs"
	"final-project-3/repositories/category_repository"
	"final-project-3/repositories/task_repository"
)

type TaskService interface {
	CreateTask(payload *dto.NewTaskRequest, userId uint) (*dto.NewTaskResponse, errs.MessageErr)
	UpdateStatus(id int, payload *dto.UpdateStatusOfATaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateCategoryIdOfTask(id int, payload *dto.UpdateCategoryIdOfATasIdkRequest) (*dto.UpdateCategoryIdOfTaskIdResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo     task_repository.TaskRepository
	categoryRepo category_repository.CategoryRepository
}

func NewTaskService(taskRepo task_repository.TaskRepository, categoryRepo category_repository.CategoryRepository) TaskService {
	return &taskService{taskRepo: taskRepo, categoryRepo: categoryRepo}
}

func (t *taskService) CreateTask(payload *dto.NewTaskRequest, userId uint) (*dto.NewTaskResponse, errs.MessageErr) {
	task := payload.TaskRequestToModel()
	task.UserId = userId

	createdTask, err := t.taskRepo.CreateTask(task)
	if err != nil {
		return nil, err
	}

	response := &dto.NewTaskResponse{
		ID:          createdTask.ID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		Status:      createdTask.Status,
		UserId:      createdTask.UserId,
		CategoryId:  createdTask.CategoryId,
		CreatedAt:   createdTask.CreatedAt,
	}

	return response, nil
}

func (t *taskService) UpdateStatus(id int, payload *dto.UpdateStatusOfATaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	task := payload.TaskRequestToModel()
	if id < 1 {
		idError := errs.NewBadRequest("Task ID value must be positive")
		return nil, idError
	}
	task.ID = uint(id)

	updatedTask, err := t.taskRepo.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateTaskResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserId:      updatedTask.UserId,
		CategoryId:  updatedTask.CategoryId,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) UpdateCategoryIdOfTask(id int, payload *dto.UpdateCategoryIdOfATasIdkRequest) (*dto.UpdateCategoryIdOfTaskIdResponse, errs.MessageErr) {
	task := payload.TaskRequestCategoryToModel()
	if id < 1 {
		idError := errs.NewBadRequest("Task ID value must be positive")
		return nil, idError
	}
	task.ID = uint(id)

	_, err := t.categoryRepo.GetCategoryById(task.CategoryId)
	if err != nil {
		return nil, err
	}

	updatedTask, err := t.taskRepo.UpdateCategoryIdOfTask(task)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateCategoryIdOfTaskIdResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserId:      updatedTask.UserId,
		CategoryId:  updatedTask.CategoryId,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return response, nil
}
