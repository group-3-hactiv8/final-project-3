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

func (t *taskPG) GetAllTasks() ([]models.Task, errs.MessageErr) {
	var tasks []models.Task
	err := t.db.Find(&tasks).Error

	if err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewInternalServerError("Failed to get all Task")
		return nil, error
	}

	return tasks, nil
}

func (t *taskPG) GetAllTasksByCategoryID(categoryID uint) (
	[]models.Task, errs.MessageErr,
) {
	var tasks []models.Task
	if err := t.db.Find(&tasks, "category_id = ?", categoryID).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to geet all task")
	}

	return tasks, nil
}

func (t *taskPG) GetTaskByID(id uint) (*models.Task, errs.MessageErr) {
	var task models.Task

	if err := t.db.First(&task, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewNotFound(fmt.Sprintf("Task with %d is not found", id))
	}

	return &task, nil
}

func (t *taskPG) UpdateTask(task *models.Task) (*models.Task, errs.MessageErr) {
	err := t.db.Model(task).Updates(task).Error

	if err != nil {
		err2 := errs.NewBadRequest(err.Error())
		return nil, err2
	}

	return task, nil
}


func (t *taskPG) UpdateTaskStatus(id uint, newStatus bool) (*models.Task, errs.MessageErr) {
	task, err := t.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Status = newStatus

	if err := t.db.Save(task).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to update task status")
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

func (t *taskPG) DeleteTask(id uint) errs.MessageErr {
	if err := t.db.Delete(&models.Task{}, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete Task with id %d", id))
	}

	return nil
}
