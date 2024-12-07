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
		carRoutes.POST("", CarController.RegisterCar)
		carRoutes.GET("", middleware.Authenticate(jwtService), CarController.GetAllCar)
		carRoutes.DELETE("/", middleware.Authenticate(jwtService), CarController.DeleteCar)
		carRoutes.PUT("/", middleware.Authenticate(jwtService), CarController.UpdateCar)
		carRoutes.GET("/:car_id", middleware.Authenticate(jwtService), CarController.GetCar)
	}
}