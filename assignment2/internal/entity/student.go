package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID           int
	Name         string `gorm:"not null"`
	Age          int    `gorm:"not null"`
	Course       string `gorm:"not null"`
	DepartmentID string `gorm:"not null"`
	Department   Department
	Gender       string `gorm:"not null"`
}
