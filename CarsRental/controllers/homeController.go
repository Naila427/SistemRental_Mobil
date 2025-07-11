package controllers

import (
	"CarsRental/config" 
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phpdave11/gofpdf"
)

func ExportPDF(c *gin.Context) {
	db, err := config.Connect()
	if err != nil {
		log.Println("Gagal koneksi DB:", err)
		c.String(http.StatusInternalServerError, "Gagal koneksi ke database")
		return
	}

	rows, err := db.Query("SELECT merk, tipe, nomor, status FROM mobil")
	if err != nil {
		log.Println("Gagal query:", err)
		c.String(http.StatusInternalServerError, "Gagal mengambil data")
		return
	}
	defer rows.Close()

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "B", 16)
	pdf.CellFormat(0, 10, "Laporan Data Mobil", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Header tabel
	pdf.SetFont("Helvetica", "B", 12)
	pdf.CellFormat(10, 10, "No", "1", 0, "C", false, 0, "")
	pdf.CellFormat(45, 10, "Merk", "1", 0, "C", false, 0, "")
	pdf.CellFormat(45, 10, "Tipe", "1", 0, "C", false, 0, "")
	pdf.CellFormat(45, 10, "Nomor", "1", 0, "C", false, 0, "")
	pdf.CellFormat(45, 10, "Status", "1", 1, "C", false, 0, "")

	// Isi data
	pdf.SetFont("Helvetica", "", 11)
	no := 1
	for rows.Next() {
		var merk, tipe, nomor, status string
		if err := rows.Scan(&merk, &tipe, &nomor, &status); err != nil {
			log.Println("Gagal scan row:", err)
			continue
		}

		pdf.CellFormat(10, 10, fmt.Sprintf("%d", no), "1", 0, "C", false, 0, "")
		pdf.CellFormat(45, 10, merk, "1", 0, "L", false, 0, "")
		pdf.CellFormat(45, 10, tipe, "1", 0, "L", false, 0, "")
		pdf.CellFormat(45, 10, nomor, "1", 0, "L", false, 0, "")
		pdf.CellFormat(45, 10, status, "1", 1, "L", false, 0, "")
		no++
	}

	// Output PDF ke buffer
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		log.Println("Gagal output PDF:", err)
		c.String(http.StatusInternalServerError, "Gagal membuat PDF")
		return
	}

	// Kirim ke browser
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=mobil.pdf")
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}
