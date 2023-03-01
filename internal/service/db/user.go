package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
	userv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
	"github.com/shopspring/decimal"
)

func (s *Store) CreateUser(
	ctx context.Context,
	r *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	queries := db.New(s.dbPool)

	userID, err := queries.CreateUser(ctx, db.CreateUserParams{
		FirebaseUid:        r.FirebaseUid,
		Region:             int32(foodv1.Region_REGION_UNSPECIFIED.Number()),
		CfLimit:            decimal.NewFromFloat(0.1),
		DietaryRequirement: int32(userv1.UserSettings_DIETARY_REQUIREMENT_UNSPECIFIED.Number()),
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &userv1.CreateUserResponse{Id: userID}, nil
}

func (s *Store) DeleteUser(
	ctx context.Context,
	firebaseUID string,
) (*userv1.DeleteUserResponse, error) {
	queries := db.New(s.dbPool)

	userID, err := queries.DeleteUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	err = queries.DeleteUserLog(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &userv1.DeleteUserResponse{}, nil
}

func (s *Store) GetUser(
	ctx context.Context,
	firebaseUID string,
) (*userv1.GetUserSettingsResponse, error) {
	queries := db.New(s.dbPool)

	user, err := queries.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	userResponse := userv1.GetUserSettingsResponse{
		UserSettings: &userv1.UserSettings{
			Region:             foodv1.Region(user.Region),
			CfLimit:            user.CfLimit.InexactFloat64(),
			DietaryRequirement: userv1.UserSettings_DietaryRequirement(user.DietaryRequirement),
		},
	}

	return &userResponse, nil
}

func (s *Store) UpdateUserSettings(
	ctx context.Context,
	r *userv1.UpdateUserSettingsRequest,
	firebaseUID string,
) (*userv1.UpdateUserSettingsResponse, error) {
	queries := db.New(s.dbPool)

	err := queries.UpdateUserSettings(ctx, db.UpdateUserSettingsParams{
		FirebaseUid:        firebaseUID,
		Region:             int32(r.GetUserSettings().GetRegion().Number()),
		CfLimit:            decimal.NewFromFloat(r.GetUserSettings().CfLimit),
		DietaryRequirement: int32(r.GetUserSettings().GetDietaryRequirement().Number()),
	})
	if err != nil {
		return nil, err
	}

	return &userv1.UpdateUserSettingsResponse{}, nil
}
