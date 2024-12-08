package controller

import (
	"car-rental-app/common"
	"car-rental-app/models"
	"car-rental-app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	RegisterCategory(ctx *gin.Context)
	GetAllCategory(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
	jwtService     service.JWTService
}

func NewCategoryController(cs service.CategoryService,  js service.JWTService) CategoryController {
	return &categoryController{
		categoryService: cs,
		jwtService: js,
	}
}

// RegisterCategory creates a new category
func (cc *categoryController) RegisterCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		res := common.BuildErrorResponse("Failed to bind data", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	newCategory, err := cc.categoryService.CreateCategory(ctx.Request.Context(), category)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create category", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := common.BuildResponse(true, "Category created successfully", newCategory)
	ctx.JSON(http.StatusCreated, res)
}

// GetAllCategory retrieves all categories
func (cc *categoryController) GetAllCategory(ctx *gin.Context) {
	categories, err := cc.categoryService.GetAllCategories(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to retrieve categories", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := common.BuildResponse(true, "Categories retrieved successfully", categories)
	ctx.JSON(http.StatusOK, res)
}

// GetCategory retrieves a category by its ID
func (cc *categoryController) GetCategory(ctx *gin.Context) {
	// Get the 'id' parameter from the URL
	idStr := ctx.Param("category_id")

	// Convert the 'id' from string to int
	id, err := strconv.Atoi(idStr) // Convert to int (for database compatibility)
	if err != nil {
		res := common.BuildErrorResponse("Invalid ID format", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Get category by ID from service
	category, err := cc.categoryService.GetCategoryByID(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to retrieve category", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := common.BuildResponse(true, "Category retrieved successfully", category)
	ctx.JSON(http.StatusOK, res)
}

// UpdateCategory updates a category's details
func (cc *categoryController) UpdateCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		res := common.BuildErrorResponse("Failed to bind data", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Get the 'id' parameter from the URL
	idStr := ctx.Param("category_id")

	// Convert the 'id' from string to int
	id, err := strconv.Atoi(idStr) // Convert to int
	if err != nil {
		res := common.BuildErrorResponse("Invalid ID format", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	category.ID = id

	// Call the service to update the category
	err = cc.categoryService.UpdateCategory(ctx.Request.Context(), category)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update category", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := common.BuildResponse(true, "Category updated successfully", category)
	ctx.JSON(http.StatusOK, res)
}

// DeleteCategory deletes a category by its ID
func (cc *categoryController) DeleteCategory(ctx *gin.Context) {
	// Get the 'id' parameter from the URL
	idStr := ctx.Param("category_id")

	// Convert the 'id' from string to int
	id, err := strconv.Atoi(idStr) // Convert to int
	if err != nil {
		res := common.BuildErrorResponse("Invalid ID format", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Call the service to delete the category
	err = cc.categoryService.DeleteCategory(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to delete category", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := common.BuildResponse(true, "Category deleted successfully", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
