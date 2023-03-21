package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/muhammadandikakurniawan/golang-todolistapp/src/entity"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/repository"
	"gorm.io/gorm"
)

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {

	return &TodoRepository{
		db: db,
	}
}

type TodoRepository struct {
	db *gorm.DB
}

func (repo TodoRepository) Insert(ctx context.Context, data *entity.Todo) (err error) {
	if data == nil {
		return
	}
	err = repo.db.Create(data).Error
	return
}

func (repo TodoRepository) Update(ctx context.Context, data *entity.TodoPartialUpdateModel) (err error) {
	if data == nil {
		return
	}

	updatedData := map[string]interface{}{}

	if data.Title != nil {
		updatedData["title"] = *data.Title
	}

	if data.Priority != nil {
		updatedData["priority"] = *data.Priority
	}

	if data.IsActive != nil {
		updatedData["is_active"] = *data.IsActive
	}

	dataIsNotUpdated := len(updatedData) <= 0
	if dataIsNotUpdated {
		return
	}

	data.UpdatedAt = time.Now()
	updatedData["updated_at"] = data.UpdatedAt

	updateRes := repo.db.Model(entity.Todo{}).Where("todo_id = ?", data.TodoId).Updates(updatedData)
	if err = updateRes.Error; err != nil {
		return
	}
	if updateRes.RowsAffected <= 0 {
		err = sql.ErrNoRows
	}
	return
}

func (repo TodoRepository) Delete(ctx context.Context, id int64) (err error) {
	deleteRes := repo.db.Model(entity.Todo{}).Where("todo_id = ?", id).Delete(&entity.Todo{})
	if err = deleteRes.Error; err != nil {
		return
	}
	if deleteRes.RowsAffected <= 0 {
		err = sql.ErrNoRows
	}
	return
}

func (repo TodoRepository) FindById(ctx context.Context, id int64) (result *entity.Todo, err error) {
	tempData := entity.Todo{}
	err = repo.db.Where("todo_id = ?", id).First(&tempData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return
	}
	result = &tempData
	return
}

func (repo TodoRepository) FindAll(ctx context.Context, filter entity.TodoPartialUpdateModel) (result []entity.Todo, err error) {
	err = repo.db.Find(&result).Error

	query := repo.db
	if filter.Title != nil {
		query = repo.db.Where("title LIKE ?", "%"+*filter.Title+"%")
	}
	if filter.ActivityGroupId != nil {
		query = repo.db.Where("activity_group_id = ?", *filter.ActivityGroupId)
	}
	if filter.Priority != nil {
		query = repo.db.Where("priority = ?", *filter.Priority)
	}
	if filter.IsActive != nil {
		query = repo.db.Where("is_active = ?", *filter.IsActive)
	}

	if err = query.Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}
