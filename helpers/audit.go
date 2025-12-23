package helpers

import (
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAuditLog(
	c *gin.Context,
	action string,
	table string,
	dataID uint,
	description string,
) {
	userIDstr, _ := c.Cookie("user_id")
	username, _ := c.Cookie("username")
	role, _ := c.Cookie("role")

	userID, _ := strconv.Atoi(userIDstr)

	log := models.AuditLog{
		UserId:      uint(userID),
		Username:    username,
		Role:        role,
		Action:      action,
		Table:       table,
		DataID:      dataID,
		Description: description,
	}

	config.DB.Create(&log)
}
