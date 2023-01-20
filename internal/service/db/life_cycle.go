package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
	"github.com/shopspring/decimal"
)

func (s *Store) CreateLifeCycle(
	ctx context.Context,
	r *foodv1.CreateLifeCycleRequest,
) (*foodv1.CreateLifeCycleResponse, error) {
	queries := db.New(s.dbPool)

	lifeCycle := db.CreateLifeCycleParams{
		CarbonFootprint: decimal.NewFromFloat(r.LifeCycle.CarbonFootprint),
		FoodItemID:      r.LifeCycle.GetFoodItemId(),
		SourceID:        r.LifeCycle.GetFoodItemId(),
	}

	lifeCycleID, err := queries.CreateLifeCycle(ctx, lifeCycle)
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &foodv1.CreateLifeCycleResponse{Id: lifeCycleID}, nil
}
