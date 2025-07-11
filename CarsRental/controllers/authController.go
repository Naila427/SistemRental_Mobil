package controllers

import (
	"CarsRental/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	db, _ := config.Connect()

	var id int
	err := db.QueryRow("SELECT id FROM user WHERE username=? AND password=?", username, password).Scan(&id)

	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Login gagal"})
		return
	}

	c.SetCookie("user_id", "1", 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/mobil")
}

func Logout(c *gin.Context) {
	c.SetCookie("user_id", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/login")
}
