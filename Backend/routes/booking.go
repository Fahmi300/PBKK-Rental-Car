package routes

import (
	"car-rental-app/controller"
	"car-rental-app/middleware"
	"car-rental-app/service"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine, BookingController controller.BookingController, jwtService service.JWTService) {
	bookingRoutes := router.Group("/api/booking")
	{
		bookingRoutes.POST("", middleware.Authenticate(jwtService), BookingController.RegisterBooking)
		bookingRoutes.GET("/my", middleware.Authenticate(jwtService), BookingController.GetAllBooking)
		bookingRoutes.DELETE("/", middleware.Authenticate(jwtService), BookingController.DeleteBooking)
		bookingRoutes.PUT("/", middleware.Authenticate(jwtService), BookingController.UpdateBooking)
		bookingRoutes.GET("/:booking_id", middleware.Authenticate(jwtService), BookingController.GetBooking)
	}
}
