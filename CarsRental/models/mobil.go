package models

type Mobil struct {
	ID     int
	Merk   string
	Tipe   string
	Nomor  string
	Status string
}

func (Mobil) TableName() string {
	return "mobil"
}
