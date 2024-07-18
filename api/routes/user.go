package routes

import (
	"github.com/Asad2730/User_Micro_Crud_API/api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserClient) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/", userController.GetUsers)
		userRoutes.GET("/:id", userController.GetUser)
		userRoutes.PUT("/", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.RemoveUser)
		userRoutes.GET("/stream", userController.StreamUsers)
	}
}
