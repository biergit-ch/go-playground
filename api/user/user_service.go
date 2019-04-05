package user

import "log"

type UserService interface {
	GetAllUsers() []*User
	CreateUser(user *User)
}

type Service struct {
	repo Repository
}

func NewUserService(r Repository) UserService {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetAllUsers() []*User {

	users, err := s.repo.FindAll()
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return users
}

func (s *Service) CreateUser(user *User) {
	s.repo.Save(user)
}
