package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID 		 int 	`gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	
	Cars []Car	`json:"cars,omitempty"`
}
