package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
	"github.com/ram02z/neutral_diet/internal/service/db"
)

type ConnectWrapper struct {
	s *db.Store
}

func NewConnectWrapper(s *db.Store) *ConnectWrapper {
	return &ConnectWrapper{s: s}
}

func (c *ConnectWrapper) CreateLifeCycle(
	ctx context.Context,
	req *connect.Request[foodv1.CreateLifeCycleRequest],
) (*connect.Response[foodv1.CreateLifeCycleResponse], error) {
	res, err := c.s.CreateLifeCycle(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}

func (c *ConnectWrapper) CreateFoodItem(
	ctx context.Context,
	req *connect.Request[foodv1.CreateFoodItemRequest],
) (*connect.Response[foodv1.CreateFoodItemResponse], error) {
	res, err := c.s.CreateFoodItem(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}

func (c *ConnectWrapper) CreateSource(
	ctx context.Context,
	req *connect.Request[foodv1.CreateSourceRequest],
) (*connect.Response[foodv1.CreateSourceResponse], error) {
	res, err := c.s.CreateSource(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}

func (c *ConnectWrapper) CreateRegion(
	ctx context.Context,
	req *connect.Request[foodv1.CreateRegionRequest],
) (*connect.Response[foodv1.CreateRegionResponse], error) {
	res, err := c.s.CreateRegion(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}

func (c *ConnectWrapper) CreateTypology(
	ctx context.Context,
	req *connect.Request[foodv1.CreateTypologyRequest],
) (*connect.Response[foodv1.CreateTypologyResponse], error) {
	res, err := c.s.CreateTypology(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}

func (c *ConnectWrapper) CreateSubTypology(
	ctx context.Context,
	req *connect.Request[foodv1.CreateSubTypologyRequest],
) (*connect.Response[foodv1.CreateSubTypologyResponse], error) {
	res, err := c.s.CreateSubTypology(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}
