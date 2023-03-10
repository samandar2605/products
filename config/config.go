package config

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	PostConfig PostgresConfig
	HttpPort   string
	SMTP       Smtp
	SecretKey  string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Database string
	Password string
}

type Smtp struct {
	Sender   string
	Password string
}

func Load(path string) Config {
	gotenv.Load(path + "/.env")

	Conf := viper.New()
	Conf.AutomaticEnv()
	cfg := Config{
		HttpPort: Conf.GetString("HTTP_PORT"),
		PostConfig: PostgresConfig{
			Host:     Conf.GetString("POSTGRES_HOST"),
			Port:     Conf.GetString("POSTGRES_PORT"),
			User:     Conf.GetString("POSTGRES_USER"),
			Database: Conf.GetString("POSTGRES_DATABASE"),
			Password: Conf.GetString("POSTGRES_PASSWORD"),
		},
		SMTP: Smtp{
			Sender:   Conf.GetString("SMTP_SENDER"),
			Password: Conf.GetString("SMTP_PASSWORD"),
		},
		SecretKey: Conf.GetString("SECRET_KEY"),
	}
	return cfg
}
