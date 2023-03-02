package db

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
	userv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
	"github.com/shopspring/decimal"
)

func (s *Store) AddFoodItemToLog(
	ctx context.Context,
	r *userv1.AddFoodItemRequest,
	firebaseUID string,
) (*userv1.AddFoodItemResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	var medianCarbonFootprint decimal.Decimal
	if r.FoodLogItem.Region == foodv1.Region_REGION_UNSPECIFIED {
		aggFoodItem, err := queries.GetAggregateFoodItem(ctx, r.FoodLogItem.FoodItemId)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		medianCarbonFootprint = aggFoodItem.MedianCarbonFootprint
	} else {
		regionalAggFoodItem, err := queries.GetRegionalAggregateFoodItem(
			ctx,
			db.GetRegionalAggregateFoodItemParams{
				FoodItemID: r.FoodLogItem.FoodItemId,
				Region:     int32(r.FoodLogItem.Region),
			},
		)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		medianCarbonFootprint = regionalAggFoodItem.MedianCarbonFootprint
	}

	weight := decimal.NewFromFloat(r.FoodLogItem.Weight)

	carbonFootprint := calculateCarbonFootprintByWeight(
		medianCarbonFootprint,
		weight,
		r.FoodLogItem.WeightUnit,
	)

	foodItemLogID, err := queries.AddFoodItemToLog(ctx, db.AddFoodItemToLogParams{
		FoodItemID: r.FoodLogItem.FoodItemId,
		Weight:     decimal.NewFromFloat(r.FoodLogItem.Weight),
		WeightUnit: int32(r.FoodLogItem.WeightUnit),
		UserID:     user.ID,
		LogDate:    mapToDate(r.FoodLogItem.GetDate()),
		Region:     int32(r.FoodLogItem.Region),
	})
	if err != nil {
		return nil, err
	}

	return &userv1.AddFoodItemResponse{
		Id:              foodItemLogID,
		CarbonFootprint: carbonFootprint.InexactFloat64(),
	}, nil
}

