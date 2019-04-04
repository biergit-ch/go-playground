package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	IsMature  bool
	Groups []*Group `gorm:"many2many:group_members;PRELOAD:true"`
}
