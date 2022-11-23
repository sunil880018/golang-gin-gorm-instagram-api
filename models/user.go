package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId string `gorm:"size:255;not null;unique" json:"userId"`
	Name   string `gorm:"size:255;not null;" json:"username"`
	Email  string `gorm:"size:255;not null;unique" json:"email"`
	DOB    string `gorm:"size:255;not null;" json:"dob"`
}

// ------------------ equals ------------------
// type User struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	UserId      string
// 	Name      string
// 	Email     string
// 	DOB      string
// }

// ----------------------- output -----------------------

// "data": {
// 	"ID": 0,
// 	"CreatedAt": "0001-01-01T00:00:00Z",
// 	"UpdatedAt": "0001-01-01T00:00:00Z",
// 	"DeletedAt": null,
// 	"UserId": "",
// 	"Name": "",
// 	"Email": "",
// 	"DOB": ""
// },
