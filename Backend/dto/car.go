package dto

// import (
// 	"github.com/google/uuid"
// )

type CarCreateDto struct {
	ID          int     `gorm:"primary_key;not_null" json:"id" form:"id"`
	Name        string  `json:"name" form:"name" binding:"required"`
	Brand       string  `json:"brand" form:"brand" binding:"required"`
	Year        int     `json:"year" form:"year" binding:"required"`
	PricePerDay float64 `json:"price_per_day" form:"priceperday" binding:"required"`
	Category    string  `json:"category" form:"category" binding:"required"`
}

type CarUpdateDto struct {
	ID          int     `gorm:"primaryKey;not_null" json:"id" form:"id"`
	Name        string  `json:"name" form:"name"`
	Brand       string  `json:"brand" form:"brand"`
	Year        int     `json:"year" form:"year"`
	PricePerDay float64 `json:"price_per_day" form:"priceperday"`
	Category    string  `json:"category" form:"category"`
}
