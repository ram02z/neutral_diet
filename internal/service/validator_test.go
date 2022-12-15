package service

import (
	"testing"

	foodv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
	"github.com/stretchr/testify/require"
)

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
		"Invalid CfType": {
			build: func() *foodv1.CreateFoodItemRequest {
				p := buildValidRequest()
				p.FoodItem.CfType = -1
				return p
			},
			expect: func(t *testing.T, err error) {
				require.EqualError(t, err, `invalid CreateFoodItemRequest.FoodItem: embedded message failed validation | caused by: invalid FoodItem.CfType: value must be one of the defined enum values`)
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
