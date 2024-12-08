package service

import (
	"car-rental-app/models"
	"car-rental-app/repository"
	"context"

	"github.com/mashingan/smapping"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, categoryDTO models.Category) (models.Category, error)
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, id int) (models.Category, error) // ID is now int
	UpdateCategory(ctx context.Context, category models.Category) error
	DeleteCategory(ctx context.Context, id int) error // ID is now int
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(cr repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: cr,
	}
}

func (cs *categoryService) CreateCategory(ctx context.Context, categoryDTO models.Category) (models.Category, error) {
	category := models.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(categoryDTO))
	if err != nil {
		return category, err
	}
	return cs.categoryRepository.CreateCategory(ctx, category)
}

func (cs *categoryService) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	return cs.categoryRepository.GetAllCategories(ctx)
}

func (cs *categoryService) GetCategoryByID(ctx context.Context, id int) (models.Category, error) { // ID is now int
	return cs.categoryRepository.GetCategoryByID(ctx, id)
}

func (cs *categoryService) UpdateCategory(ctx context.Context, category models.Category) error {
	existingCategory, err := cs.categoryRepository.GetCategoryByID(ctx, category.ID)
	if err != nil {
		return err
	}

	// Update only fields that are provided
	if category.Name != "" {
		existingCategory.Name = category.Name
	}

	return cs.categoryRepository.UpdateCategory(ctx, existingCategory)
}

func (cs *categoryService) DeleteCategory(ctx context.Context, id int) error { // ID is now int
	return cs.categoryRepository.DeleteCategory(ctx, id)
}
