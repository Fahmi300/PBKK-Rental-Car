package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID           int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string  `gorm:"type:varchar(50);not null" json:"name"`
	Brand        string  `gorm:"type:varchar(50);not null" json:"brand"`
	Seat         int     `json:"seat"`
	Transmission string  `json:"transmission"`
	Fuel         string  `json:"fuel"`
	Luggage      bool    `json:"luggage"`
	Insurance    bool    `json:"insurance"`
	Year         int     `json:"year"`
	PricePerDay  float64 `json:"price_per_day"`
	Availability bool    `json:"availability"`

	CategoryID int `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`

	Bookings []Booking `gorm:"foreignKey:CarID" json:"bookings,omitempty"`

	Image []byte `gorm:"type:longblob" json:"image,omitempty"` // Changed to use []byte with longblob for larger storage
}
