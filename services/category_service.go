package services

import (
	"final-project-3/dto"
	// "final-project-3/helpers"
	"final-project-3/pkg/errs"
	"final-project-3/repositories/category_repository"
)

type CategoryService interface {
	CreateCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo category_repository.CategoryRepository
}

func NewCategoryService(categoryRepo category_repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
} 

func (c *categoryService) CreateCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr) {
	newCty := payload.NewCategoryRequestToModel()

	createCategory, err := c.categoryRepo.CreateCategory(newCty)
	if err != nil {
		return nil, err
	}

	response := &dto.NewCategoryResponse {
		ID: int(createCategory.ID),
		Type: createCategory.Type,
		CreatedAt: createCategory.CreatedAt,
	}
	return response, nil
}