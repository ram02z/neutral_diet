import { Weight } from "@/core/weight";

export type LocalUserSettings = {
  region: number;
  cfLimit: number;
  dietaryRequirement: number;
  dirty: boolean;
};

export type LocalFoodLogItem = {
  dbId: number;
  foodItemId: number;
  name: string;
  weight: Weight;
  carbonFootprint: number;
  region: number;
};

export type FoodLogStats = {
  totalCarbonFootprint: number;
  carbonFootprintRemaining: number;
  carbonFootprintGoalPercent: number;
};

export type UserInsights = {
  overallUserAverage: number;
  overallUser: number;
  dailyGlobalAverage: number;
  dailyGlobalAverageUserDietaryRequirement: number;
  noUserEntries: number;
  streakLength: number;
  isStreakActive: boolean;
}

export type SerializableDate = {
  year: number;
  month: number;
  day: number;
}
