package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	const dsn = "root:ca-tech-dojo@tcp(localhost:3306)/sample?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
