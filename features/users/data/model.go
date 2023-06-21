package data

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	FullName		string 		`gorm:"notNull"`
	Email				string 		`gorm:"unique;notNull"`
	Password		string		`gorm:"notNull"`		
	Phone				string		`gorm:"unique:notNull"`
	Address			string		`gorm:"type:text"`
	Avatar			string
	Bank				string		
	NoRekening	string		`gorm:"unique"`
}