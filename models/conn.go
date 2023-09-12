package models

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitialDatabase() {

	var DB *gorm.DB
	DB, err := gorm.Open(sqlserver.Open("server=localhost,1433; user id=sa; password=@Qorey12"))
	if (err != nil) {
		panic(err)
	}
	fmt.Println("Koneksi berhasil")
	DB.AutoMigrate(&Kurs{})
}