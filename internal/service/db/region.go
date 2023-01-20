package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

const DefaultRegionName string = "World"

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

func (s *Store) ListRegions(
	ctx context.Context,
	r *foodv1.ListRegionsRequest,
) (*foodv1.ListRegionsResponse, error) {
	queries := db.New(s.dbPool)

	regions, err := queries.ListRegions(ctx)
	if err != nil {
		return nil, err
	}

	regionsRes := make([]*foodv1.Region, len(regions))
	for i := range regions {
		regionsRes[i] = &foodv1.Region{Name: regions[i]}
	}

	return &foodv1.ListRegionsResponse{Regions: regionsRes}, nil
}
