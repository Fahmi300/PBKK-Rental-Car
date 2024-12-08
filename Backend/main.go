package main

import (
	"car-rental-app/config"
	"car-rental-app/controller"
	"car-rental-app/database"
	"car-rental-app/repository"
	"car-rental-app/routes"
	"car-rental-app/service"

	"github.com/gin-gonic/gin"

	"log"
)

func main() {
	// Connect to the database
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	var (
		jwtService service.JWTService = service.NewJWTService()

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    service.UserService       = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)

		categoryRepository 	repository.CategoryRepository 	= repository.NewCategoryRepository(db)
		categoryService 	service.CategoryService 		= service.NewCategoryService(categoryRepository)
		categoryController	 controller.CategoryController	= controller.NewCategoryController(categoryService,  jwtService)

		carRepository repository.CarRepository = repository.NewCarRepository(db)
		carService    service.CarService       = service.NewCarService(carRepository)
		carController controller.CarController = controller.NewCarController(carService, jwtService)

		bookingRepository	repository.BookingRepository	= repository.NewBookingRepository(db)
		bookingService		service.BookingService			= service.NewBookingService(bookingRepository, carRepository)
		bookingController 	controller.BookingController 	= controller.NewBookingController(bookingService, jwtService)

	)

	// Auto-migrate models
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	if err := database.Seeder(db); err != nil {
		log.Fatalf("Error seeding database: %v", err)
	}

	// Set up routes
	r := gin.Default()
	routes.UserRoutes(r, userController, jwtService)
	routes.CarRoutes(r, carController, jwtService)
	routes.BookingRoutes(r, bookingController, jwtService)
	routes.CategoryRoutes(r, categoryController, jwtService)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Car Rental App!"})
	})

	r.Run(":8080")
}
