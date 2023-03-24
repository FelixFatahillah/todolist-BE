package form

import "gorm.io/gorm"

type ActivityUpdate struct {
	gorm.Model
	Title string `json:"title"`
	Email string `json:"email"`
}