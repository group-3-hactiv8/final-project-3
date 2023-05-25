package category_pg

import (
	"final-project-3/models"
	"final-project-3/pkg/errs"
	"final-project-3/repositories/category_repository"
	"log"
	"fmt"

	"gorm.io/gorm"
)

type categoryPG struct {
	db *gorm.DB
}

func NewCategoryPG (db *gorm.DB) (category_repository.CategoryRepository) {
	return &categoryPG{db:db}
}

func (c *categoryPG) CreateCategory (newCategory *models.Category) (*models.Category, errs.MessageErr) {

	if err := c.db.Create(newCategory).Error; err != nil {
		log.Println(err.Error())
		message := fmt.Sprintf("Failed To Create Category with Type : %s", newCategory.Type)
		error := errs.NewInternalServerError(message)
		return nil, error
	}

	return newCategory, nil
}
