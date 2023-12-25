package api

import (
	"context"

	"github.com/Arjun-P17/tax-go/proto/go/stockpb"
)

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
}
