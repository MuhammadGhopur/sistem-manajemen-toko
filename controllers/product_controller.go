package controllers

import (
	"net/http"
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/helpers"
	"sistem-manajemen-toko/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductIndex(c *gin.Context) {
	var product []models.Product

	config.DB.Preload("Category").Find(&product)

	c.HTML(http.StatusOK, "products.html", gin.H{
		"products": product,
	})
}

func CreateProductPage(c *gin.Context) {
	var categories []models.Category

	config.DB.Find(&categories)

	c.HTML(http.StatusOK, "create_product.html", gin.H{
		"categories": categories,
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	product.Name = c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))

	categoryID, _ := strconv.Atoi(c.PostForm("category_id"))
	product.CategoryID = uint(categoryID)

	product.Price = price

	config.DB.Create(&product)

	helpers.CreateAuditLog(
		c,
		"CREATE",
		"products",
		product.ID,
		"Menambah produk baru",
	)

	c.Redirect(http.StatusFound, "/product")

}

func UpdateProductPage(c *gin.Context) {
	var product models.Product

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&product, id)

	c.HTML(http.StatusOK, "update_product.html", gin.H{
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&product, id)

	product.Name = c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))

	product.Price = price

	config.DB.Save(&product)

	helpers.CreateAuditLog(
		c,
		"UPDATE",
		"products",
		product.ID,
		"Update produk",
	)

	c.Redirect(http.StatusFound, "/product")

}

func DeleteProduct(c *gin.Context) {
	var product models.Product

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&product, id)

	config.DB.Delete(&product)

	helpers.CreateAuditLog(
		c,
		"DELETE",
		"products",
		product.ID,
		"Hapus produk",
	)

	c.Redirect(http.StatusFound, "/product")
}
