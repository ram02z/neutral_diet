package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	userv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
)

func (c *ConnectWrapper) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.CreateUserResponse], error) {
	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) DeleteUser(
	ctx context.Context,
	req *connect.Request[userv1.DeleteUserRequest],
) (*connect.Response[userv1.DeleteUserResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	res, err := c.s.DeleteUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	// Delete user from firebase user table
	err = c.a.DeleteUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) AddFoodItem(
	ctx context.Context,
	req *connect.Request[userv1.AddFoodItemRequest],
) (*connect.Response[userv1.AddFoodItemResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	err = validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.AddFoodItemToLog(ctx, req.Msg, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) UpdateUserRegion(
	ctx context.Context,
	req *connect.Request[userv1.UpdateUserRegionRequest],
) (*connect.Response[userv1.UpdateUserRegionResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	err = validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.UpdateUserRegion(ctx, req.Msg, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) GetUserSettings(
	ctx context.Context,
	req *connect.Request[userv1.GetUserSettingsRequest],
) (*connect.Response[userv1.GetUserSettingsResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	res, err := c.s.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}
