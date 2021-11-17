package services

import (
	"github.com/medivh13/mahasiswaservice/internal/repository"
)

type service struct {
	mysqlrepo repository.Repository
}

func NewService(mysqlrepo repository.Repository) Services {
	return &service{mysqlrepo}
}
