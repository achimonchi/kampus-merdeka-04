package controllers

import (
	"net/http"
	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StudentController struct {
	svc *services.StudentSvc
}

func NewStudentController(svc *services.StudentSvc) *StudentController {
	return &StudentController{
		svc: svc,
	}
}

func (s *StudentController) GetAllStudents(ctx *gin.Context) {
	response := s.svc.GetAllStudents()
	WriteJsonRespnse(ctx, response)
}

func (s *StudentController) CreateStudent(ctx *gin.Context) {
	var req params.StudentCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := s.svc.CreateStudent(&req)
	WriteJsonRespnse(ctx, response)

}
