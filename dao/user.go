package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
)

func GetUser() models.User {
	var user models.User
	GetCon().Find(&user)
	return user
}

func GetUsers() []models.User {
	var users []models.User
	GetCon().Find(&users)
	return users
}

func CreateUser(person *models.User) models.User {
	GetCon().Create(person)
	var user models.User
	GetCon().First(&user)
	return user
}
