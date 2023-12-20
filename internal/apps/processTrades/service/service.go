package service

import "github.com/Arjun-P17/tax-go/internal/repository"

type Service struct {
	dbConnector *repository.Connector
}

func NewService(dbConnector *repository.Connector) *Service {
	return &Service{
		dbConnector: dbConnector,
	}
}
