package models

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	ActivityGroupId uint
	Title string
	IsActive bool
	Priority string
}