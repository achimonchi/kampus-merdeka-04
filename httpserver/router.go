package httpserver

import (
	"sesi6-gin/httpserver/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	student *controllers.StudentController
}

func NewRouter(router *gin.Engine, student *controllers.StudentController) *Router {
	return &Router{
		router:  router,
		student: student,
	}
}

func (r *Router) Start(port string) {
	// step :(1) request masuk, request keluar
	r.router.GET("/ping", controllers.HealthCheck)
	r.router.POST("/users", controllers.CreateUser)

	r.router.GET("/students", r.student.GetAllStudents)
	r.router.POST("/students", r.student.CreateStudent)
	r.router.Run(port)
}

// func (r *Router) CreateRouter() *gin.Engine {

// 	// step :(1) request masuk, request keluar
// 	r.router.GET("/ping", controllers.HealthCheck)
// 	r.router.POST("/users", controllers.CreateUser)

// 	r.router.GET("/students", r.student.GetAllStudents)

// 	return r.router
// }
