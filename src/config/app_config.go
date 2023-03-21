package config

import (
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/database"
)

type AppConfig struct {
	HttpPort       int
	DatabaseConfig database.DatabaseConfig
}
