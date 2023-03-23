package form

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Email string `json:"email" binding:"required"`
}