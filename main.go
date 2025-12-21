package main

import (
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/models"
	"sistem-manajemen-toko/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.LoadAppTemplates(r)

	config.ConnectDB()

	routes.WebRoutes(r)

	config.DB.AutoMigrate(&models.Category{}, models.Product{}, models.User{})

	r.Run(":3000")
}
