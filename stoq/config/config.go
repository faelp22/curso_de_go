package config

import (
	"os"
	"strconv"
)

const (
	DEVELOPER    = "developer"
	HOMOLOGATION = "homologation"
	PRODUCTION   = "production"
)

type Config struct {
	SRV_PORT string `toml:"srv_port" yaml:"srv_port" json:"srv_port"`
	WEB_UI   bool   `toml:"web_ui" yaml:"web_ui" json:"web_ui"`
	Mode     string `toml:"mode" yaml:"mode" json:"mode"`
	DBConfig `toml:"dbconfig" yaml:"dbconfig" json:"dbconfig"`
}

type DBConfig struct {
	DB_DRIVE string `toml:"db_drive" yaml:"db_drive" json:"db_drive"`
	DB_HOST  string `toml:"db_host" yaml:"db_host" json:"db_host"`
	DB_PORT  string `toml:"db_port" yaml:"db_port" json:"db_port"`
	DB_USER  string `toml:"db_user" yaml:"db_user" json:"db_user"`
	DB_PASS  string `toml:"db_pass" yaml:"db_pass" json:"db_pass"`
	DB_NAME  string `toml:"db_name" yaml:"db_name" json:"db_name"`
	DB_DSN   string `toml:"-" yaml:"-" json:"-"`
}

func NewConfig(confi *Config) *Config {
	var conf *Config

	if confi == nil || confi.SRV_PORT == "" {
		conf = defaultConf()
	} else {
		conf = confi
	}

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	if SRV_WEB_UI != "" {
		conf.WEB_UI, _ = strconv.ParseBool(SRV_WEB_UI)
	}

	SRV_DB_DRIVE := os.Getenv("SRV_DB_DRIVE")
	if SRV_DB_DRIVE != "" {
		conf.DBConfig.DB_DRIVE = SRV_DB_DRIVE
	}

	SRV_DB_HOST := os.Getenv("SRV_DB_HOST")
	if SRV_DB_HOST != "" {
		conf.DBConfig.DB_HOST = SRV_DB_HOST
	}

	SRV_DB_PORT := os.Getenv("SRV_DB_PORT")
	if SRV_DB_PORT != "" {
		conf.DBConfig.DB_PORT = SRV_DB_PORT
	}

	SRV_DB_USER := os.Getenv("SRV_DB_USER")
	if SRV_DB_USER != "" {
		conf.DBConfig.DB_USER = SRV_DB_USER
	}

	SRV_DB_PASS := os.Getenv("SRV_DB_PASS")
	if SRV_DB_PASS != "" {
		conf.DBConfig.DB_PASS = SRV_DB_PASS
	}

	SRV_DB_NAME := os.Getenv("SRV_DB_NAME")
	if SRV_DB_NAME != "" {
		conf.DBConfig.DB_NAME = SRV_DB_NAME
	}

	return conf
}

func defaultConf() *Config {
	default_conf := Config{
		SRV_PORT: "8080",
		WEB_UI:   true,
		DBConfig: DBConfig{
			DB_DRIVE: "sqlite3",
			// DB_DRIVE: "postgres",
			// DB_HOST: "192.168.18.2",
			// DB_PORT: "5432",
			// DB_USER: "postgres",
			// DB_PASS: "123456",
			// DB_NAME: "stoq",
			DB_NAME: "db.sqlite3",
		},
		Mode: PRODUCTION,
	}

	return &default_conf
}
