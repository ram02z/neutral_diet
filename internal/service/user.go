package service

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	config "github.com/ram02z/neutral_diet/internal"
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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}
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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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

func (c *ConnectWrapper) AddDevice(
	ctx context.Context,
	req *connect.Request[userv1.AddDeviceRequest],
) (*connect.Response[userv1.AddDeviceResponse], error) {
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

	res, err := c.s.AddDevice(ctx, req.Msg, uid)
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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

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

func (c *ConnectWrapper) AddCarbonFootprintGoal(
	ctx context.Context,
	req *connect.Request[userv1.AddCarbonFootprintGoalRequest],
) (*connect.Response[userv1.AddCarbonFootprintGoalResponse], error) {
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.AddCarbonFootprintGoal(ctx, req.Msg, uid)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) UpdateCarbonFootprintGoal(
	ctx context.Context,
	req *connect.Request[userv1.UpdateCarbonFootprintGoalRequest],
) (*connect.Response[userv1.UpdateCarbonFootprintGoalResponse], error) {
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.UpdateCarbonFootprintGoal(ctx, req.Msg, uid)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) DeleteCarbonFootprintGoal(
	ctx context.Context,
	req *connect.Request[userv1.DeleteCarbonFootprintGoalRequest],
) (*connect.Response[userv1.DeleteCarbonFootprintGoalResponse], error) {
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

	err := validate(req.Msg)
	if err != nil {
		return nil, err
	}

	res, err := c.s.DeleteCarbonFootprintGoal(ctx, req.Msg, uid)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}

func (c *ConnectWrapper) GetCarbonFootprintGoals(
	ctx context.Context,
	req *connect.Request[userv1.GetCarbonFootprintGoalsRequest],
) (*connect.Response[userv1.GetCarbonFootprintGoalsResponse], error) {
	uid, ok := ctx.Value(config.UserTokenKey).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeInternal, errors.New("user token is malformed"))
	}

	res, err := c.s.GetCarbonFootprintGoals(ctx, req.Msg, uid)
	if err != nil {
		return nil, err
	}

	out := connect.NewResponse(res)

	return out, nil
}
