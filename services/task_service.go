package services

import (
	"final-project-3/dto"
	"final-project-3/pkg/errs"
	"final-project-3/repositories/category_repository"
	"final-project-3/repositories/task_repository"
	"final-project-3/repositories/user_repository"
	"fmt"
)

type TaskService interface {
	CreateTask(payload *dto.NewTaskRequest, userId uint) (*dto.NewTaskResponse, errs.MessageErr)
	GetAllTasks() ([]dto.GetAllTasksResponse, errs.MessageErr)
	UpdateTask(id uint, payload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateStatus(id int, payload *dto.UpdateStatusOfATaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateCategoryIdOfTask(id uint, payload *dto.UpdateCategoryIdOfATasIdkRequest) (*dto.UpdateCategoryIdOfTaskIdResponse, errs.MessageErr)
	DeleteTask(id uint) (*dto.DeleteTaskResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo     task_repository.TaskRepository
	categoryRepo category_repository.CategoryRepository
	userRepo     user_repository.UserRepository
}

func NewTaskService(taskRepo task_repository.TaskRepository, categoryRepo category_repository.CategoryRepository, userRepo user_repository.UserRepository) TaskService {
	return &taskService{taskRepo, categoryRepo, userRepo}
}

func (t *taskService) CreateTask(payload *dto.NewTaskRequest, userId uint) (*dto.NewTaskResponse, errs.MessageErr) {
	task := payload.TaskRequestToModel()
	task.UserId = userId

	_, err := t.categoryRepo.GetCategoryById(task.CategoryId)
	if err != nil {
		errNotFound := errs.NewNotFound(fmt.Sprintf("Category with ID %v not found", task.CategoryId))
		return nil, errNotFound
	}

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


func (t *taskService) GetAllTasks() ([]dto.GetAllTasksResponse, errs.MessageErr) {
	tasks, err := t.taskRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllTasksResponse{}
	for _, task := range tasks {
		user, _ := t.userRepo.GetUserByID(task.UserId) 
		response = append(response, dto.GetAllTasksResponse{
			ID:          task.CategoryId,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			UserID:      task.UserId,
			CategoryID:  task.CategoryId,
			CreatedAt:   task.CreatedAt,
			User: dto.UserData{
				ID:       user.ID,
				Email:    user.Email,
				FullName: user.FullName,
			},
		})
	}

	return response, nil
}



func (t *taskService) UpdateTask(id uint, payload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.MessageErr) {
	taskUpdateRequest := payload.UpdateTaskRequestToModel()

	taskUpdateRequest.ID = id

	updatedTask, err := t.taskRepo.UpdateTask(taskUpdateRequest)
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

func (t *taskService) UpdateCategoryIdOfTask(id uint, payload *dto.UpdateCategoryIdOfATasIdkRequest) (*dto.UpdateCategoryIdOfTaskIdResponse, errs.MessageErr) {
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

func (t *taskService) DeleteTask(id uint) (*dto.DeleteTaskResponse, errs.MessageErr) {
	task, err := t.taskRepo.GetTaskByID(id)

	if err != nil {
		return nil, err
	}

	if err := t.taskRepo.DeleteTask(task.ID); err != nil {
		return nil, err
	}

	deleteResponse := &dto.DeleteTaskResponse{
		Message: "Task has been successfully deleted",
	}
	return deleteResponse, nil
}




