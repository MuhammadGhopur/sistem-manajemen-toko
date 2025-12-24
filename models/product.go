package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID uint   `json:"category_id,omitempty" gorm:"default:0"`
	Category   Category
}
