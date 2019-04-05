package group

import "log"

type GroupService interface {
	GetAllGroups() []*Group
	CreateGroup(group *Group)
}

type Service struct {
	repo Repository
}

func NewGroupService(r Repository) GroupService {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetAllGroups() []*Group {

	users, err := s.repo.FindAll()
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return users
}

func (s *Service) CreateGroup(user *Group) {
	s.repo.Save(user)
}
