package services

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/repo"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	GetAllUsers() []*models.User
	GetUserById(id string) *models.User
	CreateUser(user *models.User) *models.User
	UpdateUser(user *models.User) *models.User
	DeleteUser(id int) error
}

type userService struct {
	repo repo.UserRepository
}

func NewUserService(r repo.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) GetAllUsers() []*models.User {

	users, err := s.repo.FindAll()
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return users
}

func (s *userService) GetUserById(id string) *models.User {
	log.Debug("Get User from Repo with id ", id)
	user, err := s.repo.FindOne(id)
	if err != nil {
		log.Fatal("Failed to get user from repo")
	}
	log.Debug(user)
	return user
}

func (s *userService) CreateUser(user *models.User) *models.User {
	log.WithFields(log.Fields{
		"user_id": user.ID,
	}).Debug("Save user in repository")
	savedUser := s.repo.Save(user)
	return savedUser
}

func (s *userService) UpdateUser(user *models.User) *models.User {

	user, err := s.repo.Update(user)
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return user
}

func (s *userService) DeleteUser(id int) error {

	log.Debug("Delete User from Repo with id ", id)

	error := s.repo.Delete(id)

	return error
}
