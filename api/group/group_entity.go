package group

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	GroupName string
}
