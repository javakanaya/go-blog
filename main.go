package main

import (
	"os"
	"penugasan4/config"
	"penugasan4/controller"
	"penugasan4/repository"
	"penugasan4/routes"
	"penugasan4/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDatabaseConnection()

	jwtService := service.NewJWTService()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, jwtService)

	postRepository := repository.NewPostRepository(db, userRepository)
	postService := service.NewPostService(postRepository, userRepository)
	postController := controller.NewPostController(postService, jwtService)

	commentRepository := repository.NewCommentRepository(db, userRepository)
	commentService := service.NewCommentService(commentRepository)
	commentController := controller.NewCommentController(commentService, jwtService)

	server := gin.Default()

	routes.UserRouter(server, userController, jwtService)
	routes.PostRouter(server, postController, jwtService)
	routes.CommentRouter(server, commentController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

	config.CloseDatabaseConnection(db)

}
