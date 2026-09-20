package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

// Constructor
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {

	var categories []model.Category

	result := r.db.Find(&categories)

	return categories, result.Error
}

func (r *CategoryRepository) FindByID(id uint) (*model.Category, error) {

	var category model.Category

	result := r.db.First(&category, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (r *CategoryRepository) Create(category *model.Category) error {

	result := r.db.Create(category)

	return result.Error
}

func (r *CategoryRepository) Update(category *model.Category) error {

	result := r.db.Save(category)

	return result.Error
}

func (r *CategoryRepository) Delete(category *model.Category) error {

	result := r.db.Delete(category)

	return result.Error
}
