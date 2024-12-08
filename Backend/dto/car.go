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
	Availability bool   `json:"availability" form:"availability" binding:"required"`
	CategoryID   int    `json:"category_id" form:"category_id" binding:"required"`
}

type CarUpdateDto struct {
	Name        string  `json:"name" form:"name"`
	Brand       string  `json:"brand" form:"brand"`
	Year        int     `json:"year" form:"year"`
	PricePerDay float64 `json:"price_per_day" form:"priceperday"`
	Category    string  `json:"category" form:"category"`
	Availability bool   `json:"availability" form:"availability"`
	CategoryID   int     `json:"category_id" form:"category_id"`
}
