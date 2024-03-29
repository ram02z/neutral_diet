package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateSource(
	ctx context.Context,
	r *foodv1.CreateSourceRequest,
) (*foodv1.CreateSourceResponse, error) {
	queries := db.New(s.dbPool)

	source := db.CreateSourceParams{
		Reference: r.Source.Reference,
		Year:      r.Source.Year,
		Region:    int32(r.Source.Region.Number()),
	}

	sourceID, err := queries.CreateSource(ctx, source)
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &foodv1.CreateSourceResponse{Id: sourceID}, nil
}
