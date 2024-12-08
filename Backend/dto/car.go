package dto

type CarCreateDto struct {
	ID           int     `json:"id" form:"id"`               // ID is typically auto-incremented, so it may not be needed in creation.
	Name         string  `json:"name" form:"name"`
	Brand        string  `json:"brand" form:"brand"`
	Seat         int     `json:"seat" form:"seat"`
	Transmission string  `json:"transmission" form:"transmission"`
	Fuel         string  `json:"fuel" form:"fuel"`
	Luggage      bool    `json:"luggage" form:"luggage"`
	Insurance    bool    `json:"insurance" form:"insurance"`
	Year         int     `json:"year" form:"year"`
	PricePerDay  float64 `json:"price_per_day" form:"price_per_day"`
	Availability bool    `json:"availability" form:"availability"`
	CategoryID   int     `json:"category_id" form:"category_id"`
	Image        []byte  `json:"image" form:"image"`
}

type CarUpdateDto struct {
	Name         string  `json:"name" form:"name"`
	Brand        string  `json:"brand" form:"brand"`
	Seat         int     `json:"seat" form:"seat"`
	Transmission string  `json:"transmission" form:"transmission"`
	Fuel         string  `json:"fuel" form:"fuel"`
	Luggage      bool    `json:"luggage" form:"luggage"`
	Insurance    bool    `json:"insurance" form:"insurance"`
	Year         int     `json:"year" form:"year"`
	PricePerDay  float64 `json:"price_per_day" form:"price_per_day"`
	Availability bool    `json:"availability" form:"availability"`
	CategoryID   int     `json:"category_id" form:"category_id"`
	Image        []byte  `json:"image" form:"image"`
}
