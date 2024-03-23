package service

import (
	"github.com/Arjun-P17/tax-go/internal/repository"
)

type Service struct {
	repository repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) *Service {
	return &Service{
		repository: repo,
	}
}
