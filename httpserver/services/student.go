package services

import (
	"database/sql"
	"fmt"
	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/controllers/views"
	"sesi6-gin/httpserver/repositories"
	"sesi6-gin/httpserver/repositories/models"
	"strings"
)

type StudentSvc struct {
	repo repositories.StudentRepo
}

func NewStudentSvc(repo repositories.StudentRepo) *StudentSvc {
	return &StudentSvc{
		repo: repo,
	}
}

func (s *StudentSvc) GetAllStudents() *views.Response {
	students, err := s.repo.GetStudents()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessFindAllResponse(parseModelToStudentGetAll(students), "GET_ALL_STUDENTS")
}

func (s *StudentSvc) CreateStudent(req *params.StudentCreateRequest) *views.Response {
	student := parseRequestToModel(req)
	students, _ := s.repo.GetStudents()
	id := "S-1"
	if students != nil {
		if len(*students) > 0 {
			id = fmt.Sprintf("S-%d", len(*students)+1)
		}
	}
	student.ID = &id
	err := s.repo.CreateStudent(student)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflict(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessCreateResponse(student, "CREATE_STUDENT")
}

func parseRequestToModel(req *params.StudentCreateRequest) *models.Student {
	return &models.Student{
		Name:  req.Name,
		Age:   req.Age,
		Grade: req.Grade,
	}
}

func parseModelToStudentGetAll(mod *[]models.Student) *[]views.StudentGetAll {
	var s []views.StudentGetAll
	for _, st := range *mod {
		s = append(s, views.StudentGetAll{
			ID:   *st.ID,
			Name: st.Name,
		})
	}
	return &s
}
