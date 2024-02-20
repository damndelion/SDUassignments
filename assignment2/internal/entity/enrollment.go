package entity

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
	ID        int
	StudentID int
	Student   Student
	CourseID  int
	Course    Course
	Grade     string
}
