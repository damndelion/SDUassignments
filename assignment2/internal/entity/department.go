package entity

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	ID          int
	Code        string `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}
