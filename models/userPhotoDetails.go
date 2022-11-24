package models

import "gorm.io/gorm"

type UserPhotoDetails struct {
	gorm.Model
	UserId         string `gorm:"size:255;not null;unique" json:"userId"`
	PhotoUrl       string `gorm:"size:255;not null;unique" json:"photourl"`
	PhotoLatitude  int    `gorm:"size:255;not null;unique" json:"photolatitude"`
	PhotoLongitude int    `gorm:"size:255;not null;unique" json:"photolongitude"`
	UserLatitude   int    `gorm:"size:255;not null;unique" json:"userlatitude"`
	UserLongitude  int    `gorm:"size:255;not null;unique" json:"userlongitude"`
}
