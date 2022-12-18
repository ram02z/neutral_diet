// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: neutral_diet/food/v1/api.proto

package foodv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
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
	// FoodServiceName is the fully-qualified name of the FoodService service.
	FoodServiceName = "neutral_diet.food.v1.FoodService"
)

// FoodServiceClient is a client for the neutral_diet.food.v1.FoodService service.
type FoodServiceClient interface {
	CreateLifeCycle(context.Context, *connect_go.Request[v1.CreateLifeCycleRequest]) (*connect_go.Response[v1.CreateLifeCycleResponse], error)
	CreateFoodItem(context.Context, *connect_go.Request[v1.CreateFoodItemRequest]) (*connect_go.Response[v1.CreateFoodItemResponse], error)
	CreateTypology(context.Context, *connect_go.Request[v1.CreateTypologyRequest]) (*connect_go.Response[v1.CreateTypologyResponse], error)
	CreateSubTypology(context.Context, *connect_go.Request[v1.CreateSubTypologyRequest]) (*connect_go.Response[v1.CreateSubTypologyResponse], error)
	CreateSource(context.Context, *connect_go.Request[v1.CreateSourceRequest]) (*connect_go.Response[v1.CreateSourceResponse], error)
	CreateRegion(context.Context, *connect_go.Request[v1.CreateRegionRequest]) (*connect_go.Response[v1.CreateRegionResponse], error)
	ListAggregateFoodItems(context.Context, *connect_go.Request[v1.ListAggregateFoodItemsRequest]) (*connect_go.Response[v1.ListAggregateFoodItemsResponse], error)
	// TODO: move to own service
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
}

// NewFoodServiceClient constructs a client for the neutral_diet.food.v1.FoodService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFoodServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) FoodServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &foodServiceClient{
		createLifeCycle: connect_go.NewClient[v1.CreateLifeCycleRequest, v1.CreateLifeCycleResponse](
			httpClient,
			baseURL+"/neutral_diet.food.v1.FoodService/CreateLifeCycle",
			opts...,
		),
		createFoodItem: connect_go.NewClient[v1.CreateFoodItemRequest, v1.CreateFoodItemResponse](
			httpClient,
			baseURL+"/neutral_diet.food.v1.FoodService/CreateFoodItem",
			opts...,
		),
		createTypology: connect_go.NewClient[v1.CreateTypologyRequest, v1.CreateTypologyResponse](
			httpClient,
			baseURL+"/neutral_diet.food.v1.FoodService/CreateTypology",
			opts...,
		),
		createSubTypology: connect_go.NewClient[v1.CreateSubTypologyRequest, v1.CreateSubTypologyResponse](
			httpClient,
			baseURL+"/neutral_diet.food.v1.FoodService/CreateSubTypology",
			opts...,
		),
		createSource: connect_go.NewClient[v1.CreateSourceRequest, v1.CreateSourceResponse](
			httpClient,
			baseURL+"/neutral_diet.food.v1.FoodService/CreateSource",
			opts...,
		),
		createRegion: connect_go.NewClient[v1.CreateRegionRequest, v1.CreateRegionResponse](
			httpClient,
			baseURL+"/neutral_diet.food.v1.FoodService/CreateRegion",
			opts...,
		),
		listAggregateFoodItems: connect_go.NewClient[v1.ListAggregateFoodItemsRequest, v1.ListAggregateFoodItemsResponse](
			httpClient,
			baseURL+"/neutral_diet.food.v1.FoodService/ListAggregateFoodItems",
			opts...,
		),
		createUser: connect_go.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+"/neutral_diet.food.v1.FoodService/CreateUser",
			opts...,
		),
	}
}

// foodServiceClient implements FoodServiceClient.
type foodServiceClient struct {
	createLifeCycle        *connect_go.Client[v1.CreateLifeCycleRequest, v1.CreateLifeCycleResponse]
	createFoodItem         *connect_go.Client[v1.CreateFoodItemRequest, v1.CreateFoodItemResponse]
	createTypology         *connect_go.Client[v1.CreateTypologyRequest, v1.CreateTypologyResponse]
	createSubTypology      *connect_go.Client[v1.CreateSubTypologyRequest, v1.CreateSubTypologyResponse]
	createSource           *connect_go.Client[v1.CreateSourceRequest, v1.CreateSourceResponse]
	createRegion           *connect_go.Client[v1.CreateRegionRequest, v1.CreateRegionResponse]
	listAggregateFoodItems *connect_go.Client[v1.ListAggregateFoodItemsRequest, v1.ListAggregateFoodItemsResponse]
	createUser             *connect_go.Client[v1.CreateUserRequest, v1.CreateUserResponse]
}

