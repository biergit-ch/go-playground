package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	GroupName 	string
	Members  	[]*User  `gorm:"many2many:group_members;"`
}
