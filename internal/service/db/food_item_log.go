package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/jackc/pgtype"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	userv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
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

	carbonFootprint := new(pgtype.Numeric)
	err = carbonFootprint.Set(r.FoodLogItem.GetCarbonFootprint())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	weight := new(pgtype.Numeric)
	err = weight.Set(r.FoodLogItem.GetWeight())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	foodItemLogID, err := queries.AddFoodItemToLog(ctx, db.AddFoodItemToLogParams{
		FoodItemID:      r.FoodLogItem.FoodItemId,
		Weight:          *weight,
		CarbonFootprint: *carbonFootprint,
		UserID:          user.ID,
	})
	if err != nil {
		return nil, err
	}

	return &userv1.AddFoodItemResponse{Id: foodItemLogID}, nil
}
