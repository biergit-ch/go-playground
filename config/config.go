package config

import (
	"github.com/cloudfoundry-community/go-cfenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	MariaDB mariadb `mapstructure:"mariadb"`
	MontoDB mongodb `mapstructure:"mongodb"`
	Server server `mapstructure:"server"`
	Auth auth0 `mapstructure:"auth0"`
}

type mariadb struct {
	Host   string
	Port     int
	User     string
	Password string
	Schema  string
}

type server struct {
	Host string
	Port int
}

type auth0 struct {
	Issuer string
	Audience string
	Jwks string
}

type mongodb struct {

}

func LoadConfig(env string) *viper.Viper {

	var confFile string

	switch env {
	case "dev":
		confFile = "config.dev"
	case "stag":
		confFile = "config.stag"
	case "prod":
		confFile = "config.prod"
	default:
		confFile = "config.dev"
	}

	// Create new Viper Config Struct
	v := viper.New()
	v.SetConfigName(confFile) // name of config file (without extension)
	v.AddConfigPath("config")

	v.BindEnv("auth0.audience", "AUTH0_AUDIENCE")
	v.BindEnv("auth0.issuer","AUTH0_ISSUER")
	v.BindEnv("auth0.jwks", "AUTH0_JWKS")
	v.BindEnv("server.port", "PORT")

	v.BindEnv("cloud", "VCAP_SERVICES")

	// Configure Database in Heroku
	trans := v.BindEnv("mariadb.jdbc", "JAWSDB_MARIA_URL")
	v.BindEnv("vcapp", "VCAP_SERVICES")

	log.Debug("CF Services: ", v.GetString("vcapp"))

	// Check if Cloud Foundry is used
	if v.GetString("mariadb.jdbc") == "" {

		log.Debug("This application does not run on heroku, try read cloud foundry env")

		// Get CF Env
		appEnv, _ := cfenv.Current()

		log.Debug("CF ENV: ", appEnv)

		// Search Database Config
		dbService, err := appEnv.Services.WithName("biergit-db")

		if err != nil {
			log.Error("Failed to read db Service from cf env")
		} else {
			log.Info("CF DB Servcie: ", dbService)

			// Get DB URI
			uri, ok := dbService.CredentialString("uri")

			if ok {
				v.Set("mariadb.jdbc", uri)
			}
		}
	}

	if trans != nil {
		log.Error(trans)
	}

	log.Debug("Env", v.GetString("mariadb.jdbc"))
	log.Debug("JDBC Env: ", os.Getenv("JAWSDB_MARIA_URL"))

	err := v.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatal("Fatal error config file: ", err)
	}

	return v
}
