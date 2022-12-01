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
		Name:        r.FoodItem.GetName(),
		TypologyID:  r.FoodItem.GetTypologyId(),
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

// func (s *Store) ListFoodItems(
// 	ctx context.Context,
// 	r *foodv1.ListFoodItemsRequest,
// ) (*foodv1.ListFoodItemsResponse, error) {
// 	queries := db.New(s.dbPool)
// 	foodItemRows, err := queries.ListAggregateFoodItems(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	foodItems, err := mapToFoodItems(foodItemRows)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &foodv1.ListFoodItemsResponse{FoodItems: foodItems}, nil
// }

// func mapToFoodItems(foodItemRows []db.ListAggregateFoodItemsRow) ([]*foodv1.FoodItem, error) {
// 	foodItems := make([]*foodv1.FoodItem, len(foodItemRows))
// 	for i := range foodItemRows {
// 		var emissions string
// 		err := foodItemRows[i].MedianCarbonFootprint.AssignTo(&emissions)
// 		if err != nil {
// 			return nil, err
// 		}
// 		foodItems[i] = &foodv1.FoodItem{
// 			Id:        foodItemRows[i].N,
// 			Name:      foodItemRows[i].Name,
// 			Emissions: emissions,
// 		}
// 	}

// 	return foodItems, nil
// }
