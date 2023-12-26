package api

import (
	"context"

	"github.com/Arjun-P17/tax-go/proto/go/stockpb"
)

func (api Api) GetStockPositions(ctx context.Context, req *stockpb.StockRequest) (*stockpb.StockPositions, error) {
	stockPositions, err := api.service.GetStockPositions(ctx)
	if err != nil {
		return nil, err
	}

	var sp []*stockpb.StockPosition
	for _, position := range stockPositions {
		p := position.ToProtoModel()
		sp = append(sp, &p)
	}

	return &stockpb.StockPositions{
		Positions: sp,
	}, nil
}
