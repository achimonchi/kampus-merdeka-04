package repositories

import "sesi6-gin/httpserver/repositories/models"

type StudentRepo interface {
	GetStudents() (*[]models.Student, error)
	CreateStudent(student *models.Student) error
}
