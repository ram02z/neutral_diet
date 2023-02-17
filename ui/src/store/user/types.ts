export type LocalUserSettings = {
  region: string;
  cfLimit: number;
  dietaryRequirement: number;
  dirty: boolean;
};

export type LocalFoodLogItem = {
  dbId: number;
  name: string;
  weight: number;
  carbonFootprint: number;
};
