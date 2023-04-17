package service_test

import (
	"testing"

	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
	"github.com/stretchr/testify/require"
)

func TestCreateLifeCycle(t *testing.T) {
	cases := map[string]struct {
		build  func() *foodv1.CreateLifeCycleRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *foodv1.CreateLifeCycleRequest {
				return &foodv1.CreateLifeCycleRequest{
					LifeCycle: &foodv1.LifeCycle{},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty LifeCycle": {
			build: func() *foodv1.CreateLifeCycleRequest {
				return &foodv1.CreateLifeCycleRequest{}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			r := tc.build()
			err := r.Validate()
			tc.expect(t, err)
		})
	}
}

func TestCreateFoodItem(t *testing.T) {
	buildValidRequest := func() *foodv1.CreateFoodItemRequest {
		return &foodv1.CreateFoodItemRequest{
			FoodItem: &foodv1.FoodItem{
				Name:       "Apple",
				TypologyId: 2,
				CfType:     foodv1.FoodItem_CF_TYPE_ITEM,
			},
		}
	}

	cases := map[string]struct {
		build  func() *foodv1.CreateFoodItemRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build:  buildValidRequest,
			expect: func(t *testing.T, err error) { require.NoError(t, err) },
		},
		"Empty FoodItem": {
			build: func() *foodv1.CreateFoodItemRequest {
				return &foodv1.CreateFoodItemRequest{}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		"Invalid CfType": {
			build: func() *foodv1.CreateFoodItemRequest {
				p := buildValidRequest()
				p.FoodItem.CfType = -1
				return p
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			r := tc.build()
			err := r.Validate()
			tc.expect(t, err)
		})
	}
}

func TestCreateTypology(t *testing.T) {
	cases := map[string]struct {
		build  func() *foodv1.CreateTypologyRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *foodv1.CreateTypologyRequest {
				return &foodv1.CreateTypologyRequest{
					Typology: &foodv1.Typology{},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty Typology": {
			build: func() *foodv1.CreateTypologyRequest {
				return &foodv1.CreateTypologyRequest{}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			r := tc.build()
			err := r.Validate()
			tc.expect(t, err)
		})
	}
}

func TestCreateSubTypology(t *testing.T) {
	cases := map[string]struct {
		build  func() *foodv1.CreateSubTypologyRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *foodv1.CreateSubTypologyRequest {
				return &foodv1.CreateSubTypologyRequest{
					SubTypology: &foodv1.SubTypology{},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty SubTypology": {
			build: func() *foodv1.CreateSubTypologyRequest {
				return &foodv1.CreateSubTypologyRequest{}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			r := tc.build()
			err := r.Validate()
			tc.expect(t, err)
		})
	}
}

func TestCreateSource(t *testing.T) {
	cases := map[string]struct {
		build  func() *foodv1.CreateSourceRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *foodv1.CreateSourceRequest {
				return &foodv1.CreateSourceRequest{
					Source: &foodv1.Source{},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty Source": {
			build: func() *foodv1.CreateSourceRequest {
				return &foodv1.CreateSourceRequest{}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			r := tc.build()
			err := r.Validate()
			tc.expect(t, err)
		})
	}
}

func TestListAggregateFoodItems(t *testing.T) {
	cases := map[string]struct {
		build  func() *foodv1.ListAggregateFoodItemsRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *foodv1.ListAggregateFoodItemsRequest {
				return &foodv1.ListAggregateFoodItemsRequest{
					Region: foodv1.Region_REGION_ASIA,
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Invalid Region": {
			build: func() *foodv1.ListAggregateFoodItemsRequest {
				return &foodv1.ListAggregateFoodItemsRequest{
					Region: -1,
				}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			r := tc.build()
			err := r.Validate()
			tc.expect(t, err)
		})
	}
}

func TestGetFoodItemInfo(t *testing.T) {
	cases := map[string]struct {
		build  func() *foodv1.GetFoodItemInfoRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *foodv1.GetFoodItemInfoRequest {
				return &foodv1.GetFoodItemInfoRequest{
					Id:     1,
					Region: foodv1.Region_REGION_ASIA,
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Invalid Region": {
			build: func() *foodv1.GetFoodItemInfoRequest {
				return &foodv1.GetFoodItemInfoRequest{
					Id:     1,
					Region: -1,
				}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		"Invalid ID": {
			build: func() *foodv1.GetFoodItemInfoRequest {
				return &foodv1.GetFoodItemInfoRequest{}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			r := tc.build()
			err := r.Validate()
			tc.expect(t, err)
		})
	}
}
