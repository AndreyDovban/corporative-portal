package configs

import (
	"backend/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Mail MailConfig
}

type DbConfig struct {
	Dsn         string
	FilesFolder string
}

type MailConfig struct {
	Login    string
	Password string
	Host     string
	Port     string
	Protocol string
	Domain   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.ERROR("Error loading .env file? using default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn:         os.Getenv("DSN"),
			FilesFolder: os.Getenv("FILES_FOLDER"),
		},
		Mail: MailConfig{
			Login:    os.Getenv("LOGIN"),
			Password: os.Getenv("PASSWORD"),
			Host:     os.Getenv("HOST"),
			Port:     os.Getenv("PORT"),
			Protocol: os.Getenv("PROTOCOL"),
			Domain:   os.Getenv("DOMAIN"),
		},
	}
}
