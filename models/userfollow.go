package models

import "gorm.io/gorm"

type UserFollower struct {
	gorm.Model
	UserId    string `gorm:"size:255;not null;unique" json:"userId"`
	Followers string `gorm:"size:255;not null;" json:"followers"`
}
