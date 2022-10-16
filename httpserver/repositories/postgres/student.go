package postgres

import (
	"database/sql"
	"sesi6-gin/httpserver/repositories"
	"sesi6-gin/httpserver/repositories/models"
)

type studentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) repositories.StudentRepo {
	return &studentRepo{
		db: db,
	}
}

func (s *studentRepo) GetStudents() (*[]models.Student, error) {
	query := `
		SELECT 
			id, name, age, grade
		FROM students
		ORDER BY id ASC
	`

	// proses query data
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	// matikan koneksi
	defer rows.Close()

	var students []models.Student

	// proses scan data
	for rows.Next() {
		student := models.Student{}
		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
			&student.Grade,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}

// singleton version
// func GetStudents() (*[]models.Student, error) {
// 	db := config.GetDB()
// 	query := `
// 		SELECT
// 			id, name, age, grade
// 		FROM students
// 		ORDER BY id ASC
// 	`

// 	// proses query data
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// matikan koneksi
// 	defer rows.Close()

// 	var students []models.Student

// 	// proses scan data
// 	for rows.Next() {
// 		student := models.Student{}
// 		err := rows.Scan(
// 			&student.ID,
// 			&student.Name,
// 			&student.Age,
// 			&student.Grade,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		students = append(students, student)
// 	}

// 	return &students, nil
// }

func (s *studentRepo) CreateStudent(student *models.Student) error {
	query := `
		INSERT INTO students (id,name,age,grade)
		VALUES ($1, $2, $3, $4)
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.ID, student.Name, student.Age, student.Grade)

	return err
}
