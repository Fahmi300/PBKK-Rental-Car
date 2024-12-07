package repository

import (
	"context"
	"car-rental-app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
	GetAllUser(ctx context.Context) ([]models.User, error)
	FindUserByEmail(ctx context.Context, email string) (models.User, error)
	FindUserByID(ctx context.Context, userID int) (models.User, error)
	DeleteUser(ctx context.Context, userID int) (error)
	UpdateUser(ctx context.Context, user models.User) (error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func(db *userConnection) RegisterUser(ctx context.Context, user models.User) (models.User, error) {
	uc := db.connection.Create(&user)
	if uc.Error != nil {
		return models.User{}, uc.Error
	}
	return user, nil
}

func(db *userConnection) GetAllUser(ctx context.Context) ([]models.User, error) {
	var listUser []models.User
	tx := db.connection.Find(&listUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listUser, nil
}

func(db *userConnection) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	ux := db.connection.Where("email = ?", email).Take(&user)
	if ux.Error != nil {
		return user, ux.Error
	}
	return user, nil
}

func(db *userConnection) FindUserByID(ctx context.Context, userID int) (models.User, error) {
	var user models.User
	ux := db.connection.Where("id = ?", userID).Take(&user)
	if ux.Error != nil {
		return user, ux.Error
	}
	return user, nil
}

func(db *userConnection) DeleteUser(ctx context.Context, userID int) (error) {
	uc := db.connection.Delete(&models.User{}, &userID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *userConnection) UpdateUser(ctx context.Context, user models.User) (error) {
	uc := db.connection.Updates(&user)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}