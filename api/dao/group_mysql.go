package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

type GroupRepository interface {
	FindAll() ([]*model.Group, error)
	Save(user *model.Group)
}

type groupRepository struct {
	db *gorm.DB
}

func NewMysqlGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{
		db: db,
	}
}

func (r *groupRepository) FindAll() ([]*model.Group, error) {

	var groups []*model.Group
	r.db.Find(&groups)

	return groups, r.db.Error
}

func (r *groupRepository) Save(group *model.Group) {
	r.db.Create(&group)
}
