package controller

import (
	"net/http"
	"penugasan4/dto"
	"penugasan4/service"
	"penugasan4/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController interface {
	CreatePost(ctx *gin.Context)
	GetAllPosts(ctx *gin.Context)
	GetPostByID(ctx *gin.Context)
	LikePostByID(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
}

type postController struct {
	jwtService  service.JWTService
	postService service.PostService
}

func NewPostController(ps service.PostService, jwt service.JWTService) PostController {
	return &postController{
		jwtService:  jwt,
		postService: ps,
	}
}

func (pc *postController) CreatePost(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := pc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get userID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var postDTO dto.PostCreateDTO
	errDTO := ctx.ShouldBind(&postDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	postDTO.UserID = userID
	post, err := pc.postService.CreatePost(ctx, postDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to create post", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to create post", http.StatusOK, post)
	ctx.JSON(http.StatusOK, response)
}

func (pc *postController) GetAllPosts(ctx *gin.Context) {
	posts, err := pc.postService.GetAllPosts(ctx)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get all posts", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to get all posts", http.StatusOK, posts)
	ctx.JSON(http.StatusOK, response)
}

func (pc *postController) GetPostByID(ctx *gin.Context) {
	id := ctx.Param("id")
	postID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	post, err := pc.postService.GetPostById(ctx, postID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get post", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to get post", http.StatusOK, post)
	ctx.JSON(http.StatusOK, response)
}

func (pc *postController) LikePostByID(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := pc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get userID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	post_id := ctx.Param("id")
	postID, err := strconv.ParseUint(post_id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get postID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err = pc.postService.LikePostByPostID(ctx, userID, postID); err != nil {
		response := utils.BuildErrorResponse("Failed to like post", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to like post", http.StatusOK, nil)
	ctx.JSON(http.StatusOK, response)
}

func (pc *postController) UpdatePost(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := pc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get userID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	post_id := ctx.Param("id")
	postID, err := strconv.ParseUint(post_id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get postID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var postDTO dto.PostUpdateDTO
	if err := ctx.ShouldBindJSON(&postDTO); err != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	postDTO.UserID = userID
	postDTO.ID = postID
	if err = pc.postService.UpdatePost(ctx, postDTO); err != nil {
		response := utils.BuildErrorResponse("Failed to update post", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to update post", http.StatusOK, postDTO)
	ctx.JSON(http.StatusOK, response)
}

func (pc *postController) DeletePost(ctx *gin.Context) {
	post_id := ctx.Param("id")
	postID, err := strconv.ParseUint(post_id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get postID", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err = pc.postService.DeletePost(ctx, postID); err != nil {
		response := utils.BuildErrorResponse("Failed to delete post", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Success to delete post", http.StatusOK, nil)
	ctx.JSON(http.StatusOK, response)
}
