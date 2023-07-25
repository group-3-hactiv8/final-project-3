package http_handlers

import (
	"final-project-3/dto"
	"final-project-3/pkg/errs"
	"final-project-3/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *taskHandler {
	return &taskHandler{taskService: taskService}
}

// CreateTask godoc
//
//	@Summary		Register a task
//	@Description	Register a task by json
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		dto.NewTaskRequest	true	"Create task request body"
//	@Success		201		{object}	dto.NewTaskResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/tasks [post]
func (t *taskHandler) CreateTask(ctx *gin.Context) {
	var requestBody dto.NewTaskRequest

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

	// mustget = ambil data dari middleware authentication.
	// Tp hasil returnnya hanya empty interface, jadi harus
	// di cast dulu ke jwt.MapClaims.
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	createdTaskResponse, err3 := t.taskService.CreateTask(&requestBody, userId)
	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusCreated, createdTaskResponse)

}

// GetAllTask godoc
//
//	@Summary		Get all task
//	@Description	Get all task by json
//	@Tags			tasks
//	@Produce		json
//	@Success		200		{object}	dto.GetAllTasksResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/tasks [get]
func (t *taskHandler) GetAllTasks(ctx *gin.Context) {
	tasks, err := t.taskService.GetAllTasks()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// UpdateTask godoc
//
//	@Summary		Task a user
//	@Description	Task a user by json
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		dto.UpdateTaskRequest	true	"Update a task request body"
//	@Param			taskId	path		uint							true	"Task ID request"
//	@Success		200		{object}	dto.UpdateTaskResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/tasks/{taskId}	 [put]
func (t *taskHandler) UpdateTask(ctx *gin.Context) {
	taskId, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	var requestBody dto.UpdateTaskRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	updatedTaskResponse, err3 := t.taskService.UpdateTask(uint(taskId), &requestBody)
	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, updatedTaskResponse)
}

// UpdateStatus godoc
//
//	@Summary		Update status of a task
//	@Description	Update status of a task by json
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		dto.UpdateStatusOfATaskRequest	true	"Update status of task request body"
//	@Param			taskId	path		uint							true	"Task ID request"
//	@Success		200		{object}	dto.UpdateTaskResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/tasks/update-status/{taskId} [patch]
func (t *taskHandler) UpdateStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	var requestBody dto.UpdateStatusOfATaskRequest

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

	updatedTaskResponse, err3 := t.taskService.UpdateStatus(id, &requestBody)

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, updatedTaskResponse)
}

// UpdateCategoryId godoc
//
//	@Summary		Update CategoryId of a task
//	@Description	Update CategoryId of a task by json
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		dto.UpdateCategoryIdOfATasIdkRequest	true	"Update categoryId of task request body"
//	@Param			taskId	path		uint									true	"Task ID request"
//	@Success		200		{object}	dto.UpdateCategoryIdOfTaskIdResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/tasks/update-category/{taskId} [patch]
func (t *taskHandler) UpdateCategoryIdOfTask(ctx *gin.Context) {
	taskID, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	var requestBody dto.UpdateCategoryIdOfATasIdkRequest

	err = ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if requestBody.CategoryId == 0 {
		errorMsg := "Category ID is required"
		categoryError := errs.NewBadRequest(errorMsg)
		ctx.JSON(categoryError.StatusCode(), categoryError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	updatedTaskResponse, err3 := t.taskService.UpdateCategoryIdOfTask(uint(taskID), &requestBody)

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, updatedTaskResponse)
}

// DeleteTask	godoc
//
//	@Summary		Delete a Task
//	@Description	Delete a Task by param
//	@Tags			tasks
//	@Produce		json
//	@Param			taskId		 path		uint						true	"Task ID request"
//	@Success		200			{object}	dto.DeleteTaskResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401			{object}	errs.MessageErrData
//	@Failure		400			{object}	errs.MessageErrData
//	@Router			/tasks/{taskId} [delete]
func (t *taskHandler) DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	taskIdUint, err := strconv.ParseUint(taskId, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Task id should be in unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err2 := t.taskService.DeleteTask(uint(taskIdUint))
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
