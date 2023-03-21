package model

import (
	"time"

	"github.com/muhammadandikakurniawan/golang-todolistapp/src/entity"
)

func SetupTodoDtoFromEntity(e entity.Todo) (dto TodoDto) {
	dto = TodoDto{
		Id:              e.TodoId,
		Priority:        &e.Priority,
		ActivityGroupId: &e.ActivityGroupId,
		Title:           &e.Title,
		IsActive:        &e.IsActive,
		CreatedAt:       &e.CreatedAt,
		UpdatedAt:       &e.UpdatedAt,
	}
	return
}

// swagger:model
type TodoDto struct {
	Id              int64                `json:"id,omitempty"`
	ActivityGroupId *int64               `json:"activity_group_id,omitempty"`
	Title           *string              `json:"title,omitempty"`
	Priority        *entity.TodoPriority `json:"priority,omitempty"`
	IsActive        *bool                `json:"is_active,omitempty"`
	CreatedAt       *time.Time           `json:"created_at,omitempty"`
	UpdatedAt       *time.Time           `json:"updated_at,omitempty"`
}

func (data TodoDto) TransformToNewEntity() (newActivity entity.Todo, validationErrorMessages []string) {
	activityGroupId, title, priority, isActive := int64(0), "", entity.TodoPriority(""), false

	if data.ActivityGroupId != nil {
		activityGroupId = *data.ActivityGroupId
	}
	if data.Title != nil {
		title = *data.Title
	}
	if data.Priority != nil {
		priority = *data.Priority
	}
	if data.IsActive != nil {
		isActive = *data.IsActive
	}

	return entity.NewTodo(activityGroupId, title, priority, isActive)
}

func (data TodoDto) TransformToPartialUpdate() (activityUpdate entity.TodoPartialUpdateModel, validationErrMsg []string) {

	if data.Priority != nil {
		if !entity.ValidPriority[*data.Priority] {
			validationErrMsg = append(validationErrMsg, "invalid priority")
		}
	}

	invalidData := len(validationErrMsg) > 0
	if invalidData {
		return
	}

	activityUpdate.TodoId = data.Id
	activityUpdate.ActivityGroupId = data.ActivityGroupId
	activityUpdate.Title = data.Title
	activityUpdate.IsActive = data.IsActive
	activityUpdate.Priority = data.Priority
	return
}
