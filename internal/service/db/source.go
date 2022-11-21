package db

import (
	"context"
	"database/sql"

	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (s *Store) CreateSource(
	ctx context.Context,
	r *foodv1.CreateSourceRequest,
) (*foodv1.CreateSourceResponse, error) {
	queries := db.New(s.dbPool)

	source := &db.CreateSourceParams{
		Reference: r.GetReference(),
		Year: sql.NullInt16{
			Int16: int16(r.GetYear()),
			Valid: true,
		},
		Location: sql.NullString{
			String: r.GetLocation(),
			Valid:  true,
		},
	}

	sourceID, err := queries.CreateSource(ctx, *source)
	if err != nil {
		return nil, err
	}

	return &foodv1.CreateSourceResponse{Id: sourceID}, nil
}
