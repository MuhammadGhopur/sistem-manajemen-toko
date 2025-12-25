package controllers

import (
	"net/http"
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/helpers"
	"sistem-manajemen-toko/models"
	"sistem-manajemen-toko/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductIndex(c *gin.Context) {
	var product []models.Product

	err := services.GetProduct(&product)
	if err != nil {
		c.HTML(http.StatusBadRequest, "products.html", gin.H{
			"error": err.Error(),
		})
	}

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

	err := services.CreateProduct(&product)
	if err != nil {
		c.HTML(http.StatusBadRequest, "create_product.html", gin.H{
			"error": err.Error(),
		})
	}

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

	services.GetProductById(&product, id)

	c.HTML(http.StatusOK, "update_product.html", gin.H{
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product

	id, _ := strconv.Atoi(c.Param("id"))

	services.GetProductById(&product, id)

	product.Name = c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))

	product.Price = price

	err := services.UpdateProduct(&product)
	if err != nil {
		c.HTML(http.StatusBadRequest, "update_product.html", gin.H{
			"error": err.Error(),
		})
	}

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

	services.GetProductById(&product, id)

	err := services.DeleteProduct(&product)
	if err != nil {
		c.HTML(http.StatusBadRequest, "products.html", gin.H{
			"error": err.Error(),
		})
	}

	helpers.CreateAuditLog(
		c,
		"DELETE",
		"products",
		product.ID,
		"Hapus produk",
	)

	c.Redirect(http.StatusFound, "/product")
}
