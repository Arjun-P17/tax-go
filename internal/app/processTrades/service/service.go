package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/internal/repository"
)

type serviceInterface interface {
	ProcessTrades(ctx context.Context) error
}

// TODO: move into internal/service
type Service struct {
	serviceInterface
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}
