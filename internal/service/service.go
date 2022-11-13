package service

import (
	"context"

	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
	"github.com/ram02z/neutral_diet/internal/service/db"
)

type Service struct {
	dataStore *db.Store
}

func NewService(dataStore *db.Store) *Service {
	return &Service{
		dataStore: dataStore,
	}
}

func (s *Service) CreateFoodItem(
	ctx context.Context,
	r *foodv1.CreateFoodItemRequest,
) (*foodv1.CreateFoodItemResponse, error) {
	// TODO: Validation
	return s.dataStore.CreateFoodItem(ctx, r)
}
