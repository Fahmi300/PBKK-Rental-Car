package routes

import (
	"car-rental-app/controller"
	"car-rental-app/middleware"
	"car-rental-app/service"

	"github.com/gin-gonic/gin"
)

func CarRoutes(router *gin.Engine, CarController controller.CarController, jwtService service.JWTService) {
	carRoutes := router.Group("/api/car")
	{
		carRoutes.POST("", middleware.Authenticate(jwtService), CarController.RegisterCar)
		carRoutes.GET("", CarController.GetAllCar)
		carRoutes.DELETE("/", middleware.Authenticate(jwtService), CarController.DeleteCar)
		carRoutes.PUT("/", middleware.Authenticate(jwtService), CarController.UpdateCar)
		carRoutes.GET("/:car_id", CarController.GetCar)
		carRoutes.GET("/:car_id/image", CarController.GetCarImage)

	}
}
