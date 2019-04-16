package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	err := v.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatal("Fatal error config file: ", err)
	}

	return v
}
