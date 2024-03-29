package repository

import "database/sql"

type Repo struct {
	db *sql.DB
}

type RepoInterface interface {
	CakeRepo
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}
