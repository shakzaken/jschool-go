package dbrepo

import (
	"database/sql"
	"jschool/repository"
)


type PostgressDbRepo struct {
	DB *sql.DB
}

func NewPostgresDbRepo(db *sql.DB) repository.DatabaseRepo {
	return &PostgressDbRepo{
		DB: db,
	}
}