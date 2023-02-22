package db

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateFoodItem(
	ctx context.Context,
	r *foodv1.CreateFoodItemRequest,
) (*foodv1.CreateFoodItemResponse, error) {
	queries := db.New(s.dbPool)

	suggestedCf, err := mapToDBCfTypes(r.FoodItem.GetCfType())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	foodItem := &db.CreateFoodItemParams{
		Name:        r.FoodItem.Name,
		TypologyID:  r.FoodItem.TypologyId,
		SuggestedCf: suggestedCf,
	}

	foodItemID, err := queries.CreateFoodItem(ctx, *foodItem)
	if err != nil {
		return nil, err
	}

	return &foodv1.CreateFoodItemResponse{Id: foodItemID}, nil
}

func mapToDBCfTypes(cfType foodv1.FoodItem_CfType) (db.CfTypes, error) {
	switch cfType {
	case foodv1.FoodItem_CF_TYPE_ITEM:
		return db.CfTypesItem, nil
	case foodv1.FoodItem_CF_TYPE_SUB_TYPOLOGY:
		return db.CfTypesSubTypology, nil
	case foodv1.FoodItem_CF_TYPE_TYPOLOGY:
		return db.CfTypesTypology, nil
	}

	return "", errors.New("could not map CfType")
}

func (s *Store) ListAggregateFoodItems(
	ctx context.Context,
	r *foodv1.ListAggregateFoodItemsRequest,
) (*foodv1.ListAggregateFoodItemsResponse, error) {
	queries := db.New(s.dbPool)
	foodItemRows, err := queries.ListAggregateFoodItems(ctx)
	if err != nil {
		return nil, err
	}

	foodItems, err := mapToFoodItems(foodItemRows)
	if err != nil {
		return nil, err
	}

	return &foodv1.ListAggregateFoodItemsResponse{FoodItems: foodItems}, nil
}

func mapToFoodItems(foodItemRows []db.ListAggregateFoodItemsRow) ([]*foodv1.AggregateFoodItem, error) {
	foodItems := make([]*foodv1.AggregateFoodItem, len(foodItemRows))
	for i := range foodItemRows {
		foodItems[i] = &foodv1.AggregateFoodItem{
			Id:                    foodItemRows[i].ID,
			FoodName:              foodItemRows[i].FoodName,
			MedianCarbonFootprint: foodItemRows[i].MedianCarbonFootprint.InexactFloat64(),
			FoodItemInfo: &foodv1.FoodItemInfo{
				TypologyName:    foodItemRows[i].TypologyName,
				SubTypologyName: foodItemRows[i].SubTypologyName.String,
				NoSources:       foodItemRows[i].N,
			},
		}
	}

	return foodItems, nil
}
