package models

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	ActivityGroupId int
	Title string
	Priority string
}