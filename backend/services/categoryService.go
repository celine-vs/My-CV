package services

import (
	"backend/model"
	"backend/repository"
)

type CategoryService interface {
	CreateCategory(category *model.Category)
	GetCategoryByID(id uint) (*model.Category, error)
	GetCategories() ([]model.Category, error)
	UpdateCategory(category *model.Category)
	DeleteCategory(category *model.Category)
}

type categoryService struct {
	repository *repository.CategoryRepository
}

func NewCategoryService(categoryRepository *repository.CategoryRepository) CategoryService {
	return &categoryService{
		repository: categoryRepository,
	}
}

func (svc *categoryService) CreateCategory(category *model.Category) {
	svc.repository.Create(category)
}

func (svc *categoryService) GetCategories() ([]model.Category, error) {
	return svc.repository.FindAll()
}

func (svc *categoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return svc.repository.FindByID(id)
}

func (svc *categoryService) UpdateCategory(category *model.Category) {
	svc.repository.Update(category)
}

func (svc *categoryService) DeleteCategory(category *model.Category) {
	svc.repository.Delete(category)
}
