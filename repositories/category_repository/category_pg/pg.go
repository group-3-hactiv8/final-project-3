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

func (c *categoryPG) GetCategoryById(id uint) (*models.Category, errs.MessageErr){
	var category models.Category
	result := c.db.First(&category, id)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewNotFound(fmt.Sprintf("failed to get category by id :", category.ID))
		return nil, error
	}
	return &category, nil
}

func (c *categoryPG) UpdateCategory(category *models.Category, ctyUpdate *models.Category) (*models.Category, errs.MessageErr) {
	err := c.db.Model(category).Updates(ctyUpdate).Error
	if err != nil {
		message := fmt.Sprintf("Failed to Update Category with Id : %v", category.ID)
		err2 := errs.NewNotFound(message)
		return nil, err2
	}
	return category, nil
}

func (c *categoryPG) DeleteCategory(category *models.Category) errs.MessageErr {
	result := c.db.Delete(category)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewInternalServerError(fmt.Sprintf("Failed to delete Category by id : %v", category.ID))
		return error
	}
	return  nil
}
func (r *categoryPG) GetAllCategory() ([]models.Category, errs.MessageErr) {
	var categories []models.Category
	err := r.db.Find(&categories).Error

	if err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewInternalServerError("Failed to fetching all Category")
		return nil, error
	}

	return categories, nil
}

func (r *categoryPG) GetTasksByCategoryID(categoryId uint) ([]models.Task, errs.MessageErr) {
	var tasks []models.Task
	err := r.db.Where("category_id = ?", categoryId).Find(&tasks).Error

	if err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewNotFound(fmt.Sprintf("failed to get task by id :", categoryId))
		return nil, error
	}

	return tasks, nil
}