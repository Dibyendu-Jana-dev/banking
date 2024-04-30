package config

import (
	"banking/constants"
	db "banking/domain"
	"os"
)

type AppConfig struct {
	DB       *db.Config
}

var (
	appConfig     AppConfig
)

func Init() *AppConfig {
	cwpCollection := make(map[string]string)
	cwpCollection["findings_cwp"] = constants.FINDINGS_CWP_COLLECTION
	appConfig = AppConfig{
		DB: &db.Config{
			Host:                   os.Getenv("DB_HOST"),
			Port:                   os.Getenv("DB_PORT"),
			Username:               os.Getenv("DB_USERNAME"),
			Password:               os.Getenv("DB_PASSWORD"),
			MaxPool:                os.Getenv("MAX_POOL"),
			Database:               os.Getenv("DB_NAME"),
			CwpCollection:          cwpCollection,
		},
	}
	return &appConfig
}

// Config gets the app environment config
func Config() *AppConfig {
	return &appConfig
}