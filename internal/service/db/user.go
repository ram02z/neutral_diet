package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	userv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
)

func (s *Store) CreateUser(
	ctx context.Context,
	r *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	queries := db.New(s.dbPool)

	userID, err := queries.CreateUser(ctx, r.FirebaseUid)
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
