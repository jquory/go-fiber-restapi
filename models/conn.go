package models

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)
var DB *gorm.DB
func InitialDatabase() {
	db, err := gorm.Open(sqlserver.Open("server=localhost,1433; user id=sa; password=@Qorey12; database=product"))
	if (err != nil) {
		panic(err)
	}
	fmt.Println("Koneksi berhasil")
	db.AutoMigrate(&Kurs{})

	DB = db
}