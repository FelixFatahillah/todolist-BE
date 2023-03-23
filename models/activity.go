package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Title string
	Email string
}