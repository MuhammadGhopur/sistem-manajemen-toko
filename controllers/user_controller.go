package controllers

import (
	"net/http"
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}
func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// hash
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.String(http.StatusInternalServerError, "Gagal hash password!")
		return
	}

	// simpan ke database
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	config.DB.Create(&user)

	c.Redirect(http.StatusFound, "/login")

}
