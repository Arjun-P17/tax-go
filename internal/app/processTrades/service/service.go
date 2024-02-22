package service

import "github.com/Arjun-P17/tax-go/internal/repository"

// TODO: move into internal/service
type Service struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}
