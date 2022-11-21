package db

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateFoodItem(
	ctx context.Context,
	r *foodv1.CreateFoodItemRequest,
) (*foodv1.CreateFoodItemResponse, error) {
	queries := db.New(s.dbPool)
	carbonFootprint := new(pgtype.Numeric)
	err := carbonFootprint.Set(r.GetEmissions())
	if err != nil {
		return nil, err
	}
	foodItem := &db.CreateFoodItemParams{
		Name:            r.GetName(),
		CarbonFootprint: *carbonFootprint,
		TypologyID:      r.GetTypologyId(),
		SourceID:        r.GetSourceId(),
	}

	foodItemID, err := queries.CreateFoodItem(ctx, *foodItem)
	if err != nil {
		return nil, err
	}

	return &foodv1.CreateFoodItemResponse{Id: foodItemID}, nil
}

func (s *Store) ListFoodItems(
	ctx context.Context,
	r *foodv1.ListFoodItemsRequest,
) (*foodv1.ListFoodItemsResponse, error) {
	// TODO: implement

	return &foodv1.ListFoodItemsResponse{
		FoodItems: []*foodv1.FoodItem{},
	}, nil
}
