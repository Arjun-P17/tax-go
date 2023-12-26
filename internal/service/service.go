package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/internal/models"
	"github.com/Arjun-P17/tax-go/internal/repository"
)

type databaseInterface interface {
	GetAllStockPositions(ctx context.Context) ([]repository.StockPosition, error)
}

type serviceInterface interface {
	GetStockPositions(ctx context.Context) ([]models.StockPosition, error)
}

type Service struct {
	serviceInterface
	// database is an interface so service layer is decoupled from the database layer.
	database databaseInterface
}

func NewService(db databaseInterface) (Service, error) {
	if db == nil {
		return Service{}, nil
	}

	return Service{
		database: db,
	}, nil
}
