package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/internal/repository"
)

type serviceInterface interface {
	ProcessTrades(ctx context.Context) error
}

type Service struct {
	serviceInterface
	repository repository.RepositoryInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}
