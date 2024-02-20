package entity

import "gorm.io/gorm"

type Instructor struct {
	gorm.Model
	ID    int
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}
