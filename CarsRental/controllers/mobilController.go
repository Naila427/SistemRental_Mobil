package controllers

import (
	"CarsRental/config"
	"CarsRental/helper"
	"CarsRental/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMobil(c *gin.Context) {
	db, _ := config.Connect()
	rows, _ := db.Query("SELECT * FROM mobil")
	defer rows.Close()

	var list []models.Mobil
	for rows.Next() {
		var m models.Mobil
		rows.Scan(&m.ID, &m.Merk, &m.Tipe, &m.Nomor, &m.Status)
		list = append(list, m)
	}

	var editData models.Mobil
	idEdit := c.Query("edit")
	if idEdit != "" {
		db.QueryRow("SELECT * FROM mobil WHERE id = ?", idEdit).Scan(
			&editData.ID, &editData.Merk, &editData.Tipe, &editData.Nomor, &editData.Status)
	}

	c.HTML(http.StatusOK, "mobil.html", gin.H{
		"mobil": list,
		"edit":  idEdit != "",
		"data":  editData,
	})
}

func AddMobil(c *gin.Context) {
	db, _ := config.Connect()
	nomor := c.PostForm("nomor")

	if helper.IsNomorExist(db, nomor) {
		c.String(400, "Nomor polisi sudah terdaftar!")
		return
	}

	db.Exec("INSERT INTO mobil (merk, tipe, nomor, status) VALUES (?, ?, ?, ?)",
		c.PostForm("merk"), c.PostForm("tipe"), nomor, c.PostForm("status"))
	c.Redirect(http.StatusFound, "/mobil")
}

func UpdateMobil(c *gin.Context) {
	db, _ := config.Connect()
	db.Exec("UPDATE mobil SET merk=?, tipe=?, nomor=?, status=? WHERE id=?",
		c.PostForm("merk"), c.PostForm("tipe"), c.PostForm("nomor"), c.PostForm("status"), c.PostForm("id"))
	c.Redirect(http.StatusFound, "/mobil")
}

func DeleteMobil(c *gin.Context) {
	db, _ := config.Connect()
	db.Exec("DELETE FROM mobil WHERE id=?", c.Param("id"))
	c.Redirect(http.StatusFound, "/mobil")
}
