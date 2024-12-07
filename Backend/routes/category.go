package routes

import (
	"car-rental-app/controller"
	"car-rental-app/middleware"
	"car-rental-app/service"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine, CategoryController controller.CategoryController, jwtService service.JWTService) {
	categoryRoutes := router.Group("/api/category")
	{
		categoryRoutes.POST("", CategoryController.RegisterCategory)
		categoryRoutes.GET("", middleware.Authenticate(jwtService), CategoryController.GetAllCategory)
		categoryRoutes.DELETE("/", middleware.Authenticate(jwtService), CategoryController.DeleteCategory)
		categoryRoutes.PUT("/", middleware.Authenticate(jwtService), CategoryController.UpdateCategory)
		categoryRoutes.GET("/:category_id", middleware.Authenticate(jwtService), CategoryController.GetCategory)
	}
}
