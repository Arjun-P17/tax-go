package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/internal/models"
)

func (s Service) GetStockPositions(ctx context.Context) ([]models.StockPosition, error) {
	stockPositions, err := s.DBConnector.GetAllStockPositions(ctx)
	if err != nil {
		return nil, err
	}

	var sp []models.StockPosition
	for _, position := range stockPositions {
		sp = append(sp, position.ToModel())
	}

	return sp, nil
}
