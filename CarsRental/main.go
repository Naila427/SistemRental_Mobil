package main

import (
	"CarsRental/controllers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Fungsi tambahan untuk template (misalnya: penomoran index)
	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int { return a + b },
	})

	r.LoadHTMLGlob("templates/*")

	// Auth routes
	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	// Mobil routes
	r.GET("/mobil", controllers.GetMobil)
	r.POST("/mobil/add", controllers.AddMobil)
	r.POST("/mobil/update", controllers.UpdateMobil)
	r.GET("/mobil/delete/:id", controllers.DeleteMobil)

	// Export PDF
	r.GET("/export-pdf", controllers.ExportPDF)

	// Jalankan server
	r.Run(":8080")
}
