package services

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	GetAllUsers() []*models.User
	GetUserById(id int64) []*models.User
	CreateUser(user *models.User) *models.User
	UpdateUser(user *models.User) *models.User
	DeleteUser(id int64) error
}

type userService struct {
	repo dao.UserRepository
}

func NewUserService(r dao.UserRepository) UserService {
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

func (s *userService) GetUserById(id int64) []*models.User {
	log.Debug("Get User from Repo with id ", id)
	users, err := s.repo.Find(id)
	if err != nil {
		log.Fatal("Failed to get user from repo")
	}
	log.Debug(users)
	return users
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

func (s *userService) DeleteUser(id int64) error {

	log.Debug("Delete User from Repo with id ", id)

	error := s.repo.Delete(id)

	return error
}
