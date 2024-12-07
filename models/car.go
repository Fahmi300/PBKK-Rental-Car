package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ID			 int	 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name	     string	 `gorm:"type:varchar(50);not null" json:"name"`
	Brand        string  `gorm:"type:varchar(50);not null" json:"brand"`
	Year         int     `json:"year"`
	PricePerDay  float64 `json:"price_per_day"`
	Availability bool    `json:"availability"`

	CategoryID	 int	 `json:"category_id"`
	Category     Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`

	Bookings     []Booking `gorm:"foreignKey:CarID" json:"bookings,omitempty"`
}
