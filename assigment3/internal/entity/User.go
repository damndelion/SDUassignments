package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
