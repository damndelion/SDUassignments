package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID           int
	Name         string `gorm:"unique;not null"`
	Description  string `gorm:"not null"`
	DepartmentID string `gorm:"not null"`
	Department   Department
	InstructorID int64 `gorm:"not null"`
	Instructor   Instructor
	Enrolled     int `gorm:"default:0"`
}
