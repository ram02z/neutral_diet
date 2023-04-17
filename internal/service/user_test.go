package service_test

import (
	"testing"

	"github.com/google/uuid"
	userv1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1"
	"github.com/stretchr/testify/require"
)

func TestAddFoodItem(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.AddFoodItemRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.AddFoodItemRequest {
				return &userv1.AddFoodItemRequest{
					FoodLogItem: &userv1.FoodLogItemRequest{},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty FoodLogItem": {
			build: func() *userv1.AddFoodItemRequest {
				return &userv1.AddFoodItemRequest{}
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

func TestUpdateFoodItem(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.UpdateFoodItemRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.UpdateFoodItemRequest {
				return &userv1.UpdateFoodItemRequest{
					Id: 1,
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Invalid Id": {
			build: func() *userv1.UpdateFoodItemRequest {
				return &userv1.UpdateFoodItemRequest{}
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

func TestDeleteFoodItem(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.DeleteFoodItemRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.DeleteFoodItemRequest {
				return &userv1.DeleteFoodItemRequest{
					Id: 1,
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Invalid Id": {
			build: func() *userv1.DeleteFoodItemRequest {
				return &userv1.DeleteFoodItemRequest{}
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

func TestGetFoodItemLogDays(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.GetFoodItemLogDaysRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.GetFoodItemLogDaysRequest {
				return &userv1.GetFoodItemLogDaysRequest{
					Month: 1,
					Year:  1,
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Invalid Month": {
			build: func() *userv1.GetFoodItemLogDaysRequest {
				return &userv1.GetFoodItemLogDaysRequest{
					Month: 0,
					Year:  1,
				}
			},
			expect: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		"Invalid Year": {
			build: func() *userv1.GetFoodItemLogDaysRequest {
				return &userv1.GetFoodItemLogDaysRequest{
					Month: 1,
					Year:  0,
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

func TestGetFoodItemLog(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.GetFoodItemLogRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.GetFoodItemLogRequest {
				return &userv1.GetFoodItemLogRequest{
					Date: &userv1.Date{
						Year:  1,
						Month: 1,
						Day:   1,
					},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty Date": {
			build: func() *userv1.GetFoodItemLogRequest {
				return &userv1.GetFoodItemLogRequest{}
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

func TestCreateUser(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.CreateUserRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.CreateUserRequest {
				return &userv1.CreateUserRequest{
					FirebaseUid: uuid.New().String(),
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Invalid FirebaseUid": {
			build: func() *userv1.CreateUserRequest {
				return &userv1.CreateUserRequest{
					FirebaseUid: "",
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

func TestUpdateUserSettings(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.UpdateUserSettingsRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.UpdateUserSettingsRequest {
				return &userv1.UpdateUserSettingsRequest{
					UserSettings: &userv1.UserSettings{},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty UserSettings": {
			build: func() *userv1.UpdateUserSettingsRequest {
				return &userv1.UpdateUserSettingsRequest{}
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

func TestGetUserSettings(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.GetUserSettingsResponse
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.GetUserSettingsResponse {
				return &userv1.GetUserSettingsResponse{
					UserSettings: &userv1.UserSettings{},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty UserSettings": {
			build: func() *userv1.GetUserSettingsResponse {
				return &userv1.GetUserSettingsResponse{}
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

func TestAddCarbonFootprintGoal(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.AddCarbonFootprintGoalRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.AddCarbonFootprintGoalRequest {
				return &userv1.AddCarbonFootprintGoalRequest{
					CarbonFootprintGoal: &userv1.CarbonFootprintGoalRequest{},
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Empty CarbonFootprintGoalRequest": {
			build: func() *userv1.AddCarbonFootprintGoalRequest {
				return &userv1.AddCarbonFootprintGoalRequest{}
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

func TestUpdateCarbonFootprintGoal(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.UpdateCarbonFootprintGoalRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.UpdateCarbonFootprintGoalRequest {
				return &userv1.UpdateCarbonFootprintGoalRequest{
					Id: 1,
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Invalid Id": {
			build: func() *userv1.UpdateCarbonFootprintGoalRequest {
				return &userv1.UpdateCarbonFootprintGoalRequest{}
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

func TestDeleteCarbonFootprintGoal(t *testing.T) {
	cases := map[string]struct {
		build  func() *userv1.DeleteCarbonFootprintGoalRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: func() *userv1.DeleteCarbonFootprintGoalRequest {
				return &userv1.DeleteCarbonFootprintGoalRequest{
					Id: 1,
				}
			},
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Invalid Id": {
			build: func() *userv1.DeleteCarbonFootprintGoalRequest {
				return &userv1.DeleteCarbonFootprintGoalRequest{}
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
