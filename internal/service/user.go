package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	userv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
	"github.com/ram02z/neutral_diet/internal"
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
	uid := req.Header().Get(config.UserIDHeaderKey)
	res, err := c.s.DeleteUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	// Delete user from firebase user table
	err = c.a.DeleteUser(ctx, uid)
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
	uid := req.Header().Get(config.UserIDHeaderKey)

	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.AddFoodItemToLog(ctx, req.Msg, uid)
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
	uid := req.Header().Get(config.UserIDHeaderKey)

	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.UpdateFoodItemFromLog(ctx, req.Msg, uid)
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
	uid := req.Header().Get(config.UserIDHeaderKey)

	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.DeleteFoodItemFromLog(ctx, req.Msg, uid)
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
	uid := req.Header().Get(config.UserIDHeaderKey)

	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.UpdateUserSettings(ctx, req.Msg, uid)
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
	uid := req.Header().Get(config.UserIDHeaderKey)

	res, err := c.s.GetUser(ctx, uid)
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
	uid := req.Header().Get(config.UserIDHeaderKey)

	res, err := c.s.GetFoodItemLog(ctx, req.Msg, uid)
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
	uid := req.Header().Get(config.UserIDHeaderKey)

	res, err := c.s.GetUserInsights(ctx, req.Msg, uid)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) GetUserProgress(
	ctx context.Context,
	req *connect.Request[userv1.GetUserProgressRequest],
) (*connect.Response[userv1.GetUserProgressResponse], error) {
	uid := req.Header().Get(config.UserIDHeaderKey)

	res, err := c.s.GetUserProgress(ctx, req.Msg, uid)
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
	uid := req.Header().Get(config.UserIDHeaderKey)

	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.GetFoodItemLogDays(ctx, req.Msg, uid)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}
