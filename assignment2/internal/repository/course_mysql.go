package repository

import (
	"github.com/damndelion/SDUassignments/assignment2/internal/entity"
)

func (r *Repo) GetCourses() (courses []*entity.Course, err error) {
	err = r.db.Find(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *Repo) GetCourseByID(id string) (course *entity.Course, err error) {
	err = r.db.Where("id = ?", id).First(&course).Error
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (r *Repo) CreateCourse(course *entity.Course) (int, error) {
	err := r.db.Create(&course).Error
	if err != nil {
		return 0, err
	}

	return course.ID, nil
}

func (r *Repo) UpdateCourse(course *entity.Course) (int, error) {
	err := r.db.Model(&course).Where("id = ?", course.ID).Updates(course).Error
	if err != nil {
		return 0, err
	}

	return course.ID, nil
}

func (r *Repo) DeleteCourse(id string) error {
	err := r.db.Delete(&entity.Course{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetEnrolled() (results map[string]int64, err error) {
	err = r.db.Model(&entity.Course{}).Select("name, enrolled").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
