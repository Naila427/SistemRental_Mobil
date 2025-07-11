package helper

import "database/sql"

func IsNomorExist(db *sql.DB, nomor string) bool {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM mobil WHERE nomor = ?", nomor).Scan(&count)
	return count > 0
}
