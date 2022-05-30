package config

import (
	"fmt"
	"os"
)

type Config struct {
	SRV_PORT string
	WEB_UI   bool
	DBConfig
	Mode string
}

type DBConfig struct {
	DB_DRIVE string
	DB_HOST  string
	DB_PORT  string
	DB_USER  string
	DB_PASS  string
	DB_NAME  string
	DB_DSN   string
}

func NewConfig(SRV_PORT string, db_config DBConfig) *Config {

	var conf Config = Config{
		DBConfig: db_config,
		Mode:     DEVELOPER,
	}

	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	if SRV_WEB_UI != "" {
		conf.WEB_UI = true
	}

	if db_config.DB_DRIVE != "" {
		conf.DB_DRIVE = db_config.DB_DRIVE
	} else {
		conf.DB_DRIVE = "sqlite3"
	}

	if db_config.DB_NAME != "" {
		conf.DB_NAME = db_config.DB_NAME
	} else {
		conf.DB_NAME = "stoq"
	}

	switch conf.DB_DRIVE {
	case "sqlite3":
		conf.DB_DSN = fmt.Sprintf(conf.DB_NAME)

	case "postgresql":
		conf.DB_DSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASS, conf.DB_NAME)
	default:
		panic("Drive não implementado")
	}

	return &conf
}

const (
	DEVELOPER  = "developer"
	PRODUCTION = "production"
)
