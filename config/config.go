package config

import (
	"fmt"
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

func LoadConfig(env string) *Config {

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

	err := v.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	return &c
}
