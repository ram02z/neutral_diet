package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateSubTypology(
	ctx context.Context,
	r *foodv1.CreateSubTypologyRequest,
) (*foodv1.CreateSubTypologyResponse, error) {
	queries := db.New(s.dbPool)

	subTypologyID, err := queries.CreateSubTypology(ctx, r.SubTypology.Name)
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &foodv1.CreateSubTypologyResponse{Id: subTypologyID}, nil
}

func (s *Store) ListSubTypologyNames(
	ctx context.Context,
) (*foodv1.ListSubTypologyNamesResponse, error) {
	queries := db.New(s.dbPool)

	subTypologies, err := queries.ListSubTypologies(ctx)
	if err != nil {
		return nil, err
	}

	subTypologyNames := make([]string, len(subTypologies))
	for i, t := range subTypologies {
		subTypologyNames[i] = t.Name
	}

	return &foodv1.ListSubTypologyNamesResponse{Names: subTypologyNames}, nil
}
