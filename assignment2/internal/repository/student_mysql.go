package repository

import (
	"github.com/damndelion/SDUassignments/assignment2/internal/entity"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetStudents() (students []*entity.Student, err error) {
	err = r.db.Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (r *Repo) GetStudentByID(id string) (student *entity.Student, err error) {
	err = r.db.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (r *Repo) CreateStudent(student *entity.Student) (int, error) {
	err := r.db.Create(&student).Error
	if err != nil {
		return 0, err
	}

	return student.ID, nil
}

func (r *Repo) UpdateStudent(student *entity.Student) (int, error) {
	err := r.db.Model(&student).Where("id = ?", student.ID).Updates(student).Error
	if err != nil {
		return 0, err
	}

	return student.ID, nil
}

func (r *Repo) DeleteStudent(id string) error {
	err := r.db.Delete(&entity.Student{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetStudentByDepID(id string) (students []*entity.Student, err error) {
	err = r.db.Find(&students).Where("department_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (r *Repo) GetStudentEnrollments(id string) (enrollments []*entity.Enrollment, err error) {
	err = r.db.Find(&enrollments).Where("student_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return enrollments, nil
}

func (r *Repo) CountStudentsInDepartment(id string) (count int64, err error) {
	err = r.db.Model(&entity.Student{}).Where("department_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
