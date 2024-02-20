package repository

import (
	"github.com/damndelion/SDUassignments/assignment2/internal/entity"
)

func (r *Repo) GetDepartments() (departments []*entity.Department, err error) {
	err = r.db.Find(&departments).Error
	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *Repo) GetDepartmentByID(id string) (department *entity.Department, err error) {
	err = r.db.Where("id = ?", id).First(&department).Error
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (r *Repo) CreateDepartment(department *entity.Department) (int, error) {
	err := r.db.Create(&department).Error
	if err != nil {
		return 0, err
	}

	return department.ID, nil
}

func (r *Repo) UpdateDepartment(department *entity.Department) (int, error) {
	err := r.db.Model(&department).Where("id = ?", department.ID).Updates(department).Error
	if err != nil {
		return 0, err
	}

	return department.ID, nil
}

func (r *Repo) DeleteDepartment(id string) error {
	err := r.db.Delete(&entity.Department{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
