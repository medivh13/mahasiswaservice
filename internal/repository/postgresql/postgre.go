package repository

import (
	"github.com/medivh13/mahasiswaservice/internal/repository"

	"github.com/jmoiron/sqlx"
)

const ()

var statement PreparedStatement

type PreparedStatement struct {
}

type MySQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &MySQLRepo{Conn}
	// InitPreparedStatement(repo)
	return repo
}
