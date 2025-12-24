package repositories

import (
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/models"
)

func GetProducts(products *[]models.Product) {
	config.DB.Unscoped().Preload("Category").Find(products)
}

func GetProductById(product *models.Product, id int) {
	config.DB.Unscoped().First(product, id)
}

func CreateProduct(product *models.Product) {
	config.DB.Create(product)
}

func UpdateProduct(product *models.Product) {
	config.DB.Save(product)
}

func DeleteProduct(product *models.Product) {
	config.DB.Unscoped().Delete(product)
}
