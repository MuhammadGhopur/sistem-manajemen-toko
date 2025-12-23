package models

import "gorm.io/gorm"

type AuditLog struct {
	gorm.Model

	UserId   uint
	Username string
	Role     string

	Action string
	Table  string
	DataID uint

	Description string
}
