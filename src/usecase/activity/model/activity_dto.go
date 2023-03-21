package model

import (
	"time"

	"github.com/muhammadandikakurniawan/golang-todolistapp/src/entity"
)

func SetupActivityDtoFromEntity(e entity.Activity) (dto ActivityDto) {
	dto = ActivityDto{
		Id:        e.ActivityId,
		Title:     &e.Title,
		Email:     &e.Email,
		CreatedAt: &e.CreatedAt,
		UpdatedAt: &e.UpdatedAt,
	}
	return
}

// swagger:model
type ActivityDto struct {
	Id        int64      `json:"id,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Email     *string    `json:"email,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (data ActivityDto) TransformToNewEntity() (newActivity entity.Activity, validationErrorMessages []string) {
	title, email := "", ""

	if data.Title != nil {
		title = *data.Title
	}
	if data.Email != nil {
		email = *data.Email
	}
	return entity.NewActivity(title, email)
}

func (data ActivityDto) TransformToPartialUpdate() (activityUpdate entity.ActivityPartialUpdateModel) {
	activityUpdate.ActivityId = data.Id
	activityUpdate.Title = data.Title
	activityUpdate.Email = data.Email
	return
}
