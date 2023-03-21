//go:build wireinject
// +build wireinject

// go:build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/config"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/infrastructure/repository/mysql"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/database"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/activity"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/todo"
	"gorm.io/gorm"
)

type Container struct {
	AppConfig       config.AppConfig
	ActivityUsecase activity.ActivityUsecase
	TodoUsecase     todo.TodoUsecase
}

func NewContainer(
	AppConfig config.AppConfig,
	ActivityUsecase activity.ActivityUsecase,
	TodoUsecase todo.TodoUsecase,
) Container {
	return Container{
		AppConfig,
		ActivityUsecase,
		TodoUsecase,
	}
}

func NewMysqlClient(cfg config.AppConfig) *gorm.DB {
	db, err := database.NewMysql(cfg.DatabaseConfig)
	if err != nil {
		panic(err)
	}

	return db
}

func InitializeContainer() Container {
	wire.Build(
		config.LoadConfig,
		NewMysqlClient,
		mysql.NewActivityRepository,
		mysql.NewTodoRepository,
		activity.NewActivityUsecase,
		todo.NewTodoUsecase,
		NewContainer,
	)
	return Container{}
}
