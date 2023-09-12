package models

import (
	mssql "github.com/microsoft/go-mssqldb"
)

type Kurs struct {
	Id mssql.UniqueIdentifier `gorm:"type:uniqueidentifier;primary_key" json:"id"`
	Name string `gorm:"type:varchar(300)" json:"name"`
}