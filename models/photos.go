package models

import "gorm.io/gorm"

type Photos struct {
	gorm.Model
	PhotoId  string `gorm:"size:255;not null;unique" json:"imageid"`
	PhotoUrl string `gorm:"size:255;not null;unique" json:"photourl"`
	Title    string `gorm:"size:255;not null;unique" json:"title"`
}
