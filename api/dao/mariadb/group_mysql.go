package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao/interface"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type groupRepository struct {
	db *gorm.DB
}

func NewMysqlGroupRepository(db *gorm.DB) _interface.GroupRepository {
	return &groupRepository{
		db: db,
	}
}

func (r *groupRepository) FindAll() ([]*models.Group, error) {

	var groups []*models.Group
	r.db.Find(&groups)

	return groups, r.db.Error
}

func (r *groupRepository) Save(group *models.Group) {
	r.db.Create(&group)
}
