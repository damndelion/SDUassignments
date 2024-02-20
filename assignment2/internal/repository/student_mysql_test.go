package repository

import (
	"github.com/damndelion/SDUassignments/assignment2/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var dsn = "root:password@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func TestRepo_CountStudentsInDepartment(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int64
		wantErr   bool
	}{
		{
			name: "Test for success",
			fields: fields{
				db: db,
			},
			args: args{
				id: "1",
			},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "Test for not existing student",
			fields: fields{
				db: db,
			},
			args: args{
				id: "4",
			},
			wantCount: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db: tt.fields.db,
			}
			gotCount, err := r.CountStudentsInDepartment(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountStudentsInDepartment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("CountStudentsInDepartment() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestRepo_CreateStudent(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		student *entity.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				db: db,
			},
			args: args{
				student: &entity.Student{
					Name:         "test",
					Age:          20,
					Course:       "3",
					DepartmentID: "1",
					Gender:       "Non-binary",
				},
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db: tt.fields.db,
			}
			got, err := r.CreateStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
