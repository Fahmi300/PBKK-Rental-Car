package controller

import (
	"car-rental-app/common"
	"car-rental-app/dto"
	"car-rental-app/service"
	"net/http"
	"strconv"
	"io/ioutil"

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

	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil { // Limit to 10 MB
		res := common.BuildErrorResponse("Input Salah", "Gagal membaca multipart data", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		res := common.BuildErrorResponse("Input Salah", "Gagal membaca file gambar", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	defer file.Close()

	// Handle form values
	carDTO.Name = ctx.PostForm("name")
	carDTO.Brand = ctx.PostForm("brand")
	carDTO.Seat, _ = strconv.Atoi(ctx.PostForm("seat"))
	carDTO.Transmission = ctx.PostForm("transmission")
	carDTO.Fuel = ctx.PostForm("fuel")
	carDTO.Luggage, _ = strconv.ParseBool(ctx.PostForm("luggage"))
	carDTO.Insurance, _ = strconv.ParseBool(ctx.PostForm("insurance"))
	carDTO.Year, _ = strconv.Atoi(ctx.PostForm("year"))
	carDTO.PricePerDay, _ = strconv.ParseFloat(ctx.PostForm("pricePerDay"), 64)
	carDTO.Availability, _ = strconv.ParseBool(ctx.PostForm("availability"))
	carDTO.CategoryID, _ = strconv.Atoi(ctx.PostForm("categoryId"))

	// Convert file into a byte slice
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		res := common.BuildErrorResponse("Input Salah", "Gagal membaca isi file", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	carDTO.Image = imageData

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

func (c *carController) GetCarImage(ctx *gin.Context) {
	// Ambil ID mobil dari parameter URL
	carID := ctx.Param("car_id")
	id, err := strconv.Atoi(carID)
	if err != nil {
		res := common.BuildErrorResponse("Invalid car ID", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Ambil data mobil beserta gambar dari database
	car, err := c.carService.GetCarByID(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Mobil tidak ditemukan", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	// Jika gambar ditemukan, kembalikan sebagai blob
	if car.Image != nil {
		// Tentukan jenis konten gambar berdasarkan tipe file (misalnya jpeg atau png)
		ctx.Header("Content-Type", "image/png") // Sesuaikan dengan jenis gambar yang Anda simpan (JPEG, PNG, dll)
		ctx.Data(http.StatusOK, "image/png", car.Image) // Mengirim gambar sebagai response
		return
	}

	// Jika gambar tidak ada, kirim respons error
	res := common.BuildErrorResponse("Gambar tidak ditemukan", "Gambar tidak tersedia", common.EmptyObj{})
	ctx.JSON(http.StatusNotFound, res)
}
