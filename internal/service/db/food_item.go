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
	queries := db.New(s.dbPool)
	foodItemRows, err := queries.ListFoodItems(ctx)
	if err != nil {
		return nil, err
	}

	foodItems, err := mapToFoodItems(foodItemRows)
	if err != nil {
		return nil, err
	}

	return &foodv1.ListFoodItemsResponse{FoodItems: foodItems}, nil
}

func mapToFoodItems(foodItemRows []db.ListFoodItemsRow) ([]*foodv1.FoodItem, error) {
	foodItems := make([]*foodv1.FoodItem, len(foodItemRows))
	for i := range foodItemRows {
		var emissions string
		err := foodItemRows[i].CarbonFootprint.AssignTo(&emissions)
		if err != nil {
			return nil, err
		}
		foodItems[i] = &foodv1.FoodItem{
			Id:        foodItemRows[i].ID,
			Name:      foodItemRows[i].Name,
			Emissions: emissions,
		}
	}

	return foodItems, nil
}
