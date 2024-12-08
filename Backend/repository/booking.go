package repository

import (
	"car-rental-app/models"
	"context"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking models.Booking) (models.Booking, error)
	GetAllBookings(ctx context.Context, userID int) ([]models.Booking, error)
	FindBookingByID(ctx context.Context, bookingID int, userID int) (models.Booking, error)
	UpdateBooking(ctx context.Context, booking models.Booking) error
	DeleteBooking(ctx context.Context, bookingID int) error
	GetBookingsByCarID(ctx context.Context, carID int) ([]models.Booking, error)
}

type bookingConnection struct {
	connection *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingConnection{
		connection: db,
	}
}

func (db *bookingConnection) CreateBooking(ctx context.Context, booking models.Booking) (models.Booking, error) {
	result := db.connection.Create(&booking)
	if result.Error != nil {
		return models.Booking{}, result.Error
	}
	return booking, nil
}

func (db *bookingConnection) GetAllBookings(ctx context.Context, userID int) ([]models.Booking, error) {
	var bookings []models.Booking
	tx := db.connection.Where("user_id = ?", userID).Find(&bookings)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bookings, nil
}

func (db *bookingConnection) FindBookingByID(ctx context.Context, bookingID int , userID int) (models.Booking, error) {
	var booking models.Booking
	tx := db.connection.Where("id = ? AND user_id = ?", bookingID, userID).Take(&booking)
	if tx.Error != nil {
		return booking, tx.Error
	}
	return booking, nil
}

func (db *bookingConnection) UpdateBooking(ctx context.Context, booking models.Booking) error {
	tx := db.connection.Updates(&booking)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (db *bookingConnection) DeleteBooking(ctx context.Context, bookingID int) error {
	tx := db.connection.Delete(&models.Booking{}, bookingID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (db *bookingConnection) GetBookingsByCarID(ctx context.Context, carID int) ([]models.Booking, error) {
    var bookings []models.Booking
	tx := db.connection.Where("car_id = ?", carID).Find(&bookings)
	if tx.Error != nil {
		return nil, tx.Error
	}
    return bookings, nil
}

