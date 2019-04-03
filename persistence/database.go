package persistence

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/operations"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbInstance gorm.DB

/**

links:

http://mindbowser.com/golang-go-with-gorm-2/
 */

func SetupDatabase() {
	log.Println("Create a new db table..")

	getActiveConnection().CreateTable(operations.Person{})
	getActiveConnection().CreateTable(operations.Address{})

	// Add Values
	person := operations.Person{
		FirstName: "Jan",
		LastName: "Minder",
		Age: 25,
	}
	SavePerson(person)

}

func getActiveConnection() gorm.DB {
	if dbInstance != nil {
		log.Println("use active connection...")
	} else {
		log.Println("connecting...")
		db, err := gorm.Open("mysql", "go:123@tcp(127.0.0.1:3333)/go-basics?charset=utf8&parseTime=True")
		defer db.Close()
		if err != nil {
			log.Fatal("Failed to connect to database", err)
			db = new(gorm.DB)
		} else {
			dbInstance = db
		}
	}

	return dbInstance
}

func SavePerson(person operations.Person) {
	 getActiveConnection().Create(person)
}
