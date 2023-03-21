package entity

import (
	"net/mail"
	"strings"
	"time"
)

func NewActivity(title, email string) (newActivity Activity, validationErrorMessages []string) {

	titleIsEmpty := strings.ReplaceAll(title, " ", "") == ""
	if titleIsEmpty {
		validationErrorMessages = append(validationErrorMessages, "title cannot be empty")
	}

	email = strings.TrimSpace(email)
	_, mailErr := mail.ParseAddress(email)
	if mailErr != nil {
		validationErrorMessages = append(validationErrorMessages, "invalid email")
	}

	createdAt := time.Now()

	newActivity = Activity{
		Title:     title,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	return
}

type Activity struct {
	ActivityId int64     `json:"activity_id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"column:title"`
	Email      string    `json:"email" gorm:"column:email"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (*Activity) TableName() string {
	return "activities"
}

type ActivityPartialUpdateModel struct {
	ActivityId int64     `json:"activity_id"`
	Title      *string   `json:"title"`
	Email      *string   `json:"email"`
	UpdatedAt  time.Time `json:"updated_at"`
}
