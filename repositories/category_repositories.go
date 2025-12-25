package repositories

import (
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/models"
)

func CategoryIndex(category *[]models.Category) {
	config.DB.Find(category)
}
func GetCategoryById(category *models.Category, id int) {
	config.DB.First(category, id)
}
func CreateCategory(category *models.Category) {
	config.DB.Create(category)
}
func UpdateCategory(category *models.Category) {
	config.DB.Save(category)
}
func DeleteCategory(category *models.Category) {
	config.DB.Delete(category)
}
