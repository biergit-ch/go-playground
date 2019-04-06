package service

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/model"
	"log"
)

type GroupService interface {
	GetAllGroups() []*model.Group
	CreateGroup(group *model.Group)
}

type groupService struct {
	repo dao.GroupRepository
}

func NewGroupService(r dao.GroupRepository) GroupService {
	return &groupService{
		repo: r,
	}
}

func (e *groupService) GetAllGroups() []*model.Group {
	users, err := e.repo.FindAll()
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return users
}

func (e *groupService) CreateGroup(user *model.Group) {
	e.repo.Save(user)
}