// CreateLifeCycle calls neutral_diet.food.v1.FoodService.CreateLifeCycle.
func (c *foodServiceClient) CreateLifeCycle(ctx context.Context, req *connect_go.Request[v1.CreateLifeCycleRequest]) (*connect_go.Response[v1.CreateLifeCycleResponse], error) {
	return c.createLifeCycle.CallUnary(ctx, req)
}

// CreateFoodItem calls neutral_diet.food.v1.FoodService.CreateFoodItem.
func (c *foodServiceClient) CreateFoodItem(ctx context.Context, req *connect_go.Request[v1.CreateFoodItemRequest]) (*connect_go.Response[v1.CreateFoodItemResponse], error) {
	return c.createFoodItem.CallUnary(ctx, req)
}

// CreateTypology calls neutral_diet.food.v1.FoodService.CreateTypology.
func (c *foodServiceClient) CreateTypology(ctx context.Context, req *connect_go.Request[v1.CreateTypologyRequest]) (*connect_go.Response[v1.CreateTypologyResponse], error) {
	return c.createTypology.CallUnary(ctx, req)
}

// CreateSubTypology calls neutral_diet.food.v1.FoodService.CreateSubTypology.
func (c *foodServiceClient) CreateSubTypology(ctx context.Context, req *connect_go.Request[v1.CreateSubTypologyRequest]) (*connect_go.Response[v1.CreateSubTypologyResponse], error) {
	return c.createSubTypology.CallUnary(ctx, req)
}

// CreateSource calls neutral_diet.food.v1.FoodService.CreateSource.
func (c *foodServiceClient) CreateSource(ctx context.Context, req *connect_go.Request[v1.CreateSourceRequest]) (*connect_go.Response[v1.CreateSourceResponse], error) {
	return c.createSource.CallUnary(ctx, req)
}

// CreateRegion calls neutral_diet.food.v1.FoodService.CreateRegion.
func (c *foodServiceClient) CreateRegion(ctx context.Context, req *connect_go.Request[v1.CreateRegionRequest]) (*connect_go.Response[v1.CreateRegionResponse], error) {
	return c.createRegion.CallUnary(ctx, req)
}

// ListAggregateFoodItems calls neutral_diet.food.v1.FoodService.ListAggregateFoodItems.
func (c *foodServiceClient) ListAggregateFoodItems(ctx context.Context, req *connect_go.Request[v1.ListAggregateFoodItemsRequest]) (*connect_go.Response[v1.ListAggregateFoodItemsResponse], error) {
	return c.listAggregateFoodItems.CallUnary(ctx, req)
}

