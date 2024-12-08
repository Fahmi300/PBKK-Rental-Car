package dto

// CategoryCreateDto is used for creating a new category
type CategoryCreateDto struct {
	ID   string `gorm:"primary_key;not_null" json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
}

// CategoryUpdateDto is used for updating an existing category
type CategoryUpdateDto struct {
	Name string `json:"name" form:"name"`
}
