package response

import (
	"time"
)

type TodoResult struct {
	ID uint `gorm:"column:id" json:"id"`
	ActivityGroupId uint `gorm:"column:activity_group_id" json:"activity_group_id"`
	Title string `gorm:"column:title" json:"title"`
	IsActive bool `gorm:"column:is_active" json:"is_active"`
	Priority string `gorm:"column:priority" json:"priority"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}