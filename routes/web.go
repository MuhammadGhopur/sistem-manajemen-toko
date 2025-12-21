package routes

import (
	"sistem-manajemen-toko/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriyRoutes(r *gin.Engine) {
	r.GET("/category", controllers.CategoryIndex)
	r.GET("/category/create", controllers.CreateCategoryPage)
	r.POST("/category", controllers.CreateCategory)
	r.GET("/category/update/:id", controllers.UpdateCategoryPage)
	r.POST("/category/update/:id", controllers.UpdateCategory)
	r.GET("/category/delete/:id", controllers.DeleteCategory)
}

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", controllers.ProductIndex)
	r.GET("/product/create", controllers.CreateProductPage)
	r.POST("/product", controllers.CreateProduct)
	r.GET("/product/update/:id", controllers.UpdateProductPage)
	r.POST("/product/update/:id", controllers.UpdateProduct)
	r.GET("/product/delete/:id", controllers.DeleteProduct)
}
