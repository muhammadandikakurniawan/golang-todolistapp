package entity

import (
	"time"
)

type TodoPriority string

const (
	TodoPriority_VERY_HIGH = "very-high"
	TodoPriority_HIGH      = "high"
	TodoPriority_MEDIUM    = "medium"
	TodoPriority_LOW       = "low"
	TodoPriority_VERY_LOW  = "very-low"
)

var (
	ValidPriority = map[TodoPriority]bool{
		TodoPriority_VERY_HIGH: true,
		TodoPriority_HIGH:      true,
		TodoPriority_MEDIUM:    true,
		TodoPriority_LOW:       true,
		TodoPriority_VERY_LOW:  true,
	}
)

func NewTodo(activityGroupId int64, title string, priority TodoPriority, isActive bool) (newActivity Todo, validationErrorMessages []string) {

	if !ValidPriority[priority] {
		validationErrorMessages = append(validationErrorMessages, "invalid priority")
	}

	createdAt := time.Now()

	newActivity = Todo{
		ActivityGroupId: activityGroupId,
		Title:           title,
		Priority:        priority,
		IsActive:        isActive,
		CreatedAt:       createdAt,
		UpdatedAt:       createdAt,
	}

	return
}

type Todo struct {
	TodoId          int64        `json:"todo_id" gorm:"primaryKey"`
	ActivityGroupId int64        `json:"activity_group_id" gorm:"column:activity_group_id"`
	Title           string       `json:"title" gorm:"column:title"`
	Priority        TodoPriority `json:"priority" gorm:"column:priority"`
	IsActive        bool         `json:"is_active" gorm:"column:is_active"`
	CreatedAt       time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time    `json:"updated_at" gorm:"column:updated_at"`
}

func (*Todo) TableName() string {
	return "todos"
}

type TodoPartialUpdateModel struct {
	TodoId          int64         `json:"todo_id"`
	ActivityGroupId *int64        `json:"activity_group_id"`
	Title           *string       `json:"title"`
	Priority        *TodoPriority `json:"priority"`
	IsActive        *bool         `json:"is_active"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
