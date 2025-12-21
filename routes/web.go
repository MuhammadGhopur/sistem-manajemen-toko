package routes

import (
	"sistem-manajemen-toko/controllers"
	"sistem-manajemen-toko/middlewares"

	"github.com/gin-gonic/gin"
)

func WebRoutes(r *gin.Engine) {
	// LOGIN PUBLIK
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.LoginProcess)
	r.GET("/logout", controllers.Logout)
	r.GET("/users/create", controllers.CreateUserPage)
	r.POST("/users", controllers.CreateUser)

	// PROTEKSI
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	// CATEGORY
	auth.GET("/category", controllers.CategoryIndex)
	auth.GET("/category/create", controllers.CreateCategoryPage)
	auth.POST("/category", controllers.CreateCategory)
	auth.GET("/category/update/:id", controllers.UpdateCategoryPage)
	auth.POST("/category/update/:id", controllers.UpdateCategory)
	auth.GET("/category/delete/:id", controllers.DeleteCategory)

	// PRODUCT
	auth.GET("/product", controllers.ProductIndex)
	auth.GET("/product/create", controllers.CreateProductPage)
	auth.POST("/product", controllers.CreateProduct)
	auth.GET("/product/update/:id", controllers.UpdateProductPage)
	auth.POST("/product/update/:id", controllers.UpdateProduct)
	auth.GET("/product/delete/:id", controllers.DeleteProduct)

}
