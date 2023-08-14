package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/jeksilaen/api-builder/middlewares"
	"github.com/jeksilaen/api-builder/modules/user/helpers"
	"github.com/jeksilaen/api-builder/modules/user/models"
	"github.com/jeksilaen/api-builder/modules/user/usecases"
)

func InitUserHttpHandler(router *gin.Engine) {
	router.POST("/users/v1/login", LoginUser)
	router.POST("/users/v1/register", RegisterUser)
}

func RegisterUser(ctx *gin.Context) {
	userUsecase := usecases.NewUserCommandUsecase()
	validate := validator.New()

	// Decode the request JSON data into User object
	var req models.User
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedRegisterResponse(err.Error()))
		return
	}

	// Validate the request JSON data
	err := validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedRegisterResponse(err.Error()))
		return
	}

	// Create the user
	createdUser, err := userUsecase.CreateUser(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedRegisterResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.ReturnSucessRegisterResponse(createdUser))
}

func LoginUser(ctx *gin.Context) {
	userUsecase := usecases.NewUserCommandUsecase()
	validate := validator.New()

	// Decode the request JSON data into User object
	var req models.LoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedLoginResponse(err.Error()))
		return
	}

	// Validate the request JSON data
	err := validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ReturnFailedLoginResponse(err.Error()))
		return
	}

	// Find the user
	user, err := userUsecase.FindUserByEmailAndPassword(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, helpers.ReturnFailedLoginResponse(err.Error()))
		return
	}

	ctx.IndentedJSON(http.StatusOK, helpers.ReturnSucessLoginResponse(user))
}
