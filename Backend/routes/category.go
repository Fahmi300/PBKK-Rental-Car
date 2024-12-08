package routes

import (
	"car-rental-app/controller"
	"car-rental-app/service"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine, CategoryController controller.CategoryController, jwtService service.JWTService) {
	categoryRoutes := router.Group("/api/category")
	{
		categoryRoutes.POST("", CategoryController.RegisterCategory)
		categoryRoutes.GET("", CategoryController.GetAllCategory)
		categoryRoutes.DELETE("/", CategoryController.DeleteCategory)
		categoryRoutes.PUT("/", CategoryController.UpdateCategory)
		categoryRoutes.GET("/:category_id", CategoryController.GetCategory)
	}
}
