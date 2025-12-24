package controllers

import (
	"fmt"
	"net/http"
	"sistem-manajemen-toko/config"
	"sistem-manajemen-toko/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	if err != nil {
		fmt.Println("User not found or DB error:", err)
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "username atau password salah",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password mismatch after bcrypt comparison!")
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "username atau password salah",
		})
		return
	}

	c.SetCookie("login", "true", 3600, "/", "", false, true)
	c.SetCookie("role", user.Role, 3600, "/", "", false, true)
	c.SetCookie("user_id", strconv.Itoa(int(user.ID)), 3600, "/", "", false, true)
	c.SetCookie("username", user.Username, 3600, "/", "", false, true)
	c.SetCookie("role", user.Role, 3600, "/", "", false, true)

	c.Redirect(http.StatusFound, "/category")
	return
}

func Logout(c *gin.Context) {
	c.SetCookie("login", "", -1, "/", "", false, true)

	c.Redirect(http.StatusFound, "/login")
}
