package http_handlers

import (
	"final-project-3/dto"
	"final-project-3/pkg/errs"
	"final-project-3/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService: userService}
}

// RegisterUser godoc
//
//	@Summary		Register a user
//	@Description	Register a user by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.NewUserRequest	true	"Create user request body"
//	@Success		201		{object}	dto.NewUserResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/users [post]
func (u *userHandler) RegisterUser(ctx *gin.Context) {
	var requestBody dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdUser, err := u.userService.RegisterUser(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (u *userHandler) LoginUser(ctx *gin.Context) {

}

func (u *userHandler) UpdateUser(ctx *gin.Context) {

}

func (u *userHandler) DeleteUser(ctx *gin.Context) {

}
