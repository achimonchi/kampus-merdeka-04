package gorm

import (
	"database/sql"
	"sesi6-gin/httpserver/repositories"
	"sesi6-gin/httpserver/repositories/models"

	"github.com/jinzhu/gorm"
)

type studentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) repositories.StudentRepo {
	return &studentRepo{
		db: db,
	}
}

func (s *studentRepo) GetStudents() (*[]models.Student, error) {
	var students []models.Student
	err := s.db.Find(&students).Error
	if err != nil {

		return nil, err
	}

	if len(students) == 0 {
		return nil, sql.ErrNoRows
	}

	return &students, nil
}

func (s *studentRepo) CreateStudent(student *models.Student) error {
	err := s.db.Create(student).Error
	return err
}
