package response

import (
	"time"
)

type Result struct {
	ID uint `gorm:"column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Email string `gorm:"column:email" json:"email"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}