package dao

import (
	"42tokyo-road-to-dojo-go/pkg/database"
	"database/sql"
)

type Dao interface {
	NewUser() User
}

type dao struct {
	db *sql.DB
}

func NewDao() (Dao, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}
	return &dao{db: db}, nil
}

func (d *dao) NewUser() User {
	return &user{db: d.db}
}
