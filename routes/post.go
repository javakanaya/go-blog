package routes

import (
	"penugasan4/controller"
	"penugasan4/middleware"
	"penugasan4/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.Engine, postController controller.PostController, jwtService service.JWTService) {
	postRoutes := router.Group("/post")
	{
		postRoutes.POST("", middleware.Authenticate(jwtService), postController.CreatePost)
		postRoutes.GET("", postController.GetAllPosts)
		postRoutes.GET("/:id", postController.GetPostByID)
		postRoutes.POST("/like/:id", middleware.Authenticate(jwtService), postController.LikePostByID)
		postRoutes.DELETE("/:id", middleware.Authenticate(jwtService), postController.DeletePost)
		postRoutes.PUT("/:id", middleware.Authenticate(jwtService), postController.UpdatePost)

	}
}
