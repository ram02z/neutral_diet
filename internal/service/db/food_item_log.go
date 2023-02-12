package db

import (
	"context"
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

	foodItemLogID, err := queries.AddFoodItemToLog(ctx, db.AddFoodItemToLogParams{
		FoodItemID:      r.FoodLogItem.FoodItemId,
		Weight:          decimal.NewFromFloat(r.FoodLogItem.Weight),
		CarbonFootprint: decimal.NewFromFloat(r.FoodLogItem.CarbonFootprint),
		UserID:          user.ID,
		LogDate:         mapToDate(r.FoodLogItem.GetDate()),
	})
	if err != nil {
		return nil, err
	}

	return &userv1.AddFoodItemResponse{Id: foodItemLogID}, nil
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

func mapToFoodLogItems(foodItemLogRows []db.FoodItemLog) ([]*userv1.FoodLogItem, error) {
	foodLogItems := make([]*userv1.FoodLogItem, len(foodItemLogRows))
	for i := range foodItemLogRows {
		foodLogItems[i] = &userv1.FoodLogItem{
			FoodItemId:      foodItemLogRows[i].ID,
			Weight:          foodItemLogRows[i].Weight.InexactFloat64(),
			CarbonFootprint: foodItemLogRows[i].CarbonFootprint.InexactFloat64(),
			Date: &userv1.Date{
				Year:  int32(foodItemLogRows[i].LogDate.Time.Year()),
				Month: int32(foodItemLogRows[i].LogDate.Time.Month()),
				Day:   int32(foodItemLogRows[i].LogDate.Time.Day()),
			},
		}

		return foodLogItems, nil
	}

	return foodLogItems, nil
}
