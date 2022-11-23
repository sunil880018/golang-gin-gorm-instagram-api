package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	UserId         string `gorm:"size:255;not null;unique" json:"userId"`
	PhotoPath      string `gorm:"size:255;not null;unique" json:"photopath"`
	PhotoLatitude  int    `gorm:"size:255;not null;unique" json:"photolatitude"`
	PhotoLongitude int    `gorm:"size:255;not null;unique" json:"photolongitude"`
	UserLatitude   int    `gorm:"size:255;not null;unique" json:"userlatitude"`
	UserLongitude  int    `gorm:"size:255;not null;unique" json:"userlongitude"`
}
