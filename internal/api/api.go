package api

import (
	"context"
	"errors"

	"github.com/Arjun-P17/tax-go/internal/models"
	"github.com/Arjun-P17/tax-go/proto/go/stockpb"
)

type ServiceInterface interface {
	GetStockPositions(ctx context.Context) ([]models.StockPosition, error)
}

type ApiInterface interface {
	GetStockPositions(ctx context.Context, req *stockpb.StockRequest) (*stockpb.StockPositions, error)
}

// ForwardCompatibleApiLayer ensures forward compatiblity by providing empty implementations of new methods.
type ForwardCompatibleApiLayer struct {
	stockpb.UnimplementedStockServiceServer
}

type Api struct {
	ForwardCompatibleApiLayer
	ApiInterface
	service ServiceInterface
}

func NewApi(service ServiceInterface) (Api, error) {
	if service == nil {
		return Api{}, errors.New("service is nil")
	}

	return Api{
		service: service,
	}, nil
}
