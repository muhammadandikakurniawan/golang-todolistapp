package todo

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	pkgModel "github.com/muhammadandikakurniawan/golang-todolistapp/src/pkg/model"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/repository"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/usecase/todo/model"
)

func NewTodoUsecase(
	activityRepository repository.ActivityRepository,
	todoRepository repository.TodoRepository,
) TodoUsecase {
	return &todoUsecaseImpl{
		activityRepository: activityRepository,
		todoRepository:     todoRepository,
	}
}

type todoUsecaseImpl struct {
	activityRepository repository.ActivityRepository
	todoRepository     repository.TodoRepository
}

func (uc todoUsecaseImpl) Create(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[model.TodoDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	newActivityEntity, validationErrMsg := param.TransformToNewEntity()
	invalidRequest := len(validationErrMsg) > 0
	if invalidRequest {
		errMsg := strings.Join(validationErrMsg, ", ")
		result.Set("Bad Request", errMsg, http.StatusBadRequest, model.TodoDto{})
		return
	}

	err = uc.todoRepository.Insert(ctx, &newActivityEntity)
	if err != nil {
		return
	}

	result.Status = "Success"
	result.Message = "Success"
	result.HttpStatusCode = http.StatusOK
	result.Data = model.SetupTodoDtoFromEntity(newActivityEntity)
	return
}

func (uc todoUsecaseImpl) Update(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[*model.TodoDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	if param.Id <= 0 {
		result.Set("Not Found", fmt.Sprintf("Todo with ID %d Not Found", param.Id), http.StatusNotFound, nil)
		return
	}

	partialUpdateData, validationErrMsg := param.TransformToPartialUpdate()
	invalidRequest := len(validationErrMsg) > 0
	if invalidRequest {
		errMsg := strings.Join(validationErrMsg, ", ")
		result.Set("Bad Request", errMsg, http.StatusBadRequest, nil)
		return
	}

	err = uc.todoRepository.Update(ctx, &partialUpdateData)
	if err != nil {
		if err == sql.ErrNoRows {
			result.Set("Not Found", fmt.Sprintf("Todo with ID %d Not Found", param.Id), http.StatusNotFound, nil)
			err = nil
		}
		return
	}

	tdoData, err := uc.todoRepository.FindById(ctx, param.Id)
	if err != nil {
		return
	}
	if tdoData == nil {
		result.Set("Not Found", fmt.Sprintf("Todo with ID %d Not Found", param.Id), http.StatusNotFound, nil)
		return
	}

	result.Status = "Success"
	result.Message = "Success"
	result.HttpStatusCode = http.StatusOK
	resultData := model.SetupTodoDtoFromEntity(*tdoData)
	result.Data = &resultData
	return
}

func (uc todoUsecaseImpl) Delete(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[struct{}], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	if param.Id <= 0 {
		result.Set("Not Found", fmt.Sprintf("Todo with ID %d Not Found", param.Id), http.StatusNotFound, struct{}{})
		return
	}

	// delete
	err = uc.todoRepository.Delete(ctx, param.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			result.Set("Not Found", fmt.Sprintf("Todo with ID %d Not Found", param.Id), http.StatusNotFound, struct{}{})
			err = nil
		}
		return
	}

	result.Status = "Success"
	result.Message = "Success"
	result.HttpStatusCode = http.StatusOK
	return
}

func (uc todoUsecaseImpl) GetById(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[*model.TodoDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	if param.Id <= 0 {
		result.Set("Not Found", fmt.Sprintf("Todo with ID %d Not Found", param.Id), http.StatusNotFound, nil)
		return
	}

	todoData, err := uc.todoRepository.FindById(ctx, param.Id)
	if err != nil {
		return
	}
	if todoData == nil {
		result.Set("Not Found", fmt.Sprintf("Todo with ID %d Not Found", param.Id), http.StatusNotFound, nil)
		return
	}

	result.Status = "Success"
	result.Message = "Success"
	resultData := model.SetupTodoDtoFromEntity(*todoData)
	result.Data = &resultData
	result.HttpStatusCode = http.StatusOK
	return
}

func (uc todoUsecaseImpl) GetAll(ctx context.Context, param model.TodoDto) (result pkgModel.BaseResponseModel[[]model.TodoDto], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	filter, _ := param.TransformToPartialUpdate()
	activites, err := uc.todoRepository.FindAll(ctx, filter)
	if err != nil {
		return
	}

	for _, act := range activites {
		result.Data = append(result.Data, model.SetupTodoDtoFromEntity(act))
	}
	result.Status = "Success"
	result.Message = "Success"
	result.HttpStatusCode = http.StatusOK
	return
}
