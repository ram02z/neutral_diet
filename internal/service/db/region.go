package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateRegion(
	ctx context.Context,
	r *foodv1.CreateRegionRequest,
) (*foodv1.CreateRegionResponse, error) {
	queries := db.New(s.dbPool)

	err := queries.CreateRegion(ctx, r.GetRegion().GetName())
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &foodv1.CreateRegionResponse{}, nil
}
