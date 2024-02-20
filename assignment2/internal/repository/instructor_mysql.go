package repository

import (
	"github.com/damndelion/SDUassignments/assignment2/internal/entity"
)

func (r *Repo) GetInstructors() (instructors []*entity.Instructor, err error) {
	err = r.db.Find(&instructors).Error
	if err != nil {
		return nil, err
	}

	return instructors, nil
}

func (r *Repo) GetInstructorByID(id string) (instructor *entity.Instructor, err error) {
	err = r.db.Where("id = ?", id).First(&instructor).Error
	if err != nil {
		return nil, err
	}

	return instructor, nil
}

func (r *Repo) CreateInstructor(instructor *entity.Instructor) (int, error) {
	err := r.db.Create(&instructor).Error
	if err != nil {
		return 0, err
	}

	return instructor.ID, nil
}

func (r *Repo) UpdateInstructor(instructor *entity.Instructor) (int, error) {
	err := r.db.Model(&instructor).Where("id = ?", instructor.ID).Updates(instructor).Error
	if err != nil {
		return 0, err
	}

	return instructor.ID, nil
}

func (r *Repo) DeleteInstructor(id string) error {
	err := r.db.Delete(&entity.Instructor{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetCoursesByInstructor(id string) (courses []*entity.Course, err error) {
	err = r.db.Find(&courses).Where("instructor_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}
