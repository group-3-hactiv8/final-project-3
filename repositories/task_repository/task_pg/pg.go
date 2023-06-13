package task_pg

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
	"final-project-3/repositories/task_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type taskPG struct {
	db *gorm.DB
}

func NewTaskPG(db *gorm.DB) task_repository.TaskRepository {
	return &taskPG{db: db}
}

func (t *taskPG) CreateTask(task *models.Task) (*models.Task, errs.MessageErr) {
	if err := t.db.Create(task).Error; err != nil {
		log.Println(err.Error())
		message := fmt.Sprintf("Failed to register a new task with title %s", task.Title)
		error := errs.NewInternalServerError(message)
		return nil, error
	}

	return task, nil
}

func (t *taskPG) UpdateTask(task *models.Task) (*models.Task, errs.MessageErr) {
	err := t.db.Model(task).Updates(task).Error

	if err != nil {
		err2 := errs.NewBadRequest(err.Error())
		return nil, err2
	}

	err = t.db.Where("id = ?", task.ID).Take(&task).Error
	if err != nil {
		err2 := errs.NewInternalServerError(err.Error())
		return nil, err2
	}

	return task, nil
}

func (t *taskPG) UpdateCategoryIdOfTask(task *models.Task) (*models.Task, errs.MessageErr) {
	err := t.db.Model(task).Updates(task).Error

	if err != nil {
		err2 := errs.NewBadRequest(err.Error())
		return nil, err2
	}

	err = t.db.Where("id = ?", task.ID).Take(&task).Error
	if err != nil {
		err2 := errs.NewInternalServerError(err.Error())
		return nil, err2
	}

	return task, nil
}
