import { Weight } from '@/core/weight';

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
  overallAverageCarbonFootprint: number;
  overallCarbonFootprint: number;
  noEntries: number;
}
