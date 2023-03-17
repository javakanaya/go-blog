package routes

import (
	"penugasan4/controller"
	"penugasan4/middleware"
	"penugasan4/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("", userController.RegisterUser)
		userRoutes.GET("", middleware.Authenticate(jwtService), userController.GetUser)
		userRoutes.POST("/login", userController.LoginUser)
		userRoutes.PUT("/", middleware.Authenticate(jwtService), userController.UpdateUser)
		userRoutes.DELETE("/", middleware.Authenticate(jwtService), userController.DeleteUser)
	}
}
