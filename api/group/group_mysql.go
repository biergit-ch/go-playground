package group

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

type repo struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) FindAll() ([]*Group, error) {

	var groups []*Group
	r.db.Find(&groups)

	return groups, r.db.Error
}

func (r *repo) Save(group *Group) {
	r.db.Create(&group)
}
