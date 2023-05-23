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
