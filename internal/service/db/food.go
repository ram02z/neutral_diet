package db

import (
	"context"
	"log"

	"github.com/ram02z/neutral_diet/internal/gen/db"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
	"github.com/rs/zerolog"
)

func (s *Store) CreateFoodItem(
	ctx context.Context,
	r *foodv1.CreateFoodItemRequest,
) (*foodv1.CreateFoodItemResponse, error) {

	queries := db.New(s.dbPool)

	foodItem, err := queries.GetFoodItem(ctx, 3)
	if err != nil {
		log.Fatal(err)
	}

	logger := zerolog.Ctx(ctx)
	logger.Info().Msg(foodItem.Name)

	return nil, nil
}
