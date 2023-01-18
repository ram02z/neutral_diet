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

type Validator interface {
	Validate() error
}

func NewConnectWrapper(s *db.Store) *ConnectWrapper {
	return &ConnectWrapper{s: s}
}

func validate(r Validator) error {
	err := r.Validate()
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	return nil
}

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
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

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
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

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
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

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
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

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
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

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
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

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
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}

func (c *ConnectWrapper) CreateUser(
	ctx context.Context,
	req *connect.Request[foodv1.CreateUserRequest],
) (*connect.Response[foodv1.CreateUserResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}

func (c *ConnectWrapper) AddFoodItem(
	ctx context.Context,
	req *connect.Request[foodv1.AddFoodItemRequest],
) (*connect.Response[foodv1.AddFoodItemResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.AddFoodItemToLog(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)
	// TODO: export the headers
	out.Header().Set("API-Version", "v1")

	return out, nil
}
