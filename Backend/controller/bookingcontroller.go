package controller

import (
	"car-rental-app/common"
	"car-rental-app/dto"
	"car-rental-app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingController interface {
	RegisterBooking(ctx *gin.Context)
	GetAllBooking(ctx *gin.Context)
	GetBooking(ctx *gin.Context)
	UpdateBooking(ctx *gin.Context)
	DeleteBooking(ctx *gin.Context)
}

type bookingController struct {
	bookingService service.BookingService
	jwtService     service.JWTService
}
	
func NewBookingController(bs service.BookingService, js service.JWTService) BookingController {
	return &bookingController{
		bookingService: bs,
		jwtService:     js,
	}
}

func (b *bookingController) RegisterBooking(ctx *gin.Context) {
	var booking dto.BookingCreateDto

	token := ctx.MustGet("token").(string)
	userID, err := b.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err = ctx.ShouldBind(&booking)
	if err != nil {
		res := common.BuildErrorResponse("Invalid Input", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	isAvailable, err := b.bookingService.IsCarAvailable(ctx.Request.Context(), booking.CarID, booking.StartDate, booking.EndDate)
    if err != nil {
        res := common.BuildErrorResponse("Error checking car availability", err.Error(), common.EmptyObj{})
        ctx.JSON(http.StatusInternalServerError, res)
        return
    }
    if !isAvailable {
        res := common.BuildErrorResponse("Car not available", "The selected car is not available during the requested period", common.EmptyObj{})
        ctx.JSON(http.StatusConflict, res)
        return
    }


	result, err := b.bookingService.CreateBooking(ctx.Request.Context(), booking, userID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create booking", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := common.BuildResponse(true, "Booking Created", result)
	ctx.JSON(http.StatusCreated, res)
}

func (b *bookingController) GetAllBooking(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := b.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	bookings, err := b.bookingService.GetAllBookings(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Failed to fetch bookings", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := common.BuildResponse(true, "Bookings Retrieved", bookings)
	ctx.JSON(http.StatusOK, res)
}

func (b *bookingController) GetBooking(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := b.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	bookingID, err := strconv.Atoi(ctx.Param("booking_id"))
	if err != nil {
		res := common.BuildErrorResponse("Invalid ID", "Invalid booking ID", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	booking, err := b.bookingService.GetBooking(ctx.Request.Context(), bookingID, userID)
	if err != nil {
		res := common.BuildErrorResponse("Booking not found", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := common.BuildResponse(true, "Booking Retrieved", booking)
	ctx.JSON(http.StatusOK, res)
}

func (b *bookingController) UpdateBooking(ctx *gin.Context) {
	var bookingDTO dto.BookingUpdateDto
	err := ctx.ShouldBind(&bookingDTO)
	if err != nil {
		res := common.BuildErrorResponse("Invalid Input", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	id, err := strconv.Atoi(ctx.Param("booking_id"))
	if err != nil {
		res := common.BuildErrorResponse("Invalid ID", "Invalid booking ID", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	bookingDTO.ID = id

	// In the UpdateBooking method:
	err = b.bookingService.UpdateBooking(ctx.Request.Context(), bookingDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update booking", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := common.BuildResponse(true, "Booking Updated", common.EmptyObj{}) // Remove `result` here
	ctx.JSON(http.StatusOK, res)

}

func (b *bookingController) DeleteBooking(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("booking_id"))
	if err != nil {
		res := common.BuildErrorResponse("Invalid ID", "Invalid booking ID", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = b.bookingService.DeleteBooking(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to delete booking", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := common.BuildResponse(true, "Booking Deleted", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
