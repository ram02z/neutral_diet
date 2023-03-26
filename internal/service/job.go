package service

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	jobv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/job/v1"
)

func (c *ConnectWrapper) SendGoalNotifications(
	ctx context.Context,
	req *connect.Request[jobv1.SendGoalNotificationsRequest],
) (*connect.Response[jobv1.SendGoalNotificationsResponse], error) {
	jobCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	c.s.SendGoalNotifications(jobCtx, c.m)

	return connect.NewResponse(&jobv1.SendGoalNotificationsResponse{}), nil
}

func (c *ConnectWrapper) MarkCompletedGoals(
	ctx context.Context,
	req *connect.Request[jobv1.MarkCompletedGoalsRequest],
) (*connect.Response[jobv1.MarkCompletedGoalsResponse], error) {
	jobCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	c.s.MarkCompletedGoals(jobCtx)

	return connect.NewResponse(&jobv1.MarkCompletedGoalsResponse{}), nil
}
