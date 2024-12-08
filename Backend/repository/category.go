package repository

import (
	"car-rental-app/models"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category models.Category) (models.Category, error)
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, id int) (models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}

func (db *categoryConnection) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	// Create a new category record in the database
	tx := db.connection.Create(&category)
	if tx.Error != nil {
		return models.Category{}, tx.Error
	}
	return category, nil
}

func (db *categoryConnection) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	// Fetch all categories, including their related cars (if any)
	var categories []models.Category
	tx := db.connection.Preload("Cars").Find(&categories)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categories, nil
}

func (db *categoryConnection) GetCategoryByID(ctx context.Context, id int) (models.Category, error) {
	// Find a category by ID, including its related cars
	var category models.Category
	tx := db.connection.Preload("Cars").Where("id = ?", id).Take(&category)
	if tx.Error != nil {
		return category, tx.Error
	}
	return category, nil
}

func (db *categoryConnection) UpdateCategory(ctx context.Context, category models.Category) error {
	// Update the category with the new data
	tx := db.connection.Save(&category)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (db *categoryConnection) DeleteCategory(ctx context.Context, id int) error {
	// Delete the category by ID
	tx := db.connection.Delete(&models.Category{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
