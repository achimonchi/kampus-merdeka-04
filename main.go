package main

import (
	"fmt"
	"os"
	"sesi6-gin/config"
	"sesi6-gin/httpserver"
	"sesi6-gin/httpserver/controllers"
	"sesi6-gin/httpserver/repositories/gorm"
	"sesi6-gin/httpserver/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}
	studentRepo := gorm.NewStudentRepo(db)
	studentSvc := services.NewStudentSvc(studentRepo)
	studentHandler := controllers.NewStudentController(studentSvc)

	router := gin.Default()

	app := httpserver.NewRouter(router, studentHandler)
	app.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))

	// app := httpserver.CreateRouter()
	// app.Run(":4000")
}
