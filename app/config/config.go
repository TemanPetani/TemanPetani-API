package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	JWT_SECRET = ""
)

type AppConfig struct {
	DB_USERNAME    string
	DB_PASS        string
	DB_HOSTNAME    string
	DB_PORT        int
	DB_NAME        string
	JWT_SECRET_KEY string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	config := AppConfig{}
	isRead := false

	if val, found := os.LookupEnv("DB_USERNAME"); found {
		config.DB_USERNAME = val
		isRead = true
	}
	if val, found := os.LookupEnv("DB_PASS"); found {
		config.DB_PASS = val
		isRead = true
	}
	if val, found := os.LookupEnv("DB_HOSTNAME"); found {
		config.DB_HOSTNAME = val
		isRead = true
	}
	if val, found := os.LookupEnv("DB_PORT"); found {
		config.DB_PORT, _ = strconv.Atoi(val)
		isRead = true
	}
	if val, found := os.LookupEnv("DB_NAME"); found {
		config.DB_NAME = val
		isRead = true
	}
	if val, found := os.LookupEnv("JWT_SECRET_KEY"); found {
		config.JWT_SECRET_KEY = val
		JWT_SECRET = config.JWT_SECRET_KEY
		isRead = true
	}

	if !isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("error load config: " + err.Error())
		}

		config.DB_USERNAME = viper.GetString("DB_USERNAME")
		config.DB_PASS = viper.GetString("DB_PASS")
		config.DB_HOSTNAME = viper.GetString("DB_HOSTNAME")
		config.DB_PORT, _ = strconv.Atoi(viper.GetString("DB_PORT"))
		config.DB_NAME = viper.GetString("DB_NAME")
		config.JWT_SECRET_KEY = viper.GetString("JWT_SECRET_KEY")
	}

	return &config
}
