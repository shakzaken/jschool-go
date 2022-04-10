package handlers

import (
	"database/sql"
	"jschool/repository"
	"jschool/repository/dbrepo"
)


type AppRepository struct {
	DB repository.DatabaseRepo
}


func NewAppRepo(conn *sql.DB) *AppRepository {
	return &AppRepository{
		DB : dbrepo.NewPostgresDbRepo(conn),
	}
}