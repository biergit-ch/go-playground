package model

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	GroupName string
}
