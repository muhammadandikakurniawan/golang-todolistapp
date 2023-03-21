package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(config DatabaseConfig) (*gorm.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
