package service

import (
	"car-rental-app/dto"
	"car-rental-app/models"
	"car-rental-app/repository"
	"context"
	"time"
	"math"

	"github.com/mashingan/smapping"
)

type BookingService interface {
	CreateBooking(ctx context.Context, bookingDto dto.BookingCreateDto, userID int) (models.Booking, error)
	GetAllBookings(ctx context.Context, userID int) ([]models.Booking, error)
	GetBooking(ctx context.Context, bookingID int, userID int) (models.Booking, error)
	UpdateBooking(ctx context.Context, bookingDto dto.BookingUpdateDto) error
	DeleteBooking(ctx context.Context, bookingID int) error
	IsCarAvailable(ctx context.Context, carID int, startDate time.Time, endDate time.Time) (bool, error)
}

type bookingService struct {
	bookingRepository repository.BookingRepository
	carRepository repository.CarRepository
}

func NewBookingService(br repository.BookingRepository, cr repository.CarRepository) BookingService {
	return &bookingService{
		bookingRepository: br,
		carRepository: cr,
	}
}

func (s *bookingService) CreateBooking(ctx context.Context, bookingDTO dto.BookingCreateDto, userID int) (models.Booking, error) {
	booking := models.Booking{}
	err := smapping.FillStruct(&booking, smapping.MapFields(bookingDTO))

	totalprice, err := s.CalculateTotalPrice(ctx, booking.CarID, booking.StartDate, booking.EndDate)
    if err != nil {
        return booking, err
    }
	booking.TotalPrice = totalprice
	booking.UserID = userID
	if err != nil {
		return booking, err
	}
	return s.bookingRepository.CreateBooking(ctx, booking)
}

func (s *bookingService) GetAllBookings(ctx context.Context, userID int) ([]models.Booking, error) {
	return s.bookingRepository.GetAllBookings(ctx, userID)
}

func (s *bookingService) GetBooking(ctx context.Context, bookingID int, userID int) (models.Booking, error) {
	return s.bookingRepository.FindBookingByID(ctx, bookingID, userID)
}

func (s *bookingService) UpdateBooking(ctx context.Context, bookingDTO dto.BookingUpdateDto) error {
	booking := models.Booking{}
	err := smapping.FillStruct(&booking, smapping.MapFields(bookingDTO))
	if err != nil {
		return err
	}

	return s.bookingRepository.UpdateBooking(ctx, booking)
}

func (s *bookingService) DeleteBooking(ctx context.Context, bookingID int) error {
	return s.bookingRepository.DeleteBooking(ctx, bookingID)
}

func (s *bookingService) IsCarAvailable(ctx context.Context, carID int, startDate time.Time, endDate time.Time) (bool, error) {
    // Query the database to check for overlapping bookings
    bookings, err := s.bookingRepository.GetBookingsByCarID(ctx, carID)
    if err != nil {
        return false, err
    }

    // Check for overlapping periods
    for _, booking := range bookings {
        if (startDate.Before(booking.EndDate) && endDate.After(booking.StartDate)) {
            return false, nil
        }
    }
    return true, nil
}

func (s *bookingService) CalculateTotalPrice(ctx context.Context, carID int, startDate time.Time, endDate time.Time) (float64, error) {
    // Ambil informasi mobil berdasarkan carID dari repository car
    car, err := s.carRepository.FindCarByID(ctx, carID)
    if err != nil {
        return 0, err
    }

    // Hitung jumlah hari
    duration := endDate.Sub(startDate).Hours() / 24
    if duration < 1 {
        duration = 1 // Minimal 1 hari
    }

	duration = math.Ceil(duration)

    // Hitung total harga
    totalPrice := duration * car.PricePerDay
    return totalPrice, nil
}

