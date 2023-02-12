package db

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateTypology(
	ctx context.Context,
	r *foodv1.CreateTypologyRequest,
) (*foodv1.CreateTypologyResponse, error) {
	queries := db.New(s.dbPool)

	source := db.CreateTypologyParams{
		Name:          r.Typology.GetName(),
		SubTypologyID: pgtype.Int4{
			Int32: r.Typology.GetSubTypologyId(),
			Valid: true,
		},
	}

	typologyID, err := queries.CreateTypology(ctx, source)
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &foodv1.CreateTypologyResponse{Id: typologyID}, nil
}
