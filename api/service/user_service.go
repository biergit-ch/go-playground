package service

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/model"
	"log"
)

type UserService interface {
	GetAllUsers() []*model.User
	CreateUser(user *model.User)
}

type userService struct {
	repo dao.UserRepository
}

func NewUserService(r dao.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) GetAllUsers() []*model.User {

	users, err := s.repo.FindAll()
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return users
}

func (s *userService) CreateUser(user *model.User) {
	s.repo.Save(user)
}
