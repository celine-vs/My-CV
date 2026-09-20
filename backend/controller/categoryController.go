package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/model"
	"backend/services"
)

type CategoryController struct {
	services services.CategoryService
}

func NewCategoryController(services services.CategoryService) *CategoryController {
	return &CategoryController{
		services: services,
	}
}


// GetAllCategory godoc
// @Summary Get all categories
// @Description Get a list of all categories
// @Tags categories
// @Produce json
// @Success 200 {array} model.Category
// @Failure 400 {object} map[string]string
// @Router /categories [get]
func (ctrl *CategoryController) GetAllCategory(c *gin.Context) {
	categories, err := ctrl.services.GetCategories()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}


// GetACategory godoc
// @Summary Get category by ID
// @Description Get a category using its ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [get]
func (ctrl *CategoryController) GetACategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	Category, err := ctrl.services.GetCategoryByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	c.JSON(http.StatusOK, Category)
}

// CreateCategory godoc
// @Summary Create a category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.Category true "Category data"
// @Success 201 {object} model.Category
// @Failure 400 {object} map[string]string
// @Router /categories [post]
func (ctrl *CategoryController) CreateCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	ctrl.services.CreateCategory(&category)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update an existing category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.Category true "Category data"
// @Success 200 {object} model.Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [put]
func (ctrl *CategoryController) UpdateCategory(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalid ID",
	// 	})
	// 	return
	// }

	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	ctrl.services.UpdateCategory(&category)

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "Category not found",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, nil)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [delete]
func (ctrl *CategoryController) DeleteCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	ctrl.services.DeleteCategory(&category)

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "Category not found",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Category deleted successfully",
	// })
}
