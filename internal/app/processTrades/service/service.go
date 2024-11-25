package service

import (
	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/pkg/date"
)

type Service struct {
	repository  repository.RepositoryInterface
	currencyMap map[date.Date]float64
}

func NewService(repo repository.RepositoryInterface, currencyMap map[date.Date]float64) *Service {
	return &Service{
		repository:  repo,
		currencyMap: currencyMap,
	}
}
