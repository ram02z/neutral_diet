import { QuantityUnit } from '@/core/food_unit';

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
  quantity: QuantityUnit;
  carbonFootprint: number;
  region: number;
  meal: number;
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
};

export type UserProgress = {
  all: Record<string, number>;
  meal: Map<number, Record<string, number>>;
};

export type LocalUserGoal = {
  dbId: number;
  description: string;
  startDate: SerializableDate | undefined;
  endDate: SerializableDate | undefined;
  startCarbonFootprint: number;
  targetCarbonFootprint: number;
}

export type LocalUserGoals = {
  active: LocalUserGoal[];
  completed: LocalUserGoal[];
}

export type SerializableDate = {
  year: number;
  month: number;
  day: number;
};
