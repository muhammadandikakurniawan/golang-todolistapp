package repository

import (
	"context"

	"github.com/muhammadandikakurniawan/golang-todolistapp/src/entity"
)

type ActivityRepository interface {
	Insert(ctx context.Context, data *entity.Activity) (err error)
	Update(ctx context.Context, data *entity.ActivityPartialUpdateModel) (err error)
	Delete(ctx context.Context, id int64) (err error)
	FindById(ctx context.Context, id int64) (result *entity.Activity, err error)
	FindAll(ctx context.Context, filter entity.ActivityPartialUpdateModel) (result []entity.Activity, err error)
}
