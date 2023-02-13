package db

import (
	"context"
	"database/sql"

	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateTypology(
	ctx context.Context,
	r *foodv1.CreateTypologyRequest,
) (*foodv1.CreateTypologyResponse, error) {
	queries := db.New(s.dbPool)

	subTypologyID := sql.NullInt32{}
	if r.Typology.SubTypologyId != nil {
		subTypologyID.Valid = true
		subTypologyID.Int32 = *r.Typology.SubTypologyId
	}

	source := db.CreateTypologyParams{
		Name:          r.Typology.Name,
		SubTypologyID: subTypologyID,
	}

	typologyID, err := queries.CreateTypology(ctx, source)
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &foodv1.CreateTypologyResponse{Id: typologyID}, nil
}
