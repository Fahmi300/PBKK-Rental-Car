package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID  	   int	     `gorm:"primaryKey;autoIncrement" json:"id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalPrice float64   `json:"total_price"`

	UserID     int      `json:"user_id"`
	User	   int		`json:"user,omitempty"`
	CarID      int      `json:"car_id"`
	Car	   	   int		`json:"car,omitempty"`
}
