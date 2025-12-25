package controllers

import (
	"net/http"
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/helpers"
	"sistem-manajemen-toko/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CategoryIndex(c *gin.Context) {
	var category []models.Category

	config.DB.Find(&category)
	c.HTML(http.StatusOK, "categories.html", gin.H{
		"categories": category,
	})
}

func CreateCategoryPage(c *gin.Context) {
	c.HTML(http.StatusOK, "create_category.html", nil)
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	category.Nama = c.PostForm("name")

	config.DB.Create(&category)

	helpers.CreateAuditLog(
		c,
		"CREATE",
		"categories",
		category.ID,
		"Menambah kategori baru",
	)

	c.Redirect(http.StatusFound, "/category")
}

func UpdateCategoryPage(c *gin.Context) {

	var category models.Category

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&category, id)

	c.HTML(http.StatusOK, "update_category.html", gin.H{
		"category": category,
	})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&category, id)

	category.Nama = c.PostForm("name")

	config.DB.Save(&category)

	helpers.CreateAuditLog(
		c,
		"UPDATE",
		"categories",
		category.ID,
		"Update kategori",
	)

	c.Redirect(http.StatusFound, "/category")
}

func DeleteCategory(c *gin.Context) {
	var category models.Category

	id, _ := strconv.Atoi(c.Param("id"))

	config.DB.First(&category, id)

	config.DB.Delete(&category)

	helpers.CreateAuditLog(
		c,
		"DELETE",
		"categories",
		category.ID,
		"Hapus kategori",
	)

	c.Redirect(http.StatusFound, "/category")

}
