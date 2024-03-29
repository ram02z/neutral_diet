// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: neutral_diet/job/v1/api.proto

package jobv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/job/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// JobServiceName is the fully-qualified name of the JobService service.
	JobServiceName = "neutral_diet.job.v1.JobService"
)

// JobServiceClient is a client for the neutral_diet.job.v1.JobService service.
type JobServiceClient interface {
	SendGoalNotifications(context.Context, *connect_go.Request[v1.SendGoalNotificationsRequest]) (*connect_go.Response[v1.SendGoalNotificationsResponse], error)
	MarkCompletedGoals(context.Context, *connect_go.Request[v1.MarkCompletedGoalsRequest]) (*connect_go.Response[v1.MarkCompletedGoalsResponse], error)
	SendStreakNotifications(context.Context, *connect_go.Request[v1.SendStreakNotificationsRequest]) (*connect_go.Response[v1.SendStreakNotificationsResponse], error)
	RemoveStaleRegistrationTokens(context.Context, *connect_go.Request[v1.RemoveStaleRegistrationTokensRequest]) (*connect_go.Response[v1.RemoveStaleRegistrationTokensResponse], error)
}

// NewJobServiceClient constructs a client for the neutral_diet.job.v1.JobService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewJobServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) JobServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &jobServiceClient{
		sendGoalNotifications: connect_go.NewClient[v1.SendGoalNotificationsRequest, v1.SendGoalNotificationsResponse](
			httpClient,
			baseURL+"/neutral_diet.job.v1.JobService/SendGoalNotifications",
			opts...,
		),
		markCompletedGoals: connect_go.NewClient[v1.MarkCompletedGoalsRequest, v1.MarkCompletedGoalsResponse](
			httpClient,
			baseURL+"/neutral_diet.job.v1.JobService/MarkCompletedGoals",
			opts...,
		),
		sendStreakNotifications: connect_go.NewClient[v1.SendStreakNotificationsRequest, v1.SendStreakNotificationsResponse](
			httpClient,
			baseURL+"/neutral_diet.job.v1.JobService/SendStreakNotifications",
			opts...,
		),
		removeStaleRegistrationTokens: connect_go.NewClient[v1.RemoveStaleRegistrationTokensRequest, v1.RemoveStaleRegistrationTokensResponse](
			httpClient,
			baseURL+"/neutral_diet.job.v1.JobService/RemoveStaleRegistrationTokens",
			opts...,
		),
	}
}

// jobServiceClient implements JobServiceClient.
type jobServiceClient struct {
	sendGoalNotifications         *connect_go.Client[v1.SendGoalNotificationsRequest, v1.SendGoalNotificationsResponse]
	markCompletedGoals            *connect_go.Client[v1.MarkCompletedGoalsRequest, v1.MarkCompletedGoalsResponse]
	sendStreakNotifications       *connect_go.Client[v1.SendStreakNotificationsRequest, v1.SendStreakNotificationsResponse]
	removeStaleRegistrationTokens *connect_go.Client[v1.RemoveStaleRegistrationTokensRequest, v1.RemoveStaleRegistrationTokensResponse]
}

// SendGoalNotifications calls neutral_diet.job.v1.JobService.SendGoalNotifications.
func (c *jobServiceClient) SendGoalNotifications(ctx context.Context, req *connect_go.Request[v1.SendGoalNotificationsRequest]) (*connect_go.Response[v1.SendGoalNotificationsResponse], error) {
	return c.sendGoalNotifications.CallUnary(ctx, req)
}

// MarkCompletedGoals calls neutral_diet.job.v1.JobService.MarkCompletedGoals.
func (c *jobServiceClient) MarkCompletedGoals(ctx context.Context, req *connect_go.Request[v1.MarkCompletedGoalsRequest]) (*connect_go.Response[v1.MarkCompletedGoalsResponse], error) {
	return c.markCompletedGoals.CallUnary(ctx, req)
}

