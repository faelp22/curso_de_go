package config

import "fmt"

type Config struct {
	SRV_PORT string
	WEB_UI   bool
	DBConfig
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

func NewConfig(SRV_PORT string, db_config DBConfig, WEB_UI bool) *Config {

	var conf Config = Config{
		DBConfig: db_config,
	}

	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}

	if db_config.DB_DRIVE != "" {
		conf.DB_DRIVE = db_config.DB_DRIVE
	}

	if db_config.DB_NAME != "" {
		conf.DB_NAME = db_config.DB_NAME
	}

	conf.WEB_UI = WEB_UI

	switch conf.DB_DRIVE {
	case "sqlite":
		conf.DB_DSN = fmt.Sprintf(conf.DB_NAME)

	case "postgresql":
		conf.DB_DSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASS, conf.DB_NAME)
	default:
		panic("Drive não implementado")
	}

	return &conf
}
