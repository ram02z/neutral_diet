package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
)

func (c *ConnectWrapper) CreateLifeCycle(
	ctx context.Context,
	req *connect.Request[foodv1.CreateLifeCycleRequest],
) (*connect.Response[foodv1.CreateLifeCycleResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.CreateLifeCycle(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) CreateFoodItem(
	ctx context.Context,
	req *connect.Request[foodv1.CreateFoodItemRequest],
) (*connect.Response[foodv1.CreateFoodItemResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.CreateFoodItem(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) CreateSource(
	ctx context.Context,
	req *connect.Request[foodv1.CreateSourceRequest],
) (*connect.Response[foodv1.CreateSourceResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.CreateSource(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) CreateRegion(
	ctx context.Context,
	req *connect.Request[foodv1.CreateRegionRequest],
) (*connect.Response[foodv1.CreateRegionResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.CreateRegion(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) CreateTypology(
	ctx context.Context,
	req *connect.Request[foodv1.CreateTypologyRequest],
) (*connect.Response[foodv1.CreateTypologyResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.CreateTypology(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) CreateSubTypology(
	ctx context.Context,
	req *connect.Request[foodv1.CreateSubTypologyRequest],
) (*connect.Response[foodv1.CreateSubTypologyResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.CreateSubTypology(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) ListRegions(
	ctx context.Context,
	req *connect.Request[foodv1.ListRegionsRequest],
) (*connect.Response[foodv1.ListRegionsResponse], error) {
	res, err := c.s.ListRegions(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) ListAggregateFoodItems(
	ctx context.Context,
	req *connect.Request[foodv1.ListAggregateFoodItemsRequest],
) (*connect.Response[foodv1.ListAggregateFoodItemsResponse], error) {
	res, err := c.s.ListAggregateFoodItems(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}