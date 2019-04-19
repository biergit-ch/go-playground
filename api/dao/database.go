package dao

import (
	"context"
	"git.skydevelopment.ch/zrh-dev/go-basics-backup/api/repo"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao/mariadb"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao/mongodb"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/services"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spf13/viper"
	"net/url"
	"strconv"
)

type dao struct {
}

func EstablishMongoDbConnection(conf *viper.Viper) repo.UserRepository {

	// Check if Mongo DB DSN String is available

	// Start Mongo DB Connection
	client, err := mongo.Connect(context.TODO(), "mongodb://"+conf.GetString("mongodb.host")+":"+strconv.Itoa(conf.GetInt("mongodb.port")))

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	userCollection := client.Database("biergit").Collection("users")
	userRepo := mongodb.NewMongoDbUserRepository(userCollection)

	log.Debug("Connected to MongoDB!")

	return userRepo
}

// Establish a new connection to the configured mariaDB and initiate all repositories
func EstablishMariaDbConnection(conf *viper.Viper) (repo.UserRepository, repo.GroupRepository, repo.TransactionRepository) {

	var db *gorm.DB
	var err error

	// Check if uri is configured
	dsn := conf.GetString("mariadb.dsn")

	log.Debug("DSN: ", dsn)

	if dsn != "" {
		// Parse the DSN String
		dbUrl, _ := url.Parse(dsn)
		log.Debug("QueryString: ", dbUrl.RawQuery)

		// Create new MYSQL Connection
		db, err = gorm.Open(dbUrl.Scheme, dbUrl.User.String()+"@tcp("+
			dbUrl.Host+")"+
			dbUrl.Path+"?charset=utf8&parseTime=True")
		// TODO: Check unknown system variable of reconnect flag

	} else {
		// Create new MYSQL Connection based on configuration file
		db, err = gorm.Open("mysql", conf.GetString("mariadb.user")+":"+
			conf.GetString("mariadb.password")+"@tcp("+
			conf.GetString("mariadb.host")+":"+
			strconv.Itoa(conf.GetInt("mariadb.port"))+")/"+
			conf.GetString("mariadb.schema")+"?charset=utf8&parseTime=True")
	}

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Enable SQL Query Logs
	db.LogMode(true)

	// Migrate Database
	migrateRelationalDB(db)

	// Create Repositories
	userRepo := mariadb.NewMysqlUserRepository(db)
	groupRepo := mariadb.NewMysqlGroupRepository(db)
	transactionRepo := mariadb.NewMysqlTransactionRepository(db)

	return userRepo, groupRepo, transactionRepo
}

// Migrate all Releational Database Tables
func migrateRelationalDB(db *gorm.DB) {
	log.Debug("Migrating Database Schema")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.Transaction{})
}

func AddMockData(userService services.UserService, groupService services.GroupService, transactionService services.TransactionService) {

	log.Debug("Add Mock data")

	// Add some Mock Data
	jan := models.User{
		FirstName: "Jan",
		LastName:  "Minder",
	}

	// Add some Mock Data
	luca := models.User{
		FirstName: "Luca",
		LastName:  "Hostetter",
	}

	// Add some Mock Data
	biergit := models.Group{
		GroupName: "biergit",
	}

	// Add some Mock Data
	bspTrans := models.Transaction{
		Source:  jan,
		Target:  luca,
		Context: biergit,
		Amount:  2,
	}

	log.Debug(bspTrans)

	// Create Services
	userService.CreateUser(&jan)
	groupService.CreateGroup(&biergit)
	transactionService.CreateTransaction(&bspTrans)
}
