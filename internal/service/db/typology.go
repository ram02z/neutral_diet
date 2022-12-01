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

	subTypologyID := &sql.NullInt32{}
	err := subTypologyID.Scan(r.Typology.SubTypologyId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	source := db.CreateTypologyParams{
		Name:          r.Typology.GetName(),
		Description:   r.Typology.GetDescription(),
		SubTypologyID: *subTypologyID,
	}

	typologyID, err := queries.CreateTypology(ctx, source)
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}

	return &foodv1.CreateTypologyResponse{Id: typologyID}, nil
}
