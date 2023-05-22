package http_handlers

import (
	"final-project-3/dto"
	"final-project-3/pkg/errs"
	"final-project-3/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	createdUser, err3 := u.userService.RegisterUser(&requestBody)
	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (u *userHandler) LoginUser(ctx *gin.Context) {
	var requestBody dto.LoginUserRequest

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

	token, err := u.userService.LoginUser(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (u *userHandler) UpdateUser(ctx *gin.Context) {
	var requestBody dto.UpdateUserRequest

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

	// mustget = ambil data dari middleware authentication.
	// Tp hasil returnnya hanya empty interface, jadi harus
	// di cast dulu ke jwt.MapClaims.
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	updatedUserResponse, err2 := u.userService.UpdateUser(userId, &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, updatedUserResponse)
}

func (u *userHandler) DeleteUser(ctx *gin.Context) {
	// mustget = ambil data dari middleware authentication.
	// Tp hasil returnnya hanya empty interface, jadi harus
	// di cast dulu ke jwt.MapClaims.
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	response, err2 := u.userService.DeleteUser(userId)

	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
