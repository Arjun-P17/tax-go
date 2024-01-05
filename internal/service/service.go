package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/internal/models"
	"github.com/Arjun-P17/tax-go/internal/repository"
)

type serviceInterface interface {
	GetStockPositions(ctx context.Context) ([]models.StockPosition, error)
}

type repositoryInterface interface {
	GetAllStockPositions(ctx context.Context) ([]repository.StockPosition, error)
}

type Service struct {
	serviceInterface
	// repository is an interface so service layer is decoupled from the repository layer.
	repository repositoryInterface
}

func NewService(repo repositoryInterface) (Service, error) {
	if repo == nil {
		return Service{}, nil
	}

	return Service{
		repository: repo,
	}, nil
}
