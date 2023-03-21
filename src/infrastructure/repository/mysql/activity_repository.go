package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/muhammadandikakurniawan/golang-todolistapp/src/entity"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/repository"
	"gorm.io/gorm"
)

func NewActivityRepository(db *gorm.DB) repository.ActivityRepository {

	return &ActivityRepository{
		db: db,
	}
}

type ActivityRepository struct {
	db *gorm.DB
}

func (repo ActivityRepository) Insert(ctx context.Context, data *entity.Activity) (err error) {
	if data == nil {
		return
	}
	err = repo.db.Create(data).Error
	return
}

func (repo ActivityRepository) Update(ctx context.Context, data *entity.ActivityPartialUpdateModel) (err error) {
	if data == nil {
		return
	}

	updatedData := map[string]interface{}{}

	if data.Email != nil {
		updatedData["email"] = *data.Email
	}

	if data.Title != nil {
		updatedData["title"] = *data.Title
	}

	dataIsNotUpdated := len(updatedData) <= 0
	if dataIsNotUpdated {
		return
	}

	data.UpdatedAt = time.Now()
	updatedData["updated_at"] = data.UpdatedAt

	updateRes := repo.db.Model(entity.Activity{}).Where("activity_id = ?", data.ActivityId).Updates(updatedData)
	if err = updateRes.Error; err != nil {
		return
	}
	if updateRes.RowsAffected <= 0 {
		err = sql.ErrNoRows
	}
	return
}

func (repo ActivityRepository) Delete(ctx context.Context, id int64) (err error) {
	deleteRes := repo.db.Model(entity.Activity{}).Where("activity_id = ?", id).Delete(&entity.Activity{})
	if err = deleteRes.Error; err != nil {
		return
	}
	if deleteRes.RowsAffected <= 0 {
		err = sql.ErrNoRows
	}
	return
}

func (repo ActivityRepository) FindById(ctx context.Context, id int64) (result *entity.Activity, err error) {
	tempData := entity.Activity{}
	err = repo.db.Where("activity_id = ?", id).First(&tempData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return
	}
	result = &tempData
	return
}

func (repo ActivityRepository) FindAll(ctx context.Context, filter entity.ActivityPartialUpdateModel) (result []entity.Activity, err error) {
	err = repo.db.Find(&result).Error

	query := repo.db
	if filter.Title != nil {
		query = repo.db.Where("title LIKE ?", "%"+*filter.Title+"%")
	}
	if filter.Email != nil {
		query = repo.db.Where("email = ?", *filter.Email)
	}

	if err = query.Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}
