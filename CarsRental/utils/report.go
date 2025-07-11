package utils

import (
	"CarsRental/config"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF() *gofpdf.Fpdf {
	db, _ := config.Connect()
	rows, _ := db.Query("SELECT merk, tipe, nomor, status FROM mobil")
	defer rows.Close()

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Laporan Data Mobil")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	for rows.Next() {
		var merk, tipe, nomor, status string
		rows.Scan(&merk, &tipe, &nomor, &status)
		pdf.Cell(40, 10, merk)
		pdf.Cell(40, 10, tipe)
		pdf.Cell(40, 10, nomor)
		pdf.Cell(40, 10, status)
		pdf.Ln(10)
	}
	return pdf
}
