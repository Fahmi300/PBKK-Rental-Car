package models

import (
	"car-rental-app/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID 		 int 	`gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"unique;type:varchar(50);not null" json:"email" binding:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     string `gorm:"type:varchar(30);not null" json:"role"`

	Bookings []Booking `gorm:"foreignKey:UserID" json:"bookings,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}