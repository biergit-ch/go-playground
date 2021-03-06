package services

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/repo"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	log "github.com/sirupsen/logrus"
)

type GroupService interface {
	GetAllGroups() []*models.Group
	CreateGroup(group *models.Group)
}

type groupService struct {
	repo repo.GroupRepository
}

func NewGroupService(r repo.GroupRepository) GroupService {
	return &groupService{
		repo: r,
	}
}

func (e *groupService) GetAllGroups() []*models.Group {
	users, err := e.repo.FindAll()
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return users
}

func (e *groupService) CreateGroup(user *models.Group) {
	e.repo.Save(user)
}
