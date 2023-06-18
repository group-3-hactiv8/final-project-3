package category_repository

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) (*models.Category, errs.MessageErr)
	GetAllCategory() ([]models.Category, errs.MessageErr)
	GetTasksByCategoryID(categoryId uint) ([]models.Task, errs.MessageErr)
	GetCategoryById(id uint) (*models.Category, errs.MessageErr)
	UpdateCategory(category *models.Category, ctyUpdate *models.Category) (*models.Category, errs.MessageErr)
	DeleteCategory(category *models.Category) errs.MessageErr
}