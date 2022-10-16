package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	host           = os.Getenv("DB_HOST")
	port           = os.Getenv("DB_PORT")
	user           = os.Getenv("DB_USER")
	password       = os.Getenv("DB_PASS")
	dbname         = os.Getenv("DB_NAME")
	dbMaxIdle      = 4
	dbMaxOpenConns = 25
)

// cara singleton
// var (
// 	DB *sql.DB
// )

// func ConnectPostgres() error {
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}

// 	DB = db

// 	return nil
// }

// func GetDB() *sql.DB {
// 	return DB
// }

func ConnectPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxOpenConns)
	db.SetMaxIdleConns(dbMaxIdle)

	return db, nil
}
