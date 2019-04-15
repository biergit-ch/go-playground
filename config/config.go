package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DB database `mapstructure:"database"`
	Server server `mapstructure:"server"`
	Auth auth0 `mapstructure:"auth0"`
}

type database struct {
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

	v := viper.New()
	v.SetConfigName(confFile) // name of config file (without extension)
	v.AddConfigPath("config")

	v.SetEnvPrefix("BG")

	v.BindEnv("auth0.audience", "AUTH0_AUDIENCE")
	v.BindEnv("auth0.issuer","AUTH0_ISSUER")
	v.BindEnv("auth0.jwks", "AUTH0_JWKS")

	err := v.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatal("Fatal error config file: ", err)
	}

	return v
}
