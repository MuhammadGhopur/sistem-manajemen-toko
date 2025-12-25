package services

import (
	"sistem-manajemen-toko/models"
	"sistem-manajemen-toko/repositories"
)

func CategoryIndex(category *[]models.Category) {
	repositories.CategoryIndex(category)
}
func GetCategoryById(category *models.Category, id int) {
	repositories.GetCategoryById(category, id)
}
func CreateCategory(category *models.Category) {
	repositories.CreateCategory(category)
}
func UpdateCategory(category *models.Category) {
	repositories.UpdateCategory(category)
}
func DeleteCategory(category *models.Category) {
	repositories.DeleteCategory(category)
}
