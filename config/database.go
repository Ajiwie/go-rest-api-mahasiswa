package config

import (
	"database/sql"
	"fmt"
	"log"
)

func Database() *sql.DB {
	dsn := "root:Mysqlajiwicaksono2005@tcp(localhost:3306)/database_mahasiswa_dosen?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal Koneksi ke Database :", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database Tidak Merespons : err")
	}

	fmt.Println("Database Terkoneksi")
	return db
}
