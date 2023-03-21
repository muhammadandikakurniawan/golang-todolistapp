package repository

import (
	"context"

	"github.com/muhammadandikakurniawan/golang-todolistapp/src/entity"
)

type TodoRepository interface {
	Insert(ctx context.Context, data *entity.Todo) (err error)
	Update(ctx context.Context, data *entity.TodoPartialUpdateModel) (err error)
	Delete(ctx context.Context, id int64) (err error)
	FindById(ctx context.Context, id int64) (result *entity.Todo, err error)
	FindAll(ctx context.Context, filter entity.TodoPartialUpdateModel) (result []entity.Todo, err error)
}
