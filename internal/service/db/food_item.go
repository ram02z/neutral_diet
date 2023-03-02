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

func (s *Store) ListAggregateFoodItems(
	ctx context.Context,
	r *foodv1.ListAggregateFoodItemsRequest,
) (*foodv1.ListAggregateFoodItemsResponse, error) {
	queries := db.New(s.dbPool)

	var foodItems []*foodv1.AggregateFoodItem
	if r.Region == foodv1.Region_REGION_UNSPECIFIED {
		foodItemRows, err := queries.ListAggregateFoodItems(ctx)
		if err != nil {
			return nil, err
		}

		foodItems = make([]*foodv1.AggregateFoodItem, len(foodItemRows))
		for i := range foodItemRows {
			foodItems[i] = &foodv1.AggregateFoodItem{
				Id:                    foodItemRows[i].ID,
				FoodName:              foodItemRows[i].FoodName,
				MedianCarbonFootprint: foodItemRows[i].MedianCarbonFootprint.InexactFloat64(),
				Region:                foodv1.Region_REGION_UNSPECIFIED,
			}
		}
	} else {
		foodItemRows, err := queries.ListAggregateFoodItemsByRegion(ctx, int32(r.Region))
		if err != nil {
			return nil, err
		}

		foodItems = make([]*foodv1.AggregateFoodItem, len(foodItemRows))
		for i := range foodItemRows {
			foodItems[i] = &foodv1.AggregateFoodItem{
				Id:                    foodItemRows[i].ID,
				FoodName:              foodItemRows[i].FoodName,
				MedianCarbonFootprint: foodItemRows[i].MedianCarbonFootprint.InexactFloat64(),
				Region:                r.Region,
			}
		}
	}
	return &foodv1.ListAggregateFoodItemsResponse{FoodItems: foodItems}, nil
}

type FoodItemInfo struct {
	TypologyName    string
	SubTypologyName string
	N               int64
}

func (s *Store) GetFoodItemInfo(
	ctx context.Context,
	r *foodv1.GetFoodItemInfoRequest,
) (*foodv1.GetFoodItemInfoResponse, error) {
	queries := db.New(s.dbPool)
	var foodItemInfo FoodItemInfo
	var sources []*foodv1.Source
	if r.Region == foodv1.Region_REGION_UNSPECIFIED {
		foodItemInfoRow, err := queries.GetFoodItemInfo(ctx, r.Id)
		if err != nil {
			return nil, err
		}

		foodItemInfo = FoodItemInfo{
			TypologyName:    foodItemInfoRow.TypologyName,
			SubTypologyName: foodItemInfoRow.SubTypologyName.String,
			N:               foodItemInfoRow.N,
		}

		sourceRows, err := queries.ListSourcesByFoodItem(ctx, r.Id)
		if err != nil {
			return nil, err
		}

		sources = make([]*foodv1.Source, len(sourceRows))
		for i := range sourceRows {
			sources[i] = &foodv1.Source{
				Reference: sourceRows[i].Reference,
				Year:      sourceRows[i].Year,
				Region:    foodv1.Region(sourceRows[i].Region),
			}
		}
	} else {
		foodItemInfoRow, err := queries.GetFoodItemInfoByRegion(ctx, db.GetFoodItemInfoByRegionParams{
			FoodItemID: r.Id,
			Region:     int32(r.Region),
		})
		if err != nil {
			return nil, err
		}

		foodItemInfo = FoodItemInfo{
			TypologyName:    foodItemInfoRow.TypologyName,
			SubTypologyName: foodItemInfoRow.SubTypologyName.String,
			N:               foodItemInfoRow.N,
		}

		sourceRows, err := queries.ListSourcesByFoodItemAndRegion(
			ctx,
			db.ListSourcesByFoodItemAndRegionParams{
				FoodItemID: r.Id,
				Region:     int32(r.Region),
			},
		)
		if err != nil {
			return nil, err
		}

		sources = make([]*foodv1.Source, len(sourceRows))
		for i := range sourceRows {
			sources[i] = &foodv1.Source{
				Reference: sourceRows[i].Reference,
				Year:      sourceRows[i].Year,
				Region:    r.Region,
			}
		}

	}
	return &foodv1.GetFoodItemInfoResponse{
		FoodItemInfo: &foodv1.FoodItemInfo{
			TypologyName:     foodItemInfo.TypologyName,
			SubTypologyName:  foodItemInfo.SubTypologyName,
			NonUniqueSources: foodItemInfo.N,
			Sources:          sources,
		},
	}, nil
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
