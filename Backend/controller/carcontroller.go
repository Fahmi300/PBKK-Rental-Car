package controller

import (
	"car-rental-app/common"
	"car-rental-app/dto"
	"car-rental-app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarController interface {
	RegisterCar(ctx *gin.Context)
	GetAllCar(ctx *gin.Context)
	DeleteCar(ctx *gin.Context)
	UpdateCar(ctx *gin.Context)
	GetCar(ctx *gin.Context)
}

type carController struct {
	carService service.CarService
	jwtService service.JWTService
}

func NewCarController(cs service.CarService, js service.JWTService) CarController {
	return &carController{
		carService: cs,
		jwtService: js,
	}
}

func (c *carController) RegisterCar(ctx *gin.Context) {
	var carDTO dto.CarCreateDto

	token := ctx.MustGet("token").(string)
	isAdmin, err := c.jwtService.IsUserAdmin(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if !isAdmin {
		response := common.BuildErrorResponse("Akses Ditolak", "Hanya Admin yang Dapat Melakukan Aksi Ini", nil)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	err = ctx.ShouldBind(&carDTO)
	if err != nil {
		res := common.BuildErrorResponse("Input Salah", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	checkCar, _ := c.carService.CheckCar(ctx.Request.Context(), carDTO.Name)
	if checkCar {
		res := common.BuildErrorResponse("Mobil Sudah Ada", "Nama Sama", common.EmptyObj{})
		ctx.JSON(http.StatusConflict, res)
		return
	}

	car, err := c.carService.RegisterCar(ctx.Request.Context(), carDTO)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambah Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambah Mobil", car)
	ctx.JSON(http.StatusOK, res)
}

func (c *carController) GetAllCar(ctx *gin.Context) {
	cars, err := c.carService.GetAllCar(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menerima Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menerima Mobil", cars)
	ctx.JSON(http.StatusOK, res)
}

func (c *carController) GetCar(ctx *gin.Context) {
	// Get car ID from URL parameter
	carID := ctx.Param("car_id")
	id, err := strconv.Atoi(carID)
	if err != nil {
		res := common.BuildErrorResponse("Invalid car ID", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Fetch the car with the associated category using Preload
	car, err := c.carService.GetCarWithCategory(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menerima Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	// Return the response
	res := common.BuildResponse(true, "Berhasil Menerima Mobil", car)
	ctx.JSON(http.StatusOK, res)
}

func (c *carController) UpdateCar(ctx *gin.Context) {
	var carDTO dto.CarUpdateDto

	token := ctx.MustGet("token").(string)
	isAdmin, err := c.jwtService.IsUserAdmin(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if !isAdmin {
		response := common.BuildErrorResponse("Akses Ditolak", "Hanya Admin yang Dapat Melakukan Aksi Ini", nil)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	err = ctx.ShouldBind(&carDTO)
	if err != nil {
		res := common.BuildErrorResponse("Input Salah", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = c.carService.UpdateCar(ctx.Request.Context(), carDTO)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Membaharui Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Membaharui Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *carController) DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("car_id")
	id, err := strconv.Atoi(carID)
	if err != nil {
		res := common.BuildErrorResponse("Salah Car Id", "Car Id Harus Nomor", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = c.carService.DeleteCar(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menghapus Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
