package controller

import (
	"net/http"
	"penugasan4/dto"
	"penugasan4/service"
	"penugasan4/utils"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	jwtService  service.JWTService
	userService service.UserService
}

func NewUserController(us service.UserService, jwt service.JWTService) UserController {
	return &userController{
		jwtService:  jwt,
		userService: us,
	}
}

func (uc *userController) RegisterUser(ctx *gin.Context) {
	var userDTO dto.UserRegister

	errDTO := ctx.ShouldBind(&userDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isRegistered := uc.userService.CheckUserByEmail(ctx, userDTO.Email)
	if isRegistered {
		response := utils.BuildErrorResponse("Email already registered", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.userService.RegisterUser(ctx, userDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to register user", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to register user", http.StatusOK, user)
	ctx.JSON(http.StatusOK, response)

}

func (uc *userController) LoginUser(ctx *gin.Context) {
	var userDTO dto.UserLogin
	errDTO := ctx.ShouldBind(&userDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.userService.GetUserByEmail(ctx, userDTO.Email)
	if err != nil {
		response := utils.BuildErrorResponse("User not registered", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isSuccess, _ := uc.userService.VerifyUser(ctx, userDTO.Email, userDTO.Password)
	if !isSuccess {
		response := utils.BuildErrorResponse("Failed to verify", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := uc.jwtService.GenerateToken(user.ID, user.Role)
	response := utils.BuildSuccessResponse("Login suceccfuly", http.StatusOK, token)
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

func (uc *userController) GetUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get ID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.userService.GetUserByID(ctx, userID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to find user", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Sucess to get user", http.StatusOK, user)
	ctx.JSON(http.StatusOK, response)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	var userDTO dto.UserUpdate
	errDTO := ctx.ShouldBind(&userDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get ID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userDTO.ID = userID
	err = uc.userService.UpdateUser(ctx, userDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to update user", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Succes to update user", http.StatusOK, userDTO)
	ctx.JSON(http.StatusOK, response)

}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get ID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = uc.userService.DeleteUser(ctx, userID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to delete user", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Succes delete user", http.StatusOK, nil)
	ctx.JSON(http.StatusOK, response)
}
