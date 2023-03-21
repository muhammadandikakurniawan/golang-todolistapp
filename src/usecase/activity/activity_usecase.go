package activity

import (
	"context"

	pkgModel "github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/model"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/activity/model"
)

type ActivityUsecase interface {
	Create(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[model.ActivityDto], err error)
	Update(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[*model.ActivityDto], err error)
	Delete(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[struct{}], err error)
	GetById(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[*model.ActivityDto], err error)
	GetAll(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[[]model.ActivityDto], err error)
}
