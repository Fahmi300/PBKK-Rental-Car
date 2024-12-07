package dto

// import (
// 	"github.com/google/uuid"
// )

type UserCreateDto struct {
	ID        	int   		`gorm:"primary_key;not_null" json:"id" form:"id"`
	Name 		string 		`json:"name" form:"name" binding:"required"`
	Email 		string 		`json:"email" form:"email" binding:"required"`
	Password 	string  	`json:"password" form:"password" binding:"required"`
}

type UserUpdateDto struct {
	ID        	int   		`gorm:"primary_key;not_null" json:"id" form:"id"`
	Name 		string 		`json:"name" form:"name"`
	Email 		string 		`json:"email" form:"email"`
	Password 	string  	`json:"password" form:"password"`
}

type UserLoginDTO struct {
	Email 		string 		`json:"email" form:"email" binding:"email"`
	Password 	string  	`json:"password" form:"password" binding:"required"`
}