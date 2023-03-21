package todo

import (
	"context"

	pkgModel "github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/model"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/todo/model"
)

type TodoUsecase interface {
	Create(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[model.TodoDto], err error)
	Update(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[*model.TodoDto], err error)
	Delete(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[struct{}], err error)
	GetById(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[*model.TodoDto], err error)
	GetAll(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[[]model.TodoDto], err error)
}
