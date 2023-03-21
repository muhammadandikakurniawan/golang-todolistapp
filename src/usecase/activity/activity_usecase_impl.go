package activity

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	pkgModel "github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/model"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/repository"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/activity/model"
)

func NewActivityUsecase(activityRepository repository.ActivityRepository) ActivityUsecase {
	if activityRepository == nil {
		panic("activityRepository is null")
	}

	return &activityUsecaseImpl{
		activityRepository: activityRepository,
	}
}

type activityUsecaseImpl struct {
	activityRepository repository.ActivityRepository
}

func (uc activityUsecaseImpl) Create(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[model.ActivityDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	newActivityEntity, _ := param.TransformToNewEntity()
	err = uc.activityRepository.Insert(ctx, &newActivityEntity)
	if err != nil {
		return
	}

	result.Status = "Success"
	result.Message = "Success"
	result.HttpStatusCode = http.StatusOK
	result.Data = model.SetupActivityDtoFromEntity(newActivityEntity)
	return
}

func (uc activityUsecaseImpl) Update(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[*model.ActivityDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	if param.Id <= 0 {
		result.Set("Not Found", fmt.Sprintf("Activity with ID %d Not Found", param.Id), http.StatusNotFound, nil)
		return
	}

	partialUpdateData := param.TransformToPartialUpdate()
	err = uc.activityRepository.Update(ctx, &partialUpdateData)
	if err != nil {
		if err == sql.ErrNoRows {
			result.Set("Not Found", fmt.Sprintf("Activity with ID %d Not Found", param.Id), http.StatusNotFound, nil)
			err = nil
		}
		return
	}

	activityData, err := uc.activityRepository.FindById(ctx, param.Id)
	if err != nil {
		return
	}
	if activityData == nil {
		result.Set("Not Found", fmt.Sprintf("Activity with ID %d Not Found", param.Id), http.StatusNotFound, nil)
		return
	}

	result.Status = "Success"
	result.Message = "Success"
	result.HttpStatusCode = http.StatusOK
	resultData := model.SetupActivityDtoFromEntity(*activityData)
	result.Data = &resultData
	return
}

func (uc activityUsecaseImpl) Delete(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[struct{}], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	if param.Id <= 0 {
		result.Set("Not Found", fmt.Sprintf("Activity with ID %d Not Found", param.Id), http.StatusNotFound, struct{}{})
		return
	}

	// delete
	err = uc.activityRepository.Delete(ctx, param.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			result.Set("Not Found", fmt.Sprintf("Activity with ID %d Not Found", param.Id), http.StatusNotFound, struct{}{})
			err = nil
		}
		return
	}

	result.Status = "Success"
	result.Message = "Success"
	result.HttpStatusCode = http.StatusOK
	return
}

func (uc activityUsecaseImpl) GetById(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[*model.ActivityDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	if param.Id <= 0 {
		result.Set("Not Found", fmt.Sprintf("Activity with ID %d Not Found", param.Id), http.StatusNotFound, nil)
		return
	}

	activityData, err := uc.activityRepository.FindById(ctx, param.Id)
	if err != nil {
		return
	}
	if activityData == nil {
		result.Set("Not Found", fmt.Sprintf("Activity with ID %d Not Found", param.Id), http.StatusNotFound, nil)
		return
	}

	result.Status = "Success"
	result.Message = "Success"
	resultData := model.SetupActivityDtoFromEntity(*activityData)
	result.Data = &resultData
	result.HttpStatusCode = http.StatusOK
	return
}

func (uc activityUsecaseImpl) GetAll(ctx context.Context, param model.ActivityDto) (result pkgModel.BaseResponseModel[[]model.ActivityDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	filter := param.TransformToPartialUpdate()
	activites, err := uc.activityRepository.FindAll(ctx, filter)
	if err != nil {
		return
	}

	for _, act := range activites {
		result.Data = append(result.Data, model.SetupActivityDtoFromEntity(act))
	}
	result.Status = "Success"
	result.Message = "Success"
	result.HttpStatusCode = http.StatusOK
	return
}