func (s *Store) UpdateFoodItemFromLog(
	ctx context.Context,
	r *userv1.UpdateFoodItemRequest,
	firebaseUID string,
) (*userv1.UpdateFoodItemResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	foodItemID, err := queries.GetFoodItemIdByLogId(ctx, r.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	var medianCarbonFootprint decimal.Decimal
	if r.Region == foodv1.Region_REGION_UNSPECIFIED {
		aggFoodItem, err := queries.GetAggregateFoodItem(ctx, foodItemID)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		medianCarbonFootprint = aggFoodItem.MedianCarbonFootprint
	} else {
		regionalAggFoodItem, err := queries.GetRegionalAggregateFoodItem(
			ctx,
			db.GetRegionalAggregateFoodItemParams{
				FoodItemID: foodItemID,
				Region:     int32(r.Region),
			},
		)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		medianCarbonFootprint = regionalAggFoodItem.MedianCarbonFootprint
	}

	weight := decimal.NewFromFloat(r.Weight)

	carbonFootprint := calculateCarbonFootprintByWeight(
		medianCarbonFootprint,
		weight,
		r.WeightUnit,
	)

	err = queries.UpdateFoodItemFromLog(ctx, db.UpdateFoodItemFromLogParams{
		UserID:     user.ID,
		ID:         r.Id,
		Weight:     weight,
		WeightUnit: int32(r.WeightUnit),
	})

	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return &userv1.UpdateFoodItemResponse{CarbonFootprint: carbonFootprint.InexactFloat64()}, err
}

func (s *Store) DeleteFoodItemFromLog(
	ctx context.Context,
	r *userv1.DeleteFoodItemRequest,
	firebaseUID string,
) (*userv1.DeleteFoodItemResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	err = queries.DeleteFoodItemFromLog(ctx, db.DeleteFoodItemFromLogParams{
		UserID: user.ID,
		ID:     r.Id,
	})

	if err != nil {
		return nil, err
	}

	return &userv1.DeleteFoodItemResponse{}, nil
}

func (s *Store) GetFoodItemLog(
	ctx context.Context,
	r *userv1.GetFoodItemLogRequest,
	firebaseUID string,
) (*userv1.GetFoodItemLogResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	foodItemLog, err := queries.GetFoodItemLogByDate(ctx, db.GetFoodItemLogByDateParams{
		UserID:  user.ID,
		LogDate: mapToDate(r.GetDate()),
	})
	if err != nil {
		return nil, err
	}

	foodLogItems := make([]*userv1.FoodLogItemResponse, len(foodItemLog))
	for i := range foodItemLog {
		var medianCarbonFootprint decimal.Decimal
		if foodItemLog[i].Region == int32(foodv1.Region_REGION_UNSPECIFIED) {
			aggregateFoodItem, err := queries.GetAggregateFoodItem(ctx, foodItemLog[i].FoodItemID)
			if err != nil {
				return nil, err
			}
			medianCarbonFootprint = aggregateFoodItem.MedianCarbonFootprint
		} else {
			regionalAggregateFoodItem, err := queries.GetRegionalAggregateFoodItem(
				ctx,
				db.GetRegionalAggregateFoodItemParams{
					FoodItemID: foodItemLog[i].FoodItemID,
					Region:     foodItemLog[i].Region,
				},
			)
			if err != nil {
				return nil, err
			}
			medianCarbonFootprint = regionalAggregateFoodItem.MedianCarbonFootprint
		}

		weightUnit := userv1.WeightUnit(foodItemLog[i].WeightUnit)
		carbonFootprint := calculateCarbonFootprintByWeight(
			medianCarbonFootprint,
			foodItemLog[i].Weight,
			weightUnit,
		)
		foodLogItems[i] = &userv1.FoodLogItemResponse{
			Id:              foodItemLog[i].ID,
			FoodItemId:      foodItemLog[i].FoodItemID,
			Name:            foodItemLog[i].Name,
			Weight:          foodItemLog[i].Weight.InexactFloat64(),
			WeightUnit:      weightUnit,
			CarbonFootprint: carbonFootprint.InexactFloat64(),
			Date: &userv1.Date{
				Year:  int32(foodItemLog[i].LogDate.Time.Year()),
				Month: int32(foodItemLog[i].LogDate.Time.Month()),
				Day:   int32(foodItemLog[i].LogDate.Time.Day()),
			},
			Region: foodv1.Region(foodItemLog[i].Region),
		}
	}

	return &userv1.GetFoodItemLogResponse{
		FoodItemLog: foodLogItems,
	}, nil
}

func calculateCarbonFootprintByWeight(
	carbonFootprint decimal.Decimal,
	weight decimal.Decimal,
	weightUnit userv1.WeightUnit,
) decimal.Decimal {
	multiplier := decimal.NewFromFloat(1)
	switch weightUnit {
	case userv1.WeightUnit_WEIGHT_UNIT_KILOGRAM:
		multiplier = decimal.NewFromFloat(1)
	case userv1.WeightUnit_WEIGHT_UNIT_GRAM:
		multiplier = decimal.NewFromFloat(0.001)
	case userv1.WeightUnit_WEIGHT_UNIT_OUNCE:
		multiplier = decimal.NewFromFloat(0.02834952)
	case userv1.WeightUnit_WEIGHT_UNIT_POUND:
		multiplier = decimal.NewFromFloat(0.45359237)
	}

	weight = weight.Mul(multiplier)

	return carbonFootprint.Mul(weight)
}

func mapToDate(date *userv1.Date) pgtype.Date {
	if date == nil || date.GetYear() == 0 || date.GetMonth() == 0 || date.GetDay() == 0 {
		return pgtype.Date{Valid: false}
	}

	time := time.Date(
		int(date.GetYear()),
		time.Month(int(date.GetMonth())),
		int(date.GetDay()),
		0,
		0,

		0,
		0,
		time.UTC)
	return pgtype.Date{Time: time, Valid: true}
}
