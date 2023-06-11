package http_handlers

import (
	"final-project-3/dto"
	"final-project-3/pkg/errs"
	"final-project-3/services"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService: categoryService}
}

// CreateCategory godoc
//
//	@Summary		Register a category
//	@Description	Register a category by json
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			category  body      dto.NewCategoryRequest	true	"Create Category request body"
//	@Success		201		{object}	dto.NewCategoryResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/category [post]
func (c *categoryHandler) CreateCategory(ctx *gin.Context) {
	var requestBody dto.NewCategoryRequest

	err := ctx.ShouldBindJSON(&requestBody)

	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	newCtyResponse, err3 := c.categoryService.CreateCategory(&requestBody) 

	if err3 != nil {
		ctx.JSON(err2.StatusCode(), err3)
	}

	ctx.JSON(http.StatusOK, newCtyResponse)
}

// GetAllCategory godoc
//
//	@Summary		Get all category
//	@Description	Get all products by json
//	@Tags			products
//	@Produce		json
//	@Success		200		{object}	dto.AllProductsResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/products [get]
func (c *categoryHandler) GetAllCategory(ctx *gin.Context) {
	allCategories, err := c.categoryService.GetAllCategory()
	if err != nil {
		ctx.JSON(err.StatusCode(),err)
		return
	}
	ctx.JSON(http.StatusOK, allCategories)
}

// UpdateCategory godoc
//
//	@Summary		Update a category 
//	@Description	Update a category 
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			category		body		dto.UpdateCategoryRequest	true	"Update category request body"
//	@Param			categoryId		path		uint						true	"category ID request"
//	@Success		200				{object}	dto.UpdateCategoryResponse
//	@Failure		401				{object}	errs.MessageErrData
//	@Failure		400				{object}	errs.MessageErrData
//	@Failure		422				{object}	errs.MessageErrData
//	@Failure		500				{object}	errs.MessageErrData
//	@Router			/category/{categoryId} [patch]
func (c *categoryHandler) UpdateCategory(ctx *gin.Context) {
	
	ctyId, err := strconv.Atoi(ctx.Param("categoryId"))
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	var requestBody dto.NewCategoryRequest

	err = ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	updatedCtyResponse, err3 := c.categoryService.UpdateCategory(uint(ctyId), &requestBody)

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, updatedCtyResponse)
}

// DeleteCategory godoc
//
//	@Summary		Delete a Category
//	@Description	Delete a Category by param
//	@Tags			category
//	@Produce		json
//	@Param			categoryId	 path		uint						true	"Category ID request"
//	@Success		200			{object}	dto.DeleteCategoryResponse
//	@Failure		401			{object}	errs.MessageErrData
//	@Failure		400			{object}	errs.MessageErrData
//	@Router			/category/{id} [delete]
func (c *categoryHandler) DeleteCategory(ctx *gin.Context) {
	ctyId := ctx.Param("categoryId")
	ctyIdUint, err := strconv.ParseUint(ctyId, 10, 16)
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	response, err2 := c.categoryService.DeleteCategory(uint(ctyIdUint))

	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}