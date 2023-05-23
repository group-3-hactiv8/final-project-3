package task_repository

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
)

type TaskRepository interface {
	CreateTask(task *models.Task) (*models.Task, errs.MessageErr)
	UpdateTask(task *models.Task) (*models.Task, errs.MessageErr)
}
