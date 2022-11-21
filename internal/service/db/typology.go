package db

import (
	"context"

	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateTypology(
	ctx context.Context,
	r *foodv1.CreateTypologyRequest,
) (*foodv1.CreateTypologyResponse, error) {
	queries := db.New(s.dbPool)

	source := &db.CreateTypologyParams{
		Name:          r.GetName(),
		SubTypologies: r.GetSubTypologies(),
	}

	typologyID, err := queries.CreateTypology(ctx, *source)
	if err != nil {
		return nil, err
	}

	return &foodv1.CreateTypologyResponse{Id: typologyID}, nil
}
