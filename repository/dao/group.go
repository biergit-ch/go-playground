package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"log"
)

func GetGroup() models.Group {
	var group models.Group
	GetCon().First(&group)
	return group
}

func CreateGroup(group *models.Group) models.Group {
	var createdGroup models.Group
	log.Println("Groupname", group.GroupName)
	GetCon().Create(group)
	GetCon().First(&createdGroup)
	return createdGroup
}
