package form

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ActivityGroupId uint `json:"activity_group_id" binding:"required"`
	Title string `json:"title" binding:"required"`
	IsActive bool `json:"is_active"`
	Priority string `json:"priority" binding:"required"`
}