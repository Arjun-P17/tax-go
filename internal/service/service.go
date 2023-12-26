package service

import "github.com/Arjun-P17/tax-go/internal/repository"

// TODO: should we define service interface here?
// TODO: should the db connector be an interface?

type Service struct {
	DBConnector *repository.Connector
}

func NewService(dbConnector repository.Connector) Service {
	return Service{
		DBConnector: &dbConnector,
	}
}
