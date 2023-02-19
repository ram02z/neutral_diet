import { Weight } from "@/core/weight";

export type LocalUserSettings = {
  region: string;
  cfLimit: number;
  dietaryRequirement: number;
  dirty: boolean;
};

export type LocalFoodLogItem = {
  dbId: number;
  name: string;
  weight: Weight;
  carbonFootprint: number;
};
