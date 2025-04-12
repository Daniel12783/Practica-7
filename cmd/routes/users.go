package routes

import (
	"gin_http/cmd/controller"
	"gin_http/cmd/middleware"
	"gin_http/cmd/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userService *services.UserService) {
	admin := r.Group("/admin")
	admin.Use(middleware.APIKeyAuthMiddleware())

	//Controller
	userController := controller.NewUserController(userService)

	admin.GET("/users", userController.GetUsers)

	admin.POST("/users", userController.CreateUser)

	admin.PUT("/users/:id", userController.UpdateUser)

	admin.DELETE("/users/:id", userController.DeleteUser)
}
