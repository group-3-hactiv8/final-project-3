package services

import (
	"final-project-3/dto"
	"final-project-3/pkg/errs"
	"final-project-3/repositories/category_repository"
)

type CategoryService interface {
	CreateCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
	GetAllCategory() ([]dto.GetAllCategoryResponse, errs.MessageErr)
	UpdateCategory(id uint, payload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr)
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

func (c *categoryService) UpdateCategory(id uint, payload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)  {
	category, err := c.categoryRepo.GetCategoryById(id)

	if id < 1 {
		idError := errs.NewBadRequest("ID value must be positive")
		return nil, idError
	}

	if err != nil {
		return nil, err
	}

	updateNewCategory := payload.NewCategoryRequestToModel()

	updateCategory, err2 := c.categoryRepo.UpdateCategory(category, updateNewCategory)

	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateCategoryResponse {
		ID: updateCategory.ID,
		Type: updateCategory.Type,
		UpdatedAt: category.UpdatedAt,
	}
	return response, nil
}

func (c *categoryService) DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	category, err := c.categoryRepo.GetCategoryById(id)

	if err != nil {
		return nil, err
	}

	if err := c.categoryRepo.DeleteCategory(category); err != nil {
		return nil, err
	}

	deleteResponse := &dto.DeleteCategoryResponse{
		Message: "Category has been successfully deleted",
	}
	return deleteResponse, nil
}

func (c *categoryService) GetAllCategory() ([]dto.GetAllCategoryResponse, errs.MessageErr) {
	allCategories, err := c.categoryRepo.GetAllCategory()

	if err != nil {
		return nil, err
	}

	response := []dto.GetAllCategoryResponse{}
	for _, cty := range allCategories {
		tasks, err := c.categoryRepo.GetTasksByCategoryID(cty.ID)
		if err != nil {
			return nil, err
		}

		taskResponses := []dto.TaskForGetAllCategoryResponse{}
		for _, task := range tasks {
			taskResponses = append(taskResponses, dto.TaskForGetAllCategoryResponse{
				ID:          task.ID,
				Title:       task.Title,
				Description: task.Description,
				UserId:      task.UserId,
				CategoryId:  task.CategoryId,
				CreatedAt:   task.CreatedAt,
				UpdatedAt:   task.UpdatedAt,
			})
		}

		response = append(response, dto.GetAllCategoryResponse{
			ID:        cty.ID,
			Type:      cty.Type,
			UpdatedAt: cty.UpdatedAt,
			CreatedAt: cty.CreatedAt,
			Tasks:     taskResponses,
		})
	}

	return response, nil
}