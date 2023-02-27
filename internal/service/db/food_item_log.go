package db

import (
	"context"
	"errors"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ram02z/neutral_diet/internal/gen/db"
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

	aggFoodItem, err := queries.GetAggregateFoodItem(ctx, r.FoodLogItem.FoodItemId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	weight := decimal.NewFromFloat(r.FoodLogItem.Weight)

	carbonFootprint := calculateCarbonFootprintByWeight(
		aggFoodItem.MedianCarbonFootprint,
		weight,
		r.FoodLogItem.WeightUnit,
	)
	weightUnit, err := mapToDBWeightUnit(r.FoodLogItem.WeightUnit)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	foodItemLogID, err := queries.AddFoodItemToLog(ctx, db.AddFoodItemToLogParams{
		FoodItemID: r.FoodLogItem.FoodItemId,
		Weight:     decimal.NewFromFloat(r.FoodLogItem.Weight),
		WeightUnit: weightUnit,
		UserID:     user.ID,
		LogDate:    mapToDate(r.FoodLogItem.GetDate()),
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

	aggFoodItem, err := queries.GetAggregateFoodItem(ctx, foodItemID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	weight := decimal.NewFromFloat(r.Weight)

	carbonFootprint := calculateCarbonFootprintByWeight(
		aggFoodItem.MedianCarbonFootprint,
		weight,
		r.WeightUnit,
	)

	weightUnit, err := mapToDBWeightUnit(r.WeightUnit)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	err = queries.UpdateFoodItemFromLog(ctx, db.UpdateFoodItemFromLogParams{
		UserID:     user.ID,
		ID:         r.Id,
		Weight:     weight,
		WeightUnit: weightUnit,
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

	foodLogItems, err := mapToFoodLogItems(foodItemLog)
	if err != nil {
		return nil, err
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

func mapToDBWeightUnit(weightUnit userv1.WeightUnit) (db.WeightUnit, error) {
	switch weightUnit {
	case userv1.WeightUnit_WEIGHT_UNIT_KILOGRAM:
		return db.WeightUnitKilogram, nil
	case userv1.WeightUnit_WEIGHT_UNIT_GRAM:
		return db.WeightUnitGram, nil
	case userv1.WeightUnit_WEIGHT_UNIT_OUNCE:
		return db.WeightUnitOunce, nil
	case userv1.WeightUnit_WEIGHT_UNIT_POUND:
		return db.WeightUnitPound, nil
	}

	return "", errors.New("couldn't map WeightUnit")
}

func mapFromDBWeightUnit(dbWeighUnit db.WeightUnit) userv1.WeightUnit {
	switch dbWeighUnit {
	case db.WeightUnitKilogram:
		return userv1.WeightUnit_WEIGHT_UNIT_KILOGRAM
	case db.WeightUnitGram:
		return userv1.WeightUnit_WEIGHT_UNIT_GRAM
	case db.WeightUnitOunce:
		return userv1.WeightUnit_WEIGHT_UNIT_OUNCE
	case db.WeightUnitPound:
		return userv1.WeightUnit_WEIGHT_UNIT_POUND
	default:
		return userv1.WeightUnit_WEIGHT_UNIT_KILOGRAM
	}
}

func mapToFoodLogItems(
	foodItemLogRows []db.GetFoodItemLogByDateRow,
) ([]*userv1.FoodLogItemResponse, error) {
	foodLogItems := make([]*userv1.FoodLogItemResponse, len(foodItemLogRows))
	for i := range foodItemLogRows {
		weightUnit := mapFromDBWeightUnit(foodItemLogRows[i].WeightUnit)
		carbonFootprint := calculateCarbonFootprintByWeight(
			foodItemLogRows[i].MedianCarbonFootprint,
			foodItemLogRows[i].Weight,
			weightUnit,
		)
		foodLogItems[i] = &userv1.FoodLogItemResponse{
			Id:              foodItemLogRows[i].ID,
			FoodItemId:      foodItemLogRows[i].FoodItemID,
			Name:            foodItemLogRows[i].Name,
			Weight:          foodItemLogRows[i].Weight.InexactFloat64(),
			WeightUnit:      weightUnit,
			CarbonFootprint: carbonFootprint.InexactFloat64(),
			Date: &userv1.Date{
				Year:  int32(foodItemLogRows[i].LogDate.Time.Year()),
				Month: int32(foodItemLogRows[i].LogDate.Time.Month()),
				Day:   int32(foodItemLogRows[i].LogDate.Time.Day()),
			},
		}
	}

	return foodLogItems, nil
}
