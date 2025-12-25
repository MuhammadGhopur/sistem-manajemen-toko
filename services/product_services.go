package services

import (
	"errors"
	"sistem-manajemen-toko/models"
	"sistem-manajemen-toko/repositories"
)

func GetProduct(products *[]models.Product) error {
	if len(*products) == 0 {
		return errors.New("Produk tidak ditemukan")
	}

	err := repositories.GetProducts(products)
	if err != nil {
		return err
	}
	return nil
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

	err := repositories.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
func UpdateProduct(product *models.Product) error {

	if product.Name == "" {
		return errors.New("Nama produk tidak boleh kosong")
	}

	if product.Price <= 0 {
		return errors.New("Harga harus lebih dari 0")
	}

	err := repositories.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
func DeleteProduct(product *models.Product) error {

	err := repositories.DeleteProduct(product)
	if err != nil {
		return err
	}
	return nil
}
