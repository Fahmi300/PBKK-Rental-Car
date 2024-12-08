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
	Description	string    `json:"description"`
	Location    string	  `json:"location"`
	Need		string    `json:"need"`
	Phone		string    `json:"phone"`

	UserID     int      `json:"user_id"`
	User	   User		`json:"user,omitempty"`
	CarID      int      `json:"car_id"`
	Car	   	   Car		`json:"car,omitempty"`
}
