package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func InitDB() {

	log.Println("Initialize database")

	db = Connect()

	// Migrate the Database Schema
	migrateSchema()
}

/**
Generate or Migrate DBMS Schema
 */
func migrateSchema() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
}

func GetCon() *gorm.DB {
	if db == nil || db.Error != nil {
		log.Println("New DB Connection will be established")
		return Connect()
	} else {
		log.Println("Existing DB Connection will be used")
		return db
	}
}

func AddMockData() {

	log.Println("Adding new Person")

	// Add one User
	user1 := models.User{
		FirstName: "Jan",
		LastName:  "Minder",
	}

	user2 := models.User{
		FirstName: "Test",
		LastName:  "User",
	}

	CreateUser(&user1)
	CreateUser(&user2)

	var members []*models.User

	members = append(members, &user1)
	members = append(members, &user2)

	log.Println("append members", members[0].ID, members[0].FirstName, members[0].LastName)
	log.Println("append members", members[1].ID, members[1].FirstName, members[1].LastName)

	// Add one Group
	group := models.Group{
		GroupName: "Juventus",
		Members: members,
	}

	CreateGroup(&group)
}

func Connect() *gorm.DB {

	log.Println("connecting...")
	db, err := gorm.Open("mysql", "go:123@tcp(127.0.0.1:3333)/go-basics?charset=utf8&parseTime=True")

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	return db
}

func Close(db *gorm.DB) {
	db.Close()
}
