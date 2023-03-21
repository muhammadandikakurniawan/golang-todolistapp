package config

import (
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/database"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/utility"
	"github.com/spf13/cast"
)

func LoadConfig() AppConfig {
	return AppConfig{
		HttpPort: cast.ToInt(utility.GetEnv("HTTP_PORT", "3030")),
		DatabaseConfig: database.DatabaseConfig{
			Host:     utility.GetRequiredEnv("MYSQL_HOST"),
			Port:     cast.ToInt(utility.GetEnv("MYSQL_PORT", "3306")),
			User:     utility.GetRequiredEnv("MYSQL_USER"),
			Password: utility.GetRequiredEnv("MYSQL_PASSWORD"),
			DbName:   utility.GetRequiredEnv("MYSQL_DBNAME"),
		},
	}
}
