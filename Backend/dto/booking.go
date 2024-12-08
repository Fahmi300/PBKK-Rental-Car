package dto
import (
	"time"
)

type BookingCreateDto struct {
	ID         int    `gorm:"primary_key;not_null" json:"id" form:"id"`
	StartDate  time.Time `json:"start_date" form:"start_date" binding:"required"`
	EndDate    time.Time `json:"end_date" form:"end_date" binding:"required"`
	UserID     int    `json:"user_id"`
	CarID      int    `json:"car_id" form:"car_id" binding:"required"`
}

type BookingUpdateDto struct {
	ID         int    `gorm:"primary_key;not_null" json:"id" form:"id"`
	StartDate  time.Time `json:"start_date" form:"start_date" binding:"required"`
	EndDate    time.Time `json:"end_date" form:"end_date" binding:"required"`
	UserID     int    `json:"user_id" form:"user_id"`
	CarID      int    `json:"car_id" form:"car_id"`
}
