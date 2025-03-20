package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/internal/models"
)

func (s Service) GetStockPositions(ctx context.Context) ([]models.PortfolioPosition, error) {
	stockPositions, err := s.repository.GetAllStockPositions(ctx)
	if err != nil {
		return nil, err
	}

	var sp []models.PortfolioPosition
	for _, position := range stockPositions {
		sp = append(sp, position.ToModel())
	}

	return sp, nil
}
