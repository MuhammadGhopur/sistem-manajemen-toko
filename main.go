package main

import (
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/models"
	"sistem-manajemen-toko/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	config.ConnectDB()

	routes.CategoriyRoutes(r)
	routes.ProductRoutes(r)

	config.DB.AutoMigrate(&models.Category{}, models.Product{})

	r.Run(":3000")
}
