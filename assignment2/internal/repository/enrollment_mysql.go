package repository

import (
	"github.com/damndelion/SDUassignments/assignment2/internal/entity"
	"gorm.io/gorm"
)

func (r *Repo) GetEnrollments() (enrollments []*entity.Enrollment, err error) {
	err = r.db.Find(&enrollments).Error
	if err != nil {
		return nil, err
	}

	return enrollments, nil
}

func (r *Repo) GetEnrollmentByID(id string) (enrollment *entity.Enrollment, err error) {
	err = r.db.Where("id = ?", id).First(&enrollment).Error
	if err != nil {
		return nil, err
	}

	return enrollment, nil
}

func (r *Repo) CreateEnrollment(enrollment *entity.Enrollment) (int, error) {
	err := r.db.Create(&enrollment).Error
	if err != nil {
		return 0, err
	}

	return enrollment.ID, nil
}

func (r *Repo) UpdateEnrollment(enrollment *entity.Enrollment) (int, error) {
	err := r.db.Model(&enrollment).Where("id = ?", enrollment.ID).Updates(enrollment).Error
	if err != nil {
		return 0, err
	}

	return enrollment.ID, nil
}

func (r *Repo) DeleteEnrollment(id string) error {
	err := r.db.Delete(&entity.Enrollment{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) EnrollStudent(enrollment *entity.Enrollment) error {
	tx := r.db.Begin()

	err := tx.Create(&enrollment).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&entity.Course{}).Where("id = ?", enrollment.CourseID).
		Update("enrolled", gorm.Expr("enrolled + 1")).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *Repo) AverageGradeByCourseId(id string) (grade float64, err error) {
	err = r.db.Model(&entity.Enrollment{}).Select("avg(grade)").
		Where("course_id = ?", id).Scan(&grade).Error

	if err != nil {
		return 0, err
	}

	return grade, nil
}
