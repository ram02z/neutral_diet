// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: neutral_diet/user/v1/api.proto

package userv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
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
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "neutral_diet.user.v1.UserService"
)

// UserServiceClient is a client for the neutral_diet.user.v1.UserService service.
type UserServiceClient interface {
	AddFoodItem(context.Context, *connect_go.Request[v1.AddFoodItemRequest]) (*connect_go.Response[v1.AddFoodItemResponse], error)
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error)
	GetUserSettings(context.Context, *connect_go.Request[v1.GetUserSettingsRequest]) (*connect_go.Response[v1.GetUserSettingsResponse], error)
	UpdateUserSettings(context.Context, *connect_go.Request[v1.UpdateUserSettingsRequest]) (*connect_go.Response[v1.UpdateUserSettingsResponse], error)
}

// NewUserServiceClient constructs a client for the neutral_diet.user.v1.UserService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		addFoodItem: connect_go.NewClient[v1.AddFoodItemRequest, v1.AddFoodItemResponse](
			httpClient,
			baseURL+"/neutral_diet.user.v1.UserService/AddFoodItem",
			opts...,
		),
		createUser: connect_go.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+"/neutral_diet.user.v1.UserService/CreateUser",
			opts...,
		),
		deleteUser: connect_go.NewClient[v1.DeleteUserRequest, v1.DeleteUserResponse](
			httpClient,
			baseURL+"/neutral_diet.user.v1.UserService/DeleteUser",
			opts...,
		),
		getUserSettings: connect_go.NewClient[v1.GetUserSettingsRequest, v1.GetUserSettingsResponse](
			httpClient,
			baseURL+"/neutral_diet.user.v1.UserService/GetUserSettings",
			opts...,
		),
		updateUserSettings: connect_go.NewClient[v1.UpdateUserSettingsRequest, v1.UpdateUserSettingsResponse](
			httpClient,
			baseURL+"/neutral_diet.user.v1.UserService/UpdateUserSettings",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	addFoodItem        *connect_go.Client[v1.AddFoodItemRequest, v1.AddFoodItemResponse]
	createUser         *connect_go.Client[v1.CreateUserRequest, v1.CreateUserResponse]
	deleteUser         *connect_go.Client[v1.DeleteUserRequest, v1.DeleteUserResponse]
	getUserSettings    *connect_go.Client[v1.GetUserSettingsRequest, v1.GetUserSettingsResponse]
	updateUserSettings *connect_go.Client[v1.UpdateUserSettingsRequest, v1.UpdateUserSettingsResponse]
}

// AddFoodItem calls neutral_diet.user.v1.UserService.AddFoodItem.
func (c *userServiceClient) AddFoodItem(ctx context.Context, req *connect_go.Request[v1.AddFoodItemRequest]) (*connect_go.Response[v1.AddFoodItemResponse], error) {
	return c.addFoodItem.CallUnary(ctx, req)
}

// CreateUser calls neutral_diet.user.v1.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// DeleteUser calls neutral_diet.user.v1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// GetUserSettings calls neutral_diet.user.v1.UserService.GetUserSettings.
func (c *userServiceClient) GetUserSettings(ctx context.Context, req *connect_go.Request[v1.GetUserSettingsRequest]) (*connect_go.Response[v1.GetUserSettingsResponse], error) {
	return c.getUserSettings.CallUnary(ctx, req)
}

// UpdateUserSettings calls neutral_diet.user.v1.UserService.UpdateUserSettings.
func (c *userServiceClient) UpdateUserSettings(ctx context.Context, req *connect_go.Request[v1.UpdateUserSettingsRequest]) (*connect_go.Response[v1.UpdateUserSettingsResponse], error) {
	return c.updateUserSettings.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the neutral_diet.user.v1.UserService service.
type UserServiceHandler interface {
	AddFoodItem(context.Context, *connect_go.Request[v1.AddFoodItemRequest]) (*connect_go.Response[v1.AddFoodItemResponse], error)
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error)
	GetUserSettings(context.Context, *connect_go.Request[v1.GetUserSettingsRequest]) (*connect_go.Response[v1.GetUserSettingsResponse], error)
	UpdateUserSettings(context.Context, *connect_go.Request[v1.UpdateUserSettingsRequest]) (*connect_go.Response[v1.UpdateUserSettingsResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/neutral_diet.user.v1.UserService/AddFoodItem", connect_go.NewUnaryHandler(
		"/neutral_diet.user.v1.UserService/AddFoodItem",
		svc.AddFoodItem,
		opts...,
	))
	mux.Handle("/neutral_diet.user.v1.UserService/CreateUser", connect_go.NewUnaryHandler(
		"/neutral_diet.user.v1.UserService/CreateUser",
		svc.CreateUser,
		opts...,
	))
	mux.Handle("/neutral_diet.user.v1.UserService/DeleteUser", connect_go.NewUnaryHandler(
		"/neutral_diet.user.v1.UserService/DeleteUser",
		svc.DeleteUser,
		opts...,
	))
	mux.Handle("/neutral_diet.user.v1.UserService/GetUserSettings", connect_go.NewUnaryHandler(
		"/neutral_diet.user.v1.UserService/GetUserSettings",
		svc.GetUserSettings,
		opts...,
	))
	mux.Handle("/neutral_diet.user.v1.UserService/UpdateUserSettings", connect_go.NewUnaryHandler(
		"/neutral_diet.user.v1.UserService/UpdateUserSettings",
		svc.UpdateUserSettings,
		opts...,
	))
	return "/neutral_diet.user.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) AddFoodItem(context.Context, *connect_go.Request[v1.AddFoodItemRequest]) (*connect_go.Response[v1.AddFoodItemResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.user.v1.UserService.AddFoodItem is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.user.v1.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.user.v1.UserService.DeleteUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUserSettings(context.Context, *connect_go.Request[v1.GetUserSettingsRequest]) (*connect_go.Response[v1.GetUserSettingsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.user.v1.UserService.GetUserSettings is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUserSettings(context.Context, *connect_go.Request[v1.UpdateUserSettingsRequest]) (*connect_go.Response[v1.UpdateUserSettingsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.user.v1.UserService.UpdateUserSettings is not implemented"))
}
