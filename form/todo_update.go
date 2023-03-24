package form

import "gorm.io/gorm"

type TodoUpdate struct {
	gorm.Model
	ActivityGroupId uint `json:"activity_group_id"`
	Title string `json:"title"`
	IsActive bool `json:"is_active"`
	Priority string `json:"priority"`
}