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

	// PROTEKSI
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	admin := r.Group("/")
	admin.Use(middlewares.RoleMiddleware("admin"))

	admin.GET("/users/create", controllers.CreateUserPage)
	admin.POST("/users", controllers.CreateUser)

	// CATEGORY
	auth.GET("/category", controllers.CategoryIndex)
	admin.GET("/category/create", controllers.CreateCategoryPage)
	admin.POST("/category", controllers.CreateCategory)
	admin.GET("/category/update/:id", controllers.UpdateCategoryPage)
	admin.POST("/category/update/:id", controllers.UpdateCategory)
	admin.GET("/category/delete/:id", controllers.DeleteCategory)

	// PRODUCT
	auth.GET("/product", controllers.ProductIndex)
	admin.GET("/product/create", controllers.CreateProductPage)
	admin.POST("/product", controllers.CreateProduct)
	admin.GET("/product/update/:id", controllers.UpdateProductPage)
	admin.POST("/product/update/:id", controllers.UpdateProduct)
	admin.GET("/product/delete/:id", controllers.DeleteProduct)

}
