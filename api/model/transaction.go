package model

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Context Group `gorm:"foreignkey:ContextID"`
	Source  User `gorm:"foreignkey:SourceID"`
	Target  User `gorm:"foreignkey:TargetID"`
	Amount  int
	SourceID int `json:"-"`
	TargetID int `json:"-"`
	ContextID int `json:"-"`
}
