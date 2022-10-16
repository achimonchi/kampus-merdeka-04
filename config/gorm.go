package config

import (
	"fmt"
	"sesi6-gin/httpserver/repositories/models"

	"github.com/jinzhu/gorm"
)

func ConnectPostgresGORM() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(models.User{}, models.Student{})

	return db, nil
}
