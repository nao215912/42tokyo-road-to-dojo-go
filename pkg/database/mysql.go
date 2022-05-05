package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:ca-tech-dojo@tcp(localhost:3306)/sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}
