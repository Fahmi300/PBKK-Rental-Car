package routes

import (
	"car-rental-app/controller"
	"car-rental-app/middleware"
	"car-rental-app/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, UserController controller.UserController, jwtService service.JWTService) {
	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("", UserController.RegisterUser)
		userRoutes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
		userRoutes.POST("/login", UserController.LoginUser)
		userRoutes.DELETE("/", middleware.Authenticate(jwtService), UserController.DeleteUser)
		userRoutes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
		userRoutes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)
	}
}