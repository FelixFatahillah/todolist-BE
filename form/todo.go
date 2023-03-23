package form

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ActivityGroupId string `json:"activity_group_id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Priority string `json:"priority" binding:"required"`
}