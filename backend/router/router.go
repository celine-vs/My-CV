package router

import (
	controllers "backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(categoryController *controllers.CategoryController) *gin.Engine {
	r := gin.Default()

	category := r.Group("/categories")
	{
		category.GET("", categoryController.GetAllCategory)
		category.GET("/:id", categoryController.GetACategory)
		category.POST("", categoryController.CreateCategory)
		category.PUT("/:id", categoryController.UpdateCategory)
		category.DELETE("/:id", categoryController.DeleteCategory)
	}

	return r
}
