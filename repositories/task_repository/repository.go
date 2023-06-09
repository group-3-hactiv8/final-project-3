package task_repository

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
)

type TaskRepository interface {
	CreateTask(task *models.Task) (*models.Task, errs.MessageErr)
	GetAllTasks() ([]models.Task, errs.MessageErr)
	GetAllTasksByCategoryID(categoryId uint) ([]models.Task, errs.MessageErr)
	GetTaskByID(id uint) (*models.Task, errs.MessageErr)
	UpdateTask(task *models.Task) (*models.Task, errs.MessageErr)
	UpdateTaskStatus(id uint, newStatus bool) (*models.Task, errs.MessageErr)
	UpdateCategoryIdOfTask(task *models.Task) (*models.Task, errs.MessageErr)
	DeleteTask(id uint) errs.MessageErr
}


// GetTaskByID(id uint) (*models.Task, errs.MessageErr)
// GetTaskByID(userId uint) ([]models.Task, errs.MessageErr)