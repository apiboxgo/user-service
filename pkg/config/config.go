package config

import (
	"github.com/joho/godotenv"
	"os"
)

const DefaultEnv = "prod"

type Config struct {
	ServerName   string
	ServerDomain string
	ServerPort   string
	Env          string
	DbHost       string
	DbUser       string
	DbPassword   string
	DbName       string
	DbPort       string
	DbSSLMode    string
	DbDriver     string
}

func GetEnv() string {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = DefaultEnv
	}

	return env
}

func (config *Config) InitConfig(path string) error {
	config.Env = GetEnv()

	err := godotenv.Load(path+".env", path+".env."+config.Env)
	if err != nil {
		return err
	}

	config.ServerName = os.Getenv("SERVER_NAME")
	config.ServerDomain = os.Getenv("SERVER_DOMAINE")
	config.ServerPort = os.Getenv("SERVER_PORT")

	config.DbHost = os.Getenv("DB_HOST")
	config.DbUser = os.Getenv("DB_USER")
	config.DbPassword = os.Getenv("DB_PASSWORD")
	config.DbName = os.Getenv("DB_NAME")
	config.DbPort = os.Getenv("DB_PORT")
	config.DbSSLMode = os.Getenv("DB_SSL_MODE")
	config.DbDriver = os.Getenv("DB_DRIVER")
	return err
}