// SendStreakNotifications calls neutral_diet.job.v1.JobService.SendStreakNotifications.
func (c *jobServiceClient) SendStreakNotifications(ctx context.Context, req *connect_go.Request[v1.SendStreakNotificationsRequest]) (*connect_go.Response[v1.SendStreakNotificationsResponse], error) {
	return c.sendStreakNotifications.CallUnary(ctx, req)
}

// RemoveStaleRegistrationTokens calls neutral_diet.job.v1.JobService.RemoveStaleRegistrationTokens.
func (c *jobServiceClient) RemoveStaleRegistrationTokens(ctx context.Context, req *connect_go.Request[v1.RemoveStaleRegistrationTokensRequest]) (*connect_go.Response[v1.RemoveStaleRegistrationTokensResponse], error) {
	return c.removeStaleRegistrationTokens.CallUnary(ctx, req)
}

// JobServiceHandler is an implementation of the neutral_diet.job.v1.JobService service.
type JobServiceHandler interface {
	SendGoalNotifications(context.Context, *connect_go.Request[v1.SendGoalNotificationsRequest]) (*connect_go.Response[v1.SendGoalNotificationsResponse], error)
	MarkCompletedGoals(context.Context, *connect_go.Request[v1.MarkCompletedGoalsRequest]) (*connect_go.Response[v1.MarkCompletedGoalsResponse], error)
	SendStreakNotifications(context.Context, *connect_go.Request[v1.SendStreakNotificationsRequest]) (*connect_go.Response[v1.SendStreakNotificationsResponse], error)
	RemoveStaleRegistrationTokens(context.Context, *connect_go.Request[v1.RemoveStaleRegistrationTokensRequest]) (*connect_go.Response[v1.RemoveStaleRegistrationTokensResponse], error)
}

// NewJobServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewJobServiceHandler(svc JobServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/neutral_diet.job.v1.JobService/SendGoalNotifications", connect_go.NewUnaryHandler(
		"/neutral_diet.job.v1.JobService/SendGoalNotifications",
		svc.SendGoalNotifications,
		opts...,
	))
	mux.Handle("/neutral_diet.job.v1.JobService/MarkCompletedGoals", connect_go.NewUnaryHandler(
		"/neutral_diet.job.v1.JobService/MarkCompletedGoals",
		svc.MarkCompletedGoals,
		opts...,
	))
	mux.Handle("/neutral_diet.job.v1.JobService/SendStreakNotifications", connect_go.NewUnaryHandler(
		"/neutral_diet.job.v1.JobService/SendStreakNotifications",
		svc.SendStreakNotifications,
		opts...,
	))
	mux.Handle("/neutral_diet.job.v1.JobService/RemoveStaleRegistrationTokens", connect_go.NewUnaryHandler(
		"/neutral_diet.job.v1.JobService/RemoveStaleRegistrationTokens",
		svc.RemoveStaleRegistrationTokens,
		opts...,
	))
	return "/neutral_diet.job.v1.JobService/", mux
}

// UnimplementedJobServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedJobServiceHandler struct{}

func (UnimplementedJobServiceHandler) SendGoalNotifications(context.Context, *connect_go.Request[v1.SendGoalNotificationsRequest]) (*connect_go.Response[v1.SendGoalNotificationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.job.v1.JobService.SendGoalNotifications is not implemented"))
}

func (UnimplementedJobServiceHandler) MarkCompletedGoals(context.Context, *connect_go.Request[v1.MarkCompletedGoalsRequest]) (*connect_go.Response[v1.MarkCompletedGoalsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.job.v1.JobService.MarkCompletedGoals is not implemented"))
}

func (UnimplementedJobServiceHandler) SendStreakNotifications(context.Context, *connect_go.Request[v1.SendStreakNotificationsRequest]) (*connect_go.Response[v1.SendStreakNotificationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.job.v1.JobService.SendStreakNotifications is not implemented"))
}

func (UnimplementedJobServiceHandler) RemoveStaleRegistrationTokens(context.Context, *connect_go.Request[v1.RemoveStaleRegistrationTokensRequest]) (*connect_go.Response[v1.RemoveStaleRegistrationTokensResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.job.v1.JobService.RemoveStaleRegistrationTokens is not implemented"))
}
