package db

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/jackc/pgx/v5"
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

	quantity := decimal.NewFromFloat(r.FoodLogItem.Quantity)

	carbonFootprint := calculateCarbonFootprint(
		medianCarbonFootprint,
		quantity,
		r.FoodLogItem.Unit,
	)

	foodItemLogID, err := queries.AddFoodItemToLog(ctx, db.AddFoodItemToLogParams{
		FoodItemID:      r.FoodLogItem.FoodItemId,
		Quantity:        decimal.NewFromFloat(r.FoodLogItem.Quantity),
		Unit:            int32(r.FoodLogItem.Unit),
		UserID:          user.ID,
		LogDate:         mapToDate(r.FoodLogItem.GetDate()),
		Region:          int32(r.FoodLogItem.Region),
		CarbonFootprint: carbonFootprint,
		Meal:            int32(r.FoodLogItem.Meal),
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

	quantity := decimal.NewFromFloat(r.Quantity)

	carbonFootprint := calculateCarbonFootprint(
		medianCarbonFootprint,
		quantity,
		r.Unit,
	)

	err = queries.UpdateFoodItemFromLog(ctx, db.UpdateFoodItemFromLogParams{
		UserID:          user.ID,
		ID:              r.Id,
		Quantity:        quantity,
		Unit:            int32(r.Unit),
		CarbonFootprint: carbonFootprint,
		Meal:            int32(r.Meal),
	})

	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return &userv1.UpdateFoodItemResponse{CarbonFootprint: carbonFootprint.InexactFloat64()}, err
}

func (s *Store) GetUserInsights(
	ctx context.Context,
	r *userv1.GetUserInsightsRequest,
	firebaseUID string,
) (*userv1.GetUserInsightsResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	userDailyAverage, err := queries.GetUserDailyAverageCarbonFootprint(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	dailyAverage, err := queries.GetDailyAverageCarbonFootprint(ctx)
	if err != nil {
		dailyAverage = decimal.NewFromFloat(0)
	}

	dailyAverageDietaryRequirement, err := queries.GetDailyAverageCarbonFootprintByDietaryRequirement(
		ctx,
		user.DietaryRequirement,
	)
	if err != nil {
		dailyAverageDietaryRequirement = decimal.NewFromFloat(0)
	}

	streakRow, err := queries.GetFoodItemLogStreak(ctx, user.ID)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	streakLength := int32(streakRow.ConsecutiveDates)

	return &userv1.GetUserInsightsResponse{
		DailyUserAverage: userDailyAverage.InexactFloat64(),
		DailyAverageCarbonFootprintDietaryRequirement: dailyAverageDietaryRequirement.InexactFloat64(),
		DailyAverageCarbonFootprintOverall:            dailyAverage.InexactFloat64(),
		StreakLen:                                     streakLength,
		IsStreakActive:                                streakRow.Active,
	}, nil
}

func (s *Store) GetUserProgress(
	ctx context.Context,
	r *userv1.GetUserProgressRequest,
	firebaseUID string,
) (*userv1.GetUserProgressResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	var startDate pgtype.Date
	var endDate pgtype.Date
	if r.DateRange != nil {
		startDate = mapToDate(r.DateRange.Start)
		endDate = mapToDate(r.DateRange.End)
	} else {
		startDate = pgtype.Date{
			Time:  time.Time{},
			Valid: true,
		}
		endDate = pgtype.Date{
			Time:  time.Now().AddDate(0, 0, 1),
			Valid: true,
		}
	}

	dailySums, err := queries.GetDailyCarbonFootprintByDateRange(
		ctx,
		db.GetDailyCarbonFootprintByDateRangeParams{
			UserID:    user.ID,
			StartDate: startDate,
			EndDate:   endDate,
		},
	)

	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	dailyProgressAll := make(map[string]float64)
	dailyProgressBreakfast := make(map[string]float64)
	dailyProgressLunch := make(map[string]float64)
	dailyProgressDinner := make(map[string]float64)
	dailyProgressSnacks := make(map[string]float64)
	for _, e := range dailySums {
		date := e.LogDate.Time.Format("2006-01-02")
		emissions := e.CarbonFootprint.InexactFloat64()
		switch userv1.Meal(e.Meal) {
		case userv1.Meal_MEAL_BREAKFAST:
			dailyProgressBreakfast[date] = emissions
		case userv1.Meal_MEAL_LUNCH:
			dailyProgressLunch[date] = emissions
		case userv1.Meal_MEAL_DINNER:
			dailyProgressDinner[date] = emissions
		case userv1.Meal_MEAL_UNSPECIFIED:
			dailyProgressSnacks[date] = emissions
		}
		dailyProgressAll[date] += emissions
	}

	return &userv1.GetUserProgressResponse{
		DailyProgressAll:       dailyProgressAll,
		DailyProgressBreakfast: dailyProgressBreakfast,
		DailyProgressLunch:     dailyProgressLunch,
		DailyProgressDinner:    dailyProgressDinner,
		DailyProgressSnacks:    dailyProgressSnacks,
	}, err
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
	for i, f := range foodItemLog {
		foodLogItems[i] = &userv1.FoodLogItemResponse{
			Id:              f.ID,
			FoodItemId:      f.FoodItemID,
			Name:            f.Name,
			Quantity:        f.Quantity.InexactFloat64(),
			Unit:            userv1.Unit(f.Unit),
			CarbonFootprint: f.CarbonFootprint.InexactFloat64(),
			Date: &userv1.Date{
				Year:  int32(f.LogDate.Time.Year()),
				Month: int32(f.LogDate.Time.Month()),
				Day:   int32(f.LogDate.Time.Day()),
			},
			Region: foodv1.Region(f.Region),
			Meal:   userv1.Meal(f.Meal),
		}
	}

	return &userv1.GetFoodItemLogResponse{FoodItemLog: foodLogItems}, nil
}

func (s *Store) GetFoodItemLogDays(
	ctx context.Context,
	r *userv1.GetFoodItemLogDaysRequest,
	firebaseUID string,
) (*userv1.GetFoodItemLogDaysResponse, error) {
	queries := db.New(s.dbPool)
	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	days, err := queries.GetLoggedDaysInMonth(ctx, db.GetLoggedDaysInMonthParams{
		UserID: user.ID,
		Month:  r.Month,
		Year:   r.Year,
	})
	if err != nil {
		return nil, err
	}

	return &userv1.GetFoodItemLogDaysResponse{Days: days}, err
}

func calculateCarbonFootprint(
	carbonFootprint decimal.Decimal,
	quantity decimal.Decimal,
	unit userv1.Unit,
) decimal.Decimal {
	// Default is kilogram or litre
	multiplier := decimal.NewFromFloat(1)
	switch unit {
	case userv1.Unit_UNIT_GRAM:
		multiplier = decimal.NewFromFloat(0.001)
	case userv1.Unit_UNIT_OUNCE:
		multiplier = decimal.NewFromFloat(0.02834952)
	case userv1.Unit_UNIT_POUND:
		multiplier = decimal.NewFromFloat(0.45359237)
	case userv1.Unit_UNIT_MILLILITRE:
		multiplier = decimal.NewFromFloat(0.001)
	case userv1.Unit_UNIT_GALLON:
		multiplier = decimal.NewFromFloat(4.54609)
	case userv1.Unit_UNIT_PINT:
		multiplier = decimal.NewFromFloat(0.568261)
	}

	quantity = quantity.Mul(multiplier)

	return carbonFootprint.Mul(quantity)
}

func mapToDate(date *userv1.Date) pgtype.Date {
	if date == nil || date.GetYear() == 0 || date.GetMonth() == 0 || date.GetDay() == 0 {
		return pgtype.Date{Valid: false}
	}

	dateTime := time.Date(
		int(date.GetYear()),
		time.Month(int(date.GetMonth())),
		int(date.GetDay()),
		0,
		0,

		0,
		0,
		time.UTC)
	return pgtype.Date{Time: dateTime, Valid: true}
}
