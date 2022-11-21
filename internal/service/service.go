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

func (s *Service) CreateSource(
	ctx context.Context,
	r *foodv1.CreateSourceRequest,
) (*foodv1.CreateSourceResponse, error) {
	// TODO: Validation
	return s.dataStore.CreateSource(ctx, r)
}

func (s *Service) CreateTypology(
	ctx context.Context,
	r *foodv1.CreateTypologyRequest,
) (*foodv1.CreateTypologyResponse, error) {
	// TODO: Validation
	return s.dataStore.CreateTypology(ctx, r)
}

func (s *Service) ListFoodItems(
	ctx context.Context,
	r *foodv1.ListFoodItemsRequest,
) (*foodv1.ListFoodItemsResponse, error) {
	// TODO: Validation
	return s.dataStore.ListFoodItems(ctx, r)
}
