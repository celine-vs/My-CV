// @title Category API
// @version 1.0
// @description REST API for managing categories.
// @host localhost:8080
// @BasePath /
package main

import (
	"backend/config"
	controllers "backend/controller"
	"backend/repository"
	"backend/router"
	"backend/services"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "backend/docs"
)

func main() {

	// 1. Connect database
	db := config.ConnectToDB()

	// 2. Create repository
	categoryRepository := repository.NewCategoryRepository(db)

	// 3. Create service
	categoryService := services.NewCategoryService(categoryRepository)

	// 4. Create controller
	categoryController := controllers.NewCategoryController(categoryService)

	// 5. Create router
	r := router.SetupRouter(categoryController)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 6. Start server
	r.Run(":8080")
}

// import "github.com/gin-gonic/gin"

// func main() {
//   router := gin.Default()
//   router.GET("/ping", func(c *gin.Context) {
//     c.JSON(200, gin.H{
//       "message": "pong",
//     })
//   })
//   router.Run() // listens on 0.0.0.0:8080 by default
// }
