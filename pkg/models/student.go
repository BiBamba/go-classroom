package models

import (
	"database/sql"

	"github.com/BiBamba/go-classroom/config"
)

var db *sql.DB

type Student struct {
	ID        int64
	Firstname string
	Lastname  string
}

func Init() {
	config.Connect()
}

func addStudent(stu Student) (int64, error) {
	result, err := db.Exec("INSERT INTO ")
}
