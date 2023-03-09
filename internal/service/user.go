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

func (c *ConnectWrapper) UpdateFoodItem(
	ctx context.Context,
	req *connect.Request[userv1.UpdateFoodItemRequest],
) (*connect.Response[userv1.UpdateFoodItemResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	err = validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.UpdateFoodItemFromLog(ctx, req.Msg, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) DeleteFoodItem(
	ctx context.Context,
	req *connect.Request[userv1.DeleteFoodItemRequest],
) (*connect.Response[userv1.DeleteFoodItemResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	err = validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.DeleteFoodItemFromLog(ctx, req.Msg, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) UpdateUserSettings(
	ctx context.Context,
	req *connect.Request[userv1.UpdateUserSettingsRequest],
) (*connect.Response[userv1.UpdateUserSettingsResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	err = validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.UpdateUserSettings(ctx, req.Msg, token.UID)
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

	err = validate(out.Msg)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConnectWrapper) GetFoodItemLog(
	ctx context.Context,
	req *connect.Request[userv1.GetFoodItemLogRequest],
) (*connect.Response[userv1.GetFoodItemLogResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	res, err := c.s.GetFoodItemLog(ctx, req.Msg, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) GetUserInsights(
	ctx context.Context,
	req *connect.Request[userv1.GetUserInsightsRequest],
) (*connect.Response[userv1.GetUserInsightsResponse], error) {
	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	res, err := c.s.GetUserInsights(ctx, req.Msg, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) GetFoodItemLogDays(

	ctx context.Context,
	req *connect.Request[userv1.GetFoodItemLogDaysRequest],
) (*connect.Response[userv1.GetFoodItemLogDaysResponse], error) {

	token, err := c.verify(ctx, req.Header())
	if err != nil {
		return nil, err
	}

	err = validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.GetFoodItemLogDays(ctx, req.Msg, token.UID)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}
