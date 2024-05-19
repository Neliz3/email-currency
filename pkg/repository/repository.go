package repository

import "github.com/jmoiron/sqlx"

type Email interface {
}

type Repository struct {
	Email
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
