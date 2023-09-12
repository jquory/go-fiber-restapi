package models

type Kurs struct {
	Id int8 `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(300)" json:"name"`
}