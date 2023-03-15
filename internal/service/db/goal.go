package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	userv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
	"github.com/shopspring/decimal"
)

func (s *Store) AddCarbonFootprintGoal(
	ctx context.Context,
	r *userv1.AddCarbonFootprintGoalRequest,
	firebaseUID string,
) (*userv1.AddCarbonFootprintGoalResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	goalID, err := queries.AddCarbonFootprintGoal(ctx, db.AddCarbonFootprintGoalParams{
		UserID:                user.ID,
		Description:           r.CarbonFootprintGoal.Description,
		StartDate:             mapToDate(r.CarbonFootprintGoal.StartDate),
		EndDate:               mapToDate(r.CarbonFootprintGoal.EndDate),
		StartCarbonFootprint:  decimal.NewFromFloat(r.CarbonFootprintGoal.StartCarbonFootprint),
		TargetCarbonFootprint: decimal.NewFromFloat(r.CarbonFootprintGoal.TargetCarbonFootprint),
	})
	if err != nil {
		return nil, err
	}

	return &userv1.AddCarbonFootprintGoalResponse{Id: goalID}, nil
}

func (s *Store) UpdateCarbonFootprintGoal(
	ctx context.Context,
	r *userv1.UpdateCarbonFootprintGoalRequest,
	firebaseUID string,
) (*userv1.UpdateCarbonFootprintGoalResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	err = queries.UpdateCarbonFootprintGoal(ctx, db.UpdateCarbonFootprintGoalParams{
		UserID:    user.ID,
		ID:        r.Id,
		Completed: r.Completed,
	})
	if err != nil {
		return nil, err
	}

	return &userv1.UpdateCarbonFootprintGoalResponse{}, nil
}

func (s *Store) DeleteCarbonFootprintGoal(
	ctx context.Context,
	r *userv1.DeleteCarbonFootprintGoalRequest,
	firebaseUID string,
) (*userv1.DeleteCarbonFootprintGoalResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	err = queries.DeleteCarbonFootprintGoal(ctx, db.DeleteCarbonFootprintGoalParams{
		UserID: user.ID,
		ID:     r.Id,
	})
	if err != nil {
		return nil, err
	}

	return &userv1.DeleteCarbonFootprintGoalResponse{}, nil
}

func (s *Store) GetCarbonFootprintGoals(
	ctx context.Context,
	r *userv1.GetCarbonFootprintGoalsRequest,
	firebaseUID string,
) (*userv1.GetCarbonFootprintGoalsResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	goals, err := queries.GetCarbonFootprintGoals(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	var completedGoals []*userv1.CarbonFootprintGoalResponse
	var activeGoals []*userv1.CarbonFootprintGoalResponse
	for _, g := range goals {
		res := &userv1.CarbonFootprintGoalResponse{
			Id:          g.ID,
			Description: g.Description,
			StartDate: &userv1.Date{
				Year:  int32(g.StartDate.Time.Year()),
				Month: int32(g.StartDate.Time.Month()),
				Day:   int32(g.StartDate.Time.Day()),
			},
			EndDate: &userv1.Date{
				Year:  int32(g.EndDate.Time.Year()),
				Month: int32(g.EndDate.Time.Month()),
				Day:   int32(g.EndDate.Time.Day()),
			},
			StartCarbonFootprint:  g.StartCarbonFootprint.InexactFloat64(),
			TargetCarbonFootprint: g.TargetCarbonFootprint.InexactFloat64(),
		}
		if g.Completed {
			completedGoals = append(completedGoals, res)
		} else {
			activeGoals = append(activeGoals, res)
		}
	}

	return &userv1.GetCarbonFootprintGoalsResponse{
		Completed: completedGoals,
		Active:    activeGoals,
	}, nil
}
