package category_repository

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) (*models.Category, errs.MessageErr)
}