package controller

import (
	"net/http"
	"penugasan4/dto"
	"penugasan4/service"
	"penugasan4/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentController struct {
	jwtService     service.JWTService
	commentService service.CommentService
}

func NewCommentController(cs service.CommentService, jwt service.JWTService) CommentController {
	return &commentController{
		jwtService:     jwt,
		commentService: cs,
	}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := cc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get userID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var commentDTO dto.CreateCommentDTO
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	commentDTO.UserID = userID
	comment, err := cc.commentService.CreateComment(ctx, commentDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to create comment", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to create comment", http.StatusOK, comment)
	ctx.JSON(http.StatusOK, response)
}

func (cc *commentController) UpdateComment(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := cc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get userID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := ctx.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get commentID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var commentDTO dto.CreateCommentDTO
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	commentDTO.UserID = userID
	commentDTO.ID = commentID
	comment, err := cc.commentService.UpdateComment(ctx, commentDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to update comment", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to update comment", http.StatusOK, comment)
	ctx.JSON(http.StatusOK, response)
}

func (cc *commentController) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get commentID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = cc.commentService.DeleteComment(ctx, commentID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to delete comment", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to delete comment", http.StatusOK, nil)
	ctx.JSON(http.StatusOK, response)
}
