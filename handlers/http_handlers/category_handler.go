package http_handlers

import (
	"final-project-3/dto"
	"final-project-3/pkg/errs"
	"final-project-3/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService: categoryService}
}

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