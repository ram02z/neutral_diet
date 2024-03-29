// @generated by protoc-gen-connect-web v0.4.0 with parameter "target=ts"
// @generated from file neutral_diet/user/v1/api.proto (package neutral_diet.user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddCarbonFootprintGoalRequest, AddCarbonFootprintGoalResponse, AddDeviceRequest, AddDeviceResponse, AddFoodItemRequest, AddFoodItemResponse, CreateUserRequest, CreateUserResponse, DeleteCarbonFootprintGoalRequest, DeleteCarbonFootprintGoalResponse, DeleteFoodItemRequest, DeleteFoodItemResponse, DeleteUserRequest, DeleteUserResponse, GetCarbonFootprintGoalsRequest, GetCarbonFootprintGoalsResponse, GetFoodItemLogDaysRequest, GetFoodItemLogDaysResponse, GetFoodItemLogRequest, GetFoodItemLogResponse, GetUserInsightsRequest, GetUserInsightsResponse, GetUserProgressRequest, GetUserProgressResponse, GetUserSettingsRequest, GetUserSettingsResponse, UpdateCarbonFootprintGoalRequest, UpdateCarbonFootprintGoalResponse, UpdateFoodItemRequest, UpdateFoodItemResponse, UpdateUserSettingsRequest, UpdateUserSettingsResponse } from "./api_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service neutral_diet.user.v1.UserService
 */
export const UserService = {
  typeName: "neutral_diet.user.v1.UserService",
  methods: {
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.AddFoodItem
     */
    addFoodItem: {
      name: "AddFoodItem",
      I: AddFoodItemRequest,
      O: AddFoodItemResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.UpdateFoodItem
     */
    updateFoodItem: {
      name: "UpdateFoodItem",
      I: UpdateFoodItemRequest,
      O: UpdateFoodItemResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.DeleteFoodItem
     */
    deleteFoodItem: {
      name: "DeleteFoodItem",
      I: DeleteFoodItemRequest,
      O: DeleteFoodItemResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.GetFoodItemLog
     */
    getFoodItemLog: {
      name: "GetFoodItemLog",
      I: GetFoodItemLogRequest,
      O: GetFoodItemLogResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.GetFoodItemLogDays
     */
    getFoodItemLogDays: {
      name: "GetFoodItemLogDays",
      I: GetFoodItemLogDaysRequest,
      O: GetFoodItemLogDaysResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: CreateUserRequest,
      O: CreateUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.DeleteUser
     */
    deleteUser: {
      name: "DeleteUser",
      I: DeleteUserRequest,
      O: DeleteUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.GetUserSettings
     */
    getUserSettings: {
      name: "GetUserSettings",
      I: GetUserSettingsRequest,
      O: GetUserSettingsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.UpdateUserSettings
     */
    updateUserSettings: {
      name: "UpdateUserSettings",
      I: UpdateUserSettingsRequest,
      O: UpdateUserSettingsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.GetUserInsights
     */
    getUserInsights: {
      name: "GetUserInsights",
      I: GetUserInsightsRequest,
      O: GetUserInsightsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.GetUserProgress
     */
    getUserProgress: {
      name: "GetUserProgress",
      I: GetUserProgressRequest,
      O: GetUserProgressResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.AddCarbonFootprintGoal
     */
    addCarbonFootprintGoal: {
      name: "AddCarbonFootprintGoal",
      I: AddCarbonFootprintGoalRequest,
      O: AddCarbonFootprintGoalResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.UpdateCarbonFootprintGoal
     */
    updateCarbonFootprintGoal: {
      name: "UpdateCarbonFootprintGoal",
      I: UpdateCarbonFootprintGoalRequest,
      O: UpdateCarbonFootprintGoalResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.DeleteCarbonFootprintGoal
     */
    deleteCarbonFootprintGoal: {
      name: "DeleteCarbonFootprintGoal",
      I: DeleteCarbonFootprintGoalRequest,
      O: DeleteCarbonFootprintGoalResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.GetCarbonFootprintGoals
     */
    getCarbonFootprintGoals: {
      name: "GetCarbonFootprintGoals",
      I: GetCarbonFootprintGoalsRequest,
      O: GetCarbonFootprintGoalsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc neutral_diet.user.v1.UserService.AddDevice
     */
    addDevice: {
      name: "AddDevice",
      I: AddDeviceRequest,
      O: AddDeviceResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

