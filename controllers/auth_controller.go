package controllers

import (
	"net/http"
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/models"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginProcess(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User

	err := config.DB.
		Where("username = ?", username).
		First(&user).Error

	// ❌ user tidak ditemukan
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "username atau password salah",
		})
		return // 🔥 INI WAJIB
	}

	// ❌ password salah
	if user.Password != password {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "username atau password salah",
		})
		return // 🔥 INI WAJIB
	}

	// ✅ login berhasil
	c.SetCookie("login", "true", 3600, "/", "", false, true)

	c.Redirect(http.StatusFound, "/category")
	return // 🔥 INI JUGA WAJIB
}

func Logout(c *gin.Context) {
	c.SetCookie("login", "", -1, "/", "", false, true)

	c.Redirect(http.StatusFound, "/login")
}
