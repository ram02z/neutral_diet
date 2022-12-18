package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateUser(
	ctx context.Context,
	r *foodv1.CreateUserRequest,
) (*foodv1.CreateUserResponse, error) {
	queries := db.New(s.dbPool)

	userID, err := queries.CreateUser(ctx, r.FirebaseUid)
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &foodv1.CreateUserResponse{Id: userID}, nil
}