// CreateUser calls neutral_diet.food.v1.FoodService.CreateUser.
func (c *foodServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// FoodServiceHandler is an implementation of the neutral_diet.food.v1.FoodService service.
type FoodServiceHandler interface {
	CreateLifeCycle(context.Context, *connect_go.Request[v1.CreateLifeCycleRequest]) (*connect_go.Response[v1.CreateLifeCycleResponse], error)
	CreateFoodItem(context.Context, *connect_go.Request[v1.CreateFoodItemRequest]) (*connect_go.Response[v1.CreateFoodItemResponse], error)
	CreateTypology(context.Context, *connect_go.Request[v1.CreateTypologyRequest]) (*connect_go.Response[v1.CreateTypologyResponse], error)
	CreateSubTypology(context.Context, *connect_go.Request[v1.CreateSubTypologyRequest]) (*connect_go.Response[v1.CreateSubTypologyResponse], error)
	CreateSource(context.Context, *connect_go.Request[v1.CreateSourceRequest]) (*connect_go.Response[v1.CreateSourceResponse], error)
	CreateRegion(context.Context, *connect_go.Request[v1.CreateRegionRequest]) (*connect_go.Response[v1.CreateRegionResponse], error)
	ListAggregateFoodItems(context.Context, *connect_go.Request[v1.ListAggregateFoodItemsRequest]) (*connect_go.Response[v1.ListAggregateFoodItemsResponse], error)
	// TODO: move to own service
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
}

// NewFoodServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFoodServiceHandler(svc FoodServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/neutral_diet.food.v1.FoodService/CreateLifeCycle", connect_go.NewUnaryHandler(
		"/neutral_diet.food.v1.FoodService/CreateLifeCycle",
		svc.CreateLifeCycle,
		opts...,
	))
	mux.Handle("/neutral_diet.food.v1.FoodService/CreateFoodItem", connect_go.NewUnaryHandler(
		"/neutral_diet.food.v1.FoodService/CreateFoodItem",
		svc.CreateFoodItem,
		opts...,
	))
	mux.Handle("/neutral_diet.food.v1.FoodService/CreateTypology", connect_go.NewUnaryHandler(
		"/neutral_diet.food.v1.FoodService/CreateTypology",
		svc.CreateTypology,
		opts...,
	))
	mux.Handle("/neutral_diet.food.v1.FoodService/CreateSubTypology", connect_go.NewUnaryHandler(
		"/neutral_diet.food.v1.FoodService/CreateSubTypology",
		svc.CreateSubTypology,
		opts...,
	))
	mux.Handle("/neutral_diet.food.v1.FoodService/CreateSource", connect_go.NewUnaryHandler(
		"/neutral_diet.food.v1.FoodService/CreateSource",
		svc.CreateSource,
		opts...,
	))
	mux.Handle("/neutral_diet.food.v1.FoodService/CreateRegion", connect_go.NewUnaryHandler(
		"/neutral_diet.food.v1.FoodService/CreateRegion",
		svc.CreateRegion,
		opts...,
	))
	mux.Handle("/neutral_diet.food.v1.FoodService/ListAggregateFoodItems", connect_go.NewUnaryHandler(
		"/neutral_diet.food.v1.FoodService/ListAggregateFoodItems",
		svc.ListAggregateFoodItems,
		opts...,
	))
	mux.Handle("/neutral_diet.food.v1.FoodService/CreateUser", connect_go.NewUnaryHandler(
		"/neutral_diet.food.v1.FoodService/CreateUser",
		svc.CreateUser,
		opts...,
	))
	return "/neutral_diet.food.v1.FoodService/", mux
}

// UnimplementedFoodServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedFoodServiceHandler struct{}

func (UnimplementedFoodServiceHandler) CreateLifeCycle(context.Context, *connect_go.Request[v1.CreateLifeCycleRequest]) (*connect_go.Response[v1.CreateLifeCycleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.food.v1.FoodService.CreateLifeCycle is not implemented"))
}

func (UnimplementedFoodServiceHandler) CreateFoodItem(context.Context, *connect_go.Request[v1.CreateFoodItemRequest]) (*connect_go.Response[v1.CreateFoodItemResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.food.v1.FoodService.CreateFoodItem is not implemented"))
}

func (UnimplementedFoodServiceHandler) CreateTypology(context.Context, *connect_go.Request[v1.CreateTypologyRequest]) (*connect_go.Response[v1.CreateTypologyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.food.v1.FoodService.CreateTypology is not implemented"))
}

func (UnimplementedFoodServiceHandler) CreateSubTypology(context.Context, *connect_go.Request[v1.CreateSubTypologyRequest]) (*connect_go.Response[v1.CreateSubTypologyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.food.v1.FoodService.CreateSubTypology is not implemented"))
}

func (UnimplementedFoodServiceHandler) CreateSource(context.Context, *connect_go.Request[v1.CreateSourceRequest]) (*connect_go.Response[v1.CreateSourceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.food.v1.FoodService.CreateSource is not implemented"))
}

func (UnimplementedFoodServiceHandler) CreateRegion(context.Context, *connect_go.Request[v1.CreateRegionRequest]) (*connect_go.Response[v1.CreateRegionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.food.v1.FoodService.CreateRegion is not implemented"))
}

func (UnimplementedFoodServiceHandler) ListAggregateFoodItems(context.Context, *connect_go.Request[v1.ListAggregateFoodItemsRequest]) (*connect_go.Response[v1.ListAggregateFoodItemsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.food.v1.FoodService.ListAggregateFoodItems is not implemented"))
}

func (UnimplementedFoodServiceHandler) CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neutral_diet.food.v1.FoodService.CreateUser is not implemented"))
}
