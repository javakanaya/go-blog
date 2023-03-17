package routes

import (
	"penugasan4/controller"
	"penugasan4/middleware"
	"penugasan4/service"

	"github.com/gin-gonic/gin"
)

func CommentRouter(router *gin.Engine, commentController controller.CommentController, jwtService service.JWTService) {
	commentRoutes := router.Group("/comment")
	{
		commentRoutes.POST("", middleware.Authenticate(jwtService), commentController.CreateComment)
		commentRoutes.PUT("/:id", middleware.Authenticate(jwtService), commentController.UpdateComment)
		commentRoutes.DELETE("/:id", middleware.Authenticate(jwtService), commentController.DeleteComment)

	}
}
