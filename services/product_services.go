package services

import (
	"errors"
	"sistem-manajemen-toko/models"
	"sistem-manajemen-toko/repositories"
)

func GetProduct(products *[]models.Product) {
	repositories.GetProducts(products)
}
func GetProductById(product *models.Product, id int) {
	repositories.GetProductById(product, id)
}
func CreateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("Nama produk harus diisi")
	}
	if product.Price <= 0 {
		return errors.New("Harga tidak boleh nol")
	}
	if product.CategoryID == 0 {
		return errors.New("Isi kategori")
	}
	if product.Price <= 1000 {
		return errors.New("Harga harus lebih dari 1000")
	}

	repositories.CreateProduct(product)
	return nil
}
func UpdateProduct(product *models.Product) error {

	if product.Name == "" {
		return errors.New("Nama produk tidak boleh kosong")
	}

	if product.Price <= 0 {
		return errors.New("Harga harus lebih dari 0")
	}

	repositories.UpdateProduct(product)
	return nil
}
func DeleteProduct(product *models.Product) {
	repositories.DeleteProduct(product)
}
